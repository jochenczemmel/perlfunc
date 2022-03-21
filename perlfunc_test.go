package main

import (
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

	t.Run("uniqvalues", func(t *testing.T) {
		got := lists.Uniq(lists.Sort(hashes.Values(daten)))
		t.Logf("DEBUG: %v", got)
		want := []int{1, 2, 3}
		assert.EqualsList(t, got, want)
	})
}
