package utils

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/pkg/errors"
)

/*
This function validates and escapes the SQL query.

I copied it from somewhere, but I can't remember where is original source. Tell me if you know.
*/
func EscapeSQL(sql string, args ...interface{}) ([]byte, error) {
	buf := make([]byte, 0, len(sql))
	argPos := 0
	for i := 0; i < len(sql); i++ {
		q := strings.IndexByte(sql[i:], '%')
		if q == -1 {
			buf = append(buf, sql[i:]...)
			break
		}
		buf = append(buf, sql[i:i+q]...)
		i += q

		ch := byte(0)
		if i+1 < len(sql) {
			ch = sql[i+1] // get the specifier
		}
		switch ch {
		case 'n':
			if argPos >= len(args) {
				return nil, errors.Errorf("missing arguments, need %d-th arg, but only got %d args", argPos+1, len(args))
			}
			arg := args[argPos]
			argPos++

			v, ok := arg.(string)
			if !ok {
				return nil, errors.Errorf("expect a string identifier, got %v", arg)
			}
			buf = append(buf, '`')
			buf = append(buf, strings.ReplaceAll(v, "`", "``")...)
			buf = append(buf, '`')
			i++ // skip specifier
		case '?':
			if argPos >= len(args) {
				return nil, errors.Errorf("missing arguments, need %d-th arg, but only got %d args", argPos+1, len(args))
			}
			arg := args[argPos]
			argPos++

			if arg == nil {
				buf = append(buf, "NULL"...)
			} else {
				switch v := arg.(type) {
				case int:
					buf = strconv.AppendInt(buf, int64(v), 10)
				case int8:
					buf = strconv.AppendInt(buf, int64(v), 10)
				case int16:
					buf = strconv.AppendInt(buf, int64(v), 10)
				case int32:
					buf = strconv.AppendInt(buf, int64(v), 10)
				case int64:
					buf = strconv.AppendInt(buf, v, 10)
				case uint:
					buf = strconv.AppendUint(buf, uint64(v), 10)
				case uint8:
					buf = strconv.AppendUint(buf, uint64(v), 10)
				case uint16:
					buf = strconv.AppendUint(buf, uint64(v), 10)
				case uint32:
					buf = strconv.AppendUint(buf, uint64(v), 10)
				case uint64:
					buf = strconv.AppendUint(buf, v, 10)
				case float32:
					buf = strconv.AppendFloat(buf, float64(v), 'g', -1, 32)
				case float64:
					buf = strconv.AppendFloat(buf, v, 'g', -1, 64)
				case bool:
					buf = appendSQLArgBool(buf, v)
				case time.Time:
					if v.IsZero() {
						buf = append(buf, "'0000-00-00'"...)
					} else {
						buf = append(buf, '\'')
						buf = v.AppendFormat(buf, "2006-01-02 15:04:05.999999")
						buf = append(buf, '\'')
					}
				case json.RawMessage:
					buf = append(buf, '\'')
					buf = escapeBytesBackslash(buf, v)
					buf = append(buf, '\'')
				case []byte:
					if v == nil {
						buf = append(buf, "NULL"...)
					} else {
						buf = append(buf, "_binary'"...)
						buf = escapeBytesBackslash(buf, v)
						buf = append(buf, '\'')
					}
				case string:
					buf = appendSQLArgString(buf, v)
				case []string:
					for i, k := range v {
						if i > 0 {
							buf = append(buf, ',')
						}
						buf = append(buf, '\'')
						buf = escapeStringBackslash(buf, k)
						buf = append(buf, '\'')
					}
				case []float32:
					for i, k := range v {
						if i > 0 {
							buf = append(buf, ',')
						}
						buf = strconv.AppendFloat(buf, float64(k), 'g', -1, 32)
					}
				case []float64:
					for i, k := range v {
						if i > 0 {
							buf = append(buf, ',')
						}
						buf = strconv.AppendFloat(buf, k, 'g', -1, 64)
					}
				default:
					// slow path based on reflection
					reflectTp := reflect.TypeOf(arg)
					kind := reflectTp.Kind()
					switch kind {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						buf = strconv.AppendInt(buf, reflect.ValueOf(arg).Int(), 10)
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						buf = strconv.AppendUint(buf, reflect.ValueOf(arg).Uint(), 10)
					case reflect.Float32:
						buf = strconv.AppendFloat(buf, reflect.ValueOf(arg).Float(), 'g', -1, 32)
					case reflect.Float64:
						buf = strconv.AppendFloat(buf, reflect.ValueOf(arg).Float(), 'g', -1, 64)
					case reflect.Bool:
						buf = appendSQLArgBool(buf, reflect.ValueOf(arg).Bool())
					case reflect.String:
						buf = appendSQLArgString(buf, reflect.ValueOf(arg).String())
					default:
						return nil, errors.Errorf("unsupported %d-th argument: %v", argPos, arg)
					}
				}
			}
			i++ // skip specifier
		case '%':
			buf = append(buf, '%')
			i++ // skip specifier
		default:
			buf = append(buf, '%')
		}
	}
	return buf, nil
}

func EscapeString(s string) string {
	buf := make([]byte, 0, len(s))
	return string(escapeStringBackslash(buf, s))
}

func appendSQLArgBool(buf []byte, v bool) []byte {
	if v {
		return append(buf, '1')
	}
	return append(buf, '0')
}

func appendSQLArgString(buf []byte, s string) []byte {
	buf = append(buf, '\'')
	buf = escapeStringBackslash(buf, s)
	buf = append(buf, '\'')
	return buf
}

func escapeStringBackslash(buf []byte, v string) []byte {
	return escapeBytesBackslash(buf, unsafe.Slice(unsafe.StringData(v), len(v)))
}

// escapeBytesBackslash will escape []byte into the buffer, with backslash.
func escapeBytesBackslash(buf []byte, v []byte) []byte {
	pos := len(buf)
	buf = reserveBuffer(buf, len(v)*2)

	for _, c := range v {
		switch c {
		case '\x00':
			buf[pos] = '\\'
			buf[pos+1] = '0'
			pos += 2
		case '\n':
			buf[pos] = '\\'
			buf[pos+1] = 'n'
			pos += 2
		case '\r':
			buf[pos] = '\\'
			buf[pos+1] = 'r'
			pos += 2
		case '\x1a':
			buf[pos] = '\\'
			buf[pos+1] = 'Z'
			pos += 2
		case '\'':
			buf[pos] = '\\'
			buf[pos+1] = '\''
			pos += 2
		case '"':
			buf[pos] = '\\'
			buf[pos+1] = '"'
			pos += 2
		case '\\':
			buf[pos] = '\\'
			buf[pos+1] = '\\'
			pos += 2
		default:
			buf[pos] = c
			pos++
		}
	}

	return buf[:pos]
}

func reserveBuffer(buf []byte, appendSize int) []byte {
	newSize := len(buf) + appendSize
	if cap(buf) < newSize {
		newBuf := make([]byte, len(buf)*2+appendSize)
		copy(newBuf, buf)
		buf = newBuf
	}
	return buf[:newSize]
}
