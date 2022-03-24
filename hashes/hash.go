package hashes

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// Keys returns the keys of the map in unspecified order.
// It uses the experimental function maps.Keys().
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	return maps.Keys(m)
}

// SortKeys returns the sorted keys of the map.
// It uses the experimental function maps.Keys().
func SortKeys[M ~map[K]V, K constraints.Ordered, V any](m M) []K {
	keys := Keys(m)
	slices.Sort(keys)
	return keys
}

// Values returns the values of the map in unspecified order.
// It uses the experimental function maps.Values().
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	return maps.Values(m)
}

// Grep returns a new map with elements where f returns true.
func Grep[M ~map[K]V, K comparable, V any](m M,
	f func(K, V) bool) M {

	// result := map[K]V{}
	result := M{}
	for key, value := range m {
		if f(key, value) {
			result[key] = value
		}
	}
	return result
}

// Map returns a new map where the elements are modified by f.
func Map[M ~map[K]V, K comparable, V any](m M,
	f func(K, V) (K, V)) M {

	result := M{}
	for key, value := range m {
		k, v := f(key, value)
		result[k] = v
	}
	return result
}
