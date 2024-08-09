package builder

import "fmt"

type Pagination struct {
	Limit  interface{}
	Offset interface{}
}

func ConvertLimit(limit int) string {
	if limit == 0 {
		return ""
	}
	return fmt.Sprintf("LIMIT %d", limit)
}

func ConvertOffset(offset int) string {
	if offset == 0 {
		return ""
	}
	return fmt.Sprintf("OFFSET %d", offset)
}

func ConvertLimitOffset(limit, offset int) string {
	if limit == 0 && offset == 0 {
		return ""
	}
	return fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)
}

func ConvertPagination(p Pagination) string {
	limit := ""
	if p.Limit != nil {
		limit = fmt.Sprintf("LIMIT %d", p.Limit)
	}
	offset := ""
	if p.Offset != nil {
		offset = fmt.Sprintf("OFFSET %d", p.Offset)
	}
	return fmt.Sprintf("%s %s", limit, offset)
}
