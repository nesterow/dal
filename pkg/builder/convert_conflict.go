package builder

import (
	"fmt"
	"strings"

	utils "pkg/utils"
)

func convertConflict(ctx Dialect, fields ...string) string {
	keys := utils.Map(fields, ctx.GetColumnName)
	return fmt.Sprintf("ON CONFLICT (%s)", strings.Join(keys, ","))
}
