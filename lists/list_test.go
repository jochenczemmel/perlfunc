package lists_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jochenczemmel/perlfunc/assert"
	"github.com/jochenczemmel/perlfunc/lists"
)

func TestListPush(t *testing.T) {

	list := []int{}
	lists.Push(&list, 1, 2, 3)
	wantLen := 3
	assert.Equals(t, len(list), wantLen)

	lists.Push(&list, 4)
	wantLen = 4
	assert.Equals(t, len(list), wantLen)
}

func TestListPop(t *testing.T) {

	list := []int{1, 2, 3}

	for i, want := range []int{3, 2, 1, 0, 0} {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := lists.Pop(&list)
			assert.Equals(t, got, want)
		})
	}
}

func TestListUnshift(t *testing.T) {

	list := []int{}
	lists.Unshift(&list, 1, 2, 3)
	wantLen := 3
	assert.Equals(t, len(list), wantLen)

	lists.Unshift(&list, 4, 5)
	wantLen = 5
	assert.Equals(t, len(list), wantLen)

	for i, want := range []int{3, 2, 1, 5, 4, 0} {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := lists.Pop(&list)
			assert.Equals(t, got, want)
		})
	}
}

func TestListShift(t *testing.T) {

	list := []int{1, 2, 3}

	for i, want := range []int{1, 2, 3, 0, 0} {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := lists.Shift(&list)
			assert.Equals(t, got, want)
		})
	}
}

func TestListGrep(t *testing.T) {

	list := []string{"eins", "zwei", "drei", "vier", "fünf"}

	filtered := lists.Grep(list, func(value string) bool {
		return strings.Contains(value, "ei")
	})

	for i, want := range []string{"eins", "zwei", "drei", ""} {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := lists.Shift(&filtered)
			assert.Equals(t, got, want)
		})
	}
}

func TestListMap(t *testing.T) {

	list := []string{"eins", "zwei", "drei", "vier", "fünf"}
	mapped := lists.Map(list, func(value string) string {
		return strings.ToUpper(value)
	})

	for i, want := range []string{
		"EINS", "ZWEI", "DREI", "VIER", "FÜNF"} {

		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := lists.Shift(&mapped)
			assert.Equals(t, got, want)
		})
	}
}
