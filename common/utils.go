package common

import (
	"sort"
)

func SortAndRemoveEmptyDuplicates(a []string) (ret []string) {
	sort.Strings(a)
	for i := 0; i < len(a); i++ {
		if a[i] == "" || (i > 0 && a[i] == a[i-1]) {
			continue
		}
		ret = append(ret, a[i])
	}
	return ret
}
