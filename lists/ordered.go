package lists

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func Max[T constraints.Ordered](list []T) T {
	var result T
	for _, value := range list {
		if value > result {
			result = value
		}
	}
	return result
}

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

func Sort[T constraints.Ordered](list []T) []T {
	slices.Sort(list)
	return list
}
