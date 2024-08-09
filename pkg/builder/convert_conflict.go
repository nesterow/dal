package builder

import (
	"fmt"
	"strings"

	utils "l12.xyz/dal/utils"
)

func ConvertConflict(ctx Context, fields ...string) string {
	keys := utils.Map(fields, ctx.GetColumnName)
	return fmt.Sprintf("ON CONFLICT (%s) DO", strings.Join(keys, ","))
}
