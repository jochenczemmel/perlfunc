package hashes

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	return maps.Keys(m)
}

func SortKeys[M ~map[K]V, K constraints.Ordered, V any](m M) []K {
	keys := maps.Keys(m)
	slices.Sort(keys)
	return keys
}

func Values[M ~map[K]V, K comparable, V any](m M) []V {
	return maps.Values(m)
}

func Grep[M ~map[K]V, K comparable, V any](m M,
	f func(K, V) bool) map[K]V {

	result := map[K]V{}
	for key, value := range m {
		if f(key, value) {
			result[key] = value
		}
	}
	return result
}

func Map[M ~map[K]V, K comparable, V any](m M,
	f func(K, V) (K, V)) map[K]V {

	result := map[K]V{}
	for key, value := range m {
		k, v := f(key, value)
		result[k] = v
	}
	return result
}
