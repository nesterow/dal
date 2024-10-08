package proto

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Response) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "i":
			z.Id, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "ra":
			z.RowsAffected, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "RowsAffected")
				return
			}
		case "li":
			z.LastInsertId, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "LastInsertId")
				return
			}
		case "e":
			z.Error, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Error")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Response) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "i"
	err = en.Append(0x84, 0xa1, 0x69)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.Id)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	// write "ra"
	err = en.Append(0xa2, 0x72, 0x61)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.RowsAffected)
	if err != nil {
		err = msgp.WrapError(err, "RowsAffected")
		return
	}
	// write "li"
	err = en.Append(0xa2, 0x6c, 0x69)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.LastInsertId)
	if err != nil {
		err = msgp.WrapError(err, "LastInsertId")
		return
	}
	// write "e"
	err = en.Append(0xa1, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Error)
	if err != nil {
		err = msgp.WrapError(err, "Error")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Response) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "i"
	o = append(o, 0x84, 0xa1, 0x69)
	o = msgp.AppendUint32(o, z.Id)
	// string "ra"
	o = append(o, 0xa2, 0x72, 0x61)
	o = msgp.AppendInt64(o, z.RowsAffected)
	// string "li"
	o = append(o, 0xa2, 0x6c, 0x69)
	o = msgp.AppendInt64(o, z.LastInsertId)
	// string "e"
	o = append(o, 0xa1, 0x65)
	o = msgp.AppendString(o, z.Error)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Response) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "i":
			z.Id, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "ra":
			z.RowsAffected, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RowsAffected")
				return
			}
		case "li":
			z.LastInsertId, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastInsertId")
				return
			}
		case "e":
			z.Error, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Error")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Response) Msgsize() (s int) {
	s = 1 + 2 + msgp.Uint32Size + 3 + msgp.Int64Size + 3 + msgp.Int64Size + 2 + msgp.StringPrefixSize + len(z.Error)
	return
}
