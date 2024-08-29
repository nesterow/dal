package builder

import (
	"fmt"
	"strings"

	"l12.xyz/x/dal/pkg/utils"
)

func convertConflict(ctx Dialect, fields ...string) string {
	keys := utils.Map(fields, ctx.GetColumnName)
	return fmt.Sprintf("ON CONFLICT (%s)", strings.Join(keys, ","))
}
