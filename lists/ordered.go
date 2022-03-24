package lists

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Max returns the highest value in the list.
// If the list is empty, the zero value is returned.
func Max[T constraints.Ordered](list []T) T {
	var result T
	if len(list) == 0 {
		return result
	}
	result = list[0]
	for _, value := range list[1:] {
		if value > result {
			result = value
		}
	}
	return result
}

// Min returns the lowest value in the list.
// If the list is empty, the zero value is returned.
func Min[T constraints.Ordered](list []T) T {
	var result T
	if len(list) == 0 {
		return result
	}
	result = list[0]
	for _, value := range list[1:] {
		if value < result {
			result = value
		}
	}
	return result
}

// Sort returns the given list in sorted order.
// It returns a copy of the list and leaves the original
// list unchanged.
func Sort[L ~[]T, T constraints.Ordered](list L) L {
	result := make(L, len(list))
	copy(result, list)
	slices.Sort(result)
	return result
}
