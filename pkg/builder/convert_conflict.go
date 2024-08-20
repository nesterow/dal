package builder

import (
	"fmt"
	"strings"

	utils "github.com/nesterow/dal/pkg/utils"
)

func convertConflict(ctx Dialect, fields ...string) string {
	keys := utils.Map(fields, ctx.GetColumnName)
	return fmt.Sprintf("ON CONFLICT (%s)", strings.Join(keys, ","))
}
