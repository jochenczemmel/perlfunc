package hashes_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/jochenczemmel/perlfunc/assert"
	"github.com/jochenczemmel/perlfunc/hashes"
)

var m1 = map[int]int{1: 2, 2: 4, 4: 8, 8: 16}
var m2 = map[int]string{
	1: "eins",
	2: "zwei",
	4: "vier",
	8: "acht",
}

func TestMap(t *testing.T) {
	want := map[int]string{
		1: "EINS",
		2: "ZWEI",
		4: "VIER",
		8: "ACHT",
	}
	got := hashes.Map(m2, func(key int, value string) (int, string) {
		return key, strings.ToUpper(value)
	})
	assert.EqualsMap(t, got, want)

	want = map[int]string{
		10: "eins",
		20: "zwei",
		40: "vier",
		80: "acht",
	}
	got = hashes.Map(m2, func(key int, value string) (int, string) {
		return key * 10, value
	})
	assert.EqualsMap(t, got, want)

}

func TestGrep(t *testing.T) {
	want := map[int]string{
		1: "eins",
		2: "zwei",
		4: "vier",
	}
	got := hashes.Grep(m2, func(_ int, value string) bool {
		return strings.Contains(value, "i")
	})
	assert.EqualsMap(t, got, want)

	got = hashes.Grep(m2, func(_ int, value string) bool {
		return strings.Contains(value, "X")
	})
	assert.EqualsMap(t, got, map[int]string{})

	want = map[int]string{
		4: "vier",
		8: "acht",
	}
	got = hashes.Grep(m2, func(key int, _ string) bool {
		return key > 2
	})
	assert.EqualsMap(t, got, want)
}

func TestKeys(t *testing.T) {
	want := []int{1, 2, 4, 8}

	got1 := hashes.Keys(m1)
	sort.Ints(got1)
	assert.EqualsList(t, got1, want)

	got2 := hashes.Keys(m2)
	sort.Ints(got2)
	assert.EqualsList(t, got2, want)
}

func TestSortKeys(t *testing.T) {
	want := []int{1, 2, 4, 8}

	got1 := hashes.SortKeys(m1)
	assert.EqualsList(t, got1, want)

	got2 := hashes.SortKeys(m2)
	assert.EqualsList(t, got2, want)
}

func TestValues(t *testing.T) {
	want := []string{"acht", "eins", "vier", "zwei"}

	got := hashes.Values(m2)
	sort.Strings(got)
	assert.EqualsList(t, got, want)
}
