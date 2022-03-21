package lists

import "golang.org/x/exp/slices"

func Uniq[L ~[]T, T comparable](list L) L {
	return slices.Compact(list)
}
