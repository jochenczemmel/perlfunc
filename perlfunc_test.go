package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jochenczemmel/perlfunc/assert"
	"github.com/jochenczemmel/perlfunc/hashes"
	"github.com/jochenczemmel/perlfunc/lists"
)

func TestPerlfunc(t *testing.T) {
	daten := map[string]int{
		"Fukuoka":   1,
		"Kyoto":     1,
		"Osaka":     2,
		"Kagoshima": 3,
		"Kanazawa":  2,
		"Takamatsu": 2,
		"Kochi":     1,
		"Tokyo":     3,
		"Nagoya":    2,
	}

	t.Run("sortkeys", func(t *testing.T) {
		got := lists.Head(lists.Sort(hashes.Keys(daten)), 3)
		t.Logf("DEBUG: %v", got)

		want := []string{
			"Fukuoka",
			"Kagoshima",
			"Kanazawa",
		}
		assert.EqualsList(t, got, want)
	})

	t.Run("mapsortkeys", func(t *testing.T) {

		got := lists.Head(lists.Sort(hashes.Keys(
			hashes.Map(daten, func(k string, v int) (string, int) {
				return strings.ToUpper(k), v
			}))), 3)
		t.Logf("DEBUG: %v", got)

		want := []string{
			"FUKUOKA",
			"KAGOSHIMA",
			"KANAZAWA",
		}
		assert.EqualsList(t, got, want)
	})

	t.Run("uniqvalues", func(t *testing.T) {
		got := lists.Uniq(lists.Sort(hashes.Values(daten)))
		t.Logf("DEBUG: %v", got)
		want := []int{1, 2, 3}
		assert.EqualsList(t, got, want)
	})
}

type PointList []int32

func (p PointList) String() string {
	result := "["
	for i, v := range p {
		if i > 0 {
			result += " "
		}
		result += fmt.Sprintf("%d", v)
	}
	result += "]"
	return result
}

func TestListDataType(t *testing.T) {

	var p PointList = []int32{1, 2, 2, -1}

	t.Run("uniq", func(t *testing.T) {
		got := lists.Uniq(p).String()
		want := "[1 2 -1]"
		assert.Equals(t, got, want)
	})

	t.Run("grep", func(t *testing.T) {
		got := lists.Grep(p, func(a int32) bool {
			return a%2 == 0
		}).String()
		want := "[2 2]"
		assert.Equals(t, got, want)
	})

	t.Run("head", func(t *testing.T) {
		got := lists.Head(p, 2).String()
		want := "[1 2]"
		assert.Equals(t, got, want)
	})

	t.Run("sort", func(t *testing.T) {
		got := lists.Sort(p).String()
		want := "[-1 1 2 2]"
		assert.Equals(t, got, want)
	})
}

type PointMap map[int]string

func (p PointMap) String() string {
	result := "["
	first := true
	for i, v := range p {
		if first {
			first = false
		} else {
			result += " "
		}
		result += fmt.Sprintf("%d: %s", i, v)
	}
	result += "]"
	return result
}

func TestMapDataType(t *testing.T) {
	p := PointMap{
		1: "eins",
		2: "zwei",
		4: "vier",
		8: "acht",
	}

	t.Run("grep", func(t *testing.T) {
		got := hashes.Grep(p, func(k int, v string) bool {
			return strings.Contains(v, "i")
		})
		want := PointMap{
			1: "eins",
			2: "zwei",
			4: "vier",
		}
		assert.EqualsMap(t, got, want)
		// compiler checks method
		t.Logf("DEBUG: String: %s ", got.String())
	})

	t.Run("map", func(t *testing.T) {
		got := hashes.Map(p, func(k int, v string) (int, string) {
			return k, strings.ToUpper(v)
		})
		want := PointMap{
			1: "EINS",
			2: "ZWEI",
			4: "VIER",
			8: "ACHT",
		}
		assert.EqualsMap(t, got, want)
		t.Logf("DEBUG: String: %s ", got.String())
	})

	t.Run("sortkeys", func(t *testing.T) {
		got := hashes.SortKeys(p)
		want := []int{1, 2, 4, 8}
		assert.EqualsList(t, got, want)
	})

	t.Run("values", func(t *testing.T) {
		got := lists.Sort(hashes.Values(p))
		want := []string{
			"acht",
			"eins",
			"vier",
			"zwei",
		}
		assert.EqualsList(t, got, want)
	})
}
