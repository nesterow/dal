package builder

import (
	"fmt"
	"strings"

	"github.com/nesterow/dal/pkg/utils"
)

func convertGroup(ctx Dialect, keys []string) string {
	set := utils.Map(keys, ctx.GetColumnName)
	return fmt.Sprintf("GROUP BY %s", strings.Join(set, ", "))
}
