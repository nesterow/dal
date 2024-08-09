package builder

import (
	"strings"

	"l12.xyz/dal/utils"
)

func ConvertGroup(ctx Context, keys []string) string {
	set := utils.Map(keys, ctx.GetColumnName)
	return strings.Join(set, ", ")
}
