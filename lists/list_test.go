package lists_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jochenczemmel/perlfunc/assert"
	"github.com/jochenczemmel/perlfunc/lists"
)

func TestListInit(t *testing.T) {

	got := lists.New[int]()
	wantLen := 0
	assert.Equals(t, got.Len(), wantLen)

	got = lists.New(1, 2, 3)
	wantLen = 3
	assert.Equals(t, got.Len(), wantLen)
}

func TestListPush(t *testing.T) {

	list := lists.New[int]()
	list.Push(1, 2, 3)
	wantLen := 3
	assert.Equals(t, list.Len(), wantLen)

	list.Push(4)
	wantLen = 4
	assert.Equals(t, list.Len(), wantLen)
}

func TestListUnshift(t *testing.T) {

	list := lists.New[int]()
	list.Unshift(1, 2, 3)
	wantLen := 3
	assert.Equals(t, list.Len(), wantLen)

	list.Unshift(4, 5)
	wantLen = 5
	assert.Equals(t, list.Len(), wantLen)

	for i, want := range []int{3, 2, 1, 5, 4, 0} {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := list.Pop()
			assert.Equals(t, got, want)
		})
	}
}

func TestListPop(t *testing.T) {

	list := lists.New(1, 2, 3)

	for i, want := range []int{3, 2, 1, 0, 0} {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := list.Pop()
			assert.Equals(t, got, want)
		})
	}
}

func TestListShift(t *testing.T) {

	list := lists.New(1, 2, 3)

	for i, want := range []int{1, 2, 3, 0, 0} {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := list.Shift()
			assert.Equals(t, got, want)
		})
	}
}

func TestListGrep(t *testing.T) {

	list := lists.New("eins", "zwei", "drei", "vier", "fünf")
	filtered := list.Grep(func(value string) bool {
		return strings.Contains(value, "ei")
	})

	for i, want := range []string{"eins", "zwei", "drei", ""} {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := filtered.Shift()
			assert.Equals(t, got, want)
		})
	}
}

func TestListMap(t *testing.T) {

	list := lists.New("eins", "zwei", "drei", "vier", "fünf")
	mapped := list.Map(func(value string) string {
		return strings.ToUpper(value)
	})

	for i, want := range []string{
		"EINS", "ZWEI", "DREI", "VIER", "FÜNF"} {

		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := mapped.Shift()
			assert.Equals(t, got, want)
		})
	}
}
