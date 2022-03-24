package lists

import "golang.org/x/exp/slices"

// Uniq returns the unique values in the list.
// It uses the experimental function slices.Compact.
// It returns a copy of the list and leaves the original
// list unchanged.
func Uniq[L ~[]T, T comparable](list L) L {
	result := make(L, len(list))
	copy(result, list)
	return slices.Compact(result)
}
