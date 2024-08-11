package builder

import (
	"fmt"
	"strings"

	"l12.xyz/dal/utils"
)

func convertGroup(ctx Dialect, keys []string) string {
	set := utils.Map(keys, ctx.GetColumnName)
	return fmt.Sprintf("GROUP BY %s", strings.Join(set, ", "))
}
