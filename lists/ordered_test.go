package lists_test

import (
	"testing"

	"github.com/jochenczemmel/assert/assert"
	"github.com/jochenczemmel/perlfunc/lists"
)

func TestMax(t *testing.T) {

	assert.Equal(t, lists.Max([]int{}), 0)
	assert.Equal(t, lists.Max([]int{0, 0, 0}), 0)
	assert.Equal(t, lists.Max([]int{2, 2, 2}), 2)
	assert.Equal(t, lists.Max([]int{3, 2, 9, 5, 6}), 9)
	assert.Equal(t, lists.Max([]int{-3, -2, -9, -5, -6}), -2)
}

func TestMin(t *testing.T) {

	assert.Equal(t, lists.Min([]int{}), 0)
	assert.Equal(t, lists.Min([]int{2}), 2)
	assert.Equal(t, lists.Min([]int{2, 3}), 2)
	assert.Equal(t, lists.Min([]int{3, 2}), 2)
	assert.Equal(t, lists.Min([]int{2, 3, 1}), 1)
	assert.Equal(t, lists.Min([]int{3, 2, 9, 5, 6}), 2)
	assert.Equal(t, lists.Min([]int{-3, -2, -9, -5, -6}), -9)
}

func TestSort(t *testing.T) {
	data := []int{3, 2, 9, 5, 6}
	original := []int{3, 2, 9, 5, 6}
	want := []int{2, 3, 5, 6, 9}
	got := lists.Sort(data)
	assert.EqualList(t, got, want)
	assert.EqualList(t, data, original)
}
