package builder

import "sort"

func AggregateSortedKeys(maps []Map) []string {
	set := make(map[string]int)
	keys := make([]string, 0)
	for _, item := range maps {
		for k := range item {
			if set[k] == 1 {
				continue
			}
			keys = append(keys, k)
			set[k] = 1
		}
	}
	set = nil
	sort.Strings(keys)
	return keys
}
