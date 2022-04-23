package lists_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jochenczemmel/assert/assert"
	"github.com/jochenczemmel/perlfunc/lists"
)

func TestPush(t *testing.T) {

	list := []int{}
	lists.Push(&list, 1, 2, 3)
	wantLen := 3
	assert.Equal(t, len(list), wantLen)

	lists.Push(&list, 4)
	wantLen = 4
	assert.Equal(t, len(list), wantLen)
}

func TestPop(t *testing.T) {

	list := []int{1, 2, 3}

	for i, want := range []int{3, 2, 1, 0, 0} {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := lists.Pop(&list)
			assert.Equal(t, got, want)
		})
	}
}

func TestUnshift(t *testing.T) {

	list := []int{}
	lists.Unshift(&list, 1, 2, 3)
	wantLen := 3
	assert.Equal(t, len(list), wantLen)

	lists.Unshift(&list, 4, 5)
	wantLen = 5
	assert.Equal(t, len(list), wantLen)

	for i, want := range []int{3, 2, 1, 5, 4, 0} {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := lists.Pop(&list)
			assert.Equal(t, got, want)
		})
	}
}

func TestShift(t *testing.T) {

	list := []int{1, 2, 3}

	for i, want := range []int{1, 2, 3, 0, 0} {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := lists.Shift(&list)
			assert.Equal(t, got, want)
		})
	}
}

func TestFind(t *testing.T) {

	type MyString []string
	list := MyString{"eins", "zwei", "drei", "vier", "fünf"}
	// list := []string{"eins", "zwei", "drei", "vier", "fünf"}

	candidates := []struct {
		search    string
		want      string
		wantMatch bool
	}{
		{"ei", "eins", true},
		{"dr", "drei", true},
		{"X", "", false},
	}

	for _, c := range candidates {

		t.Run(c.search, func(t *testing.T) {
			got, gotMatch := lists.Find(list, func(value string) bool {
				return strings.Contains(value, c.search)
			})
			assert.Equal(t, got, c.want)
			assert.Equal(t, gotMatch, c.wantMatch)
		})
	}

}

func TestGrep(t *testing.T) {

	list := []string{"eins", "zwei", "drei", "vier", "fünf"}

	filtered := lists.Grep(list, func(value string) bool {
		return strings.Contains(value, "ei")
	})

	for i, want := range []string{"eins", "zwei", "drei", ""} {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := lists.Shift(&filtered)
			assert.Equal(t, got, want)
		})
	}
}

func TestMap(t *testing.T) {

	list := []string{"eins", "zwei", "drei", "vier", "fünf"}
	mapped := lists.Map(list, func(value string) string {
		return strings.ToUpper(value)
	})

	for i, want := range []string{
		"EINS", "ZWEI", "DREI", "VIER", "FÜNF"} {

		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			got := lists.Shift(&mapped)
			assert.Equal(t, got, want)
		})
	}
}

func TestAnyAll(t *testing.T) {

	list := []int{1, 2, 3, 4, 5, 6, 7}
	listObject := lists.New(1, 2, 3, 4, 5, 6, 7)

	candidates := []struct {
		name             string
		fn               func(int) bool
		wantAll, wantAny bool
	}{
		{
			name:    "gt 0",
			fn:      func(v int) bool { return v > 0 },
			wantAll: true,
			wantAny: true,
		},
		{
			name:    "gt 3",
			fn:      func(v int) bool { return v > 3 },
			wantAny: true,
		},
		{
			name:    "gt 6",
			fn:      func(v int) bool { return v > 6 },
			wantAny: true,
		},
		{
			name: "gt 7",
			fn:   func(v int) bool { return v > 7 },
		},
	}

	for _, c := range candidates {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, lists.All(list, c.fn), c.wantAll)
			assert.Equal(t, lists.NotAll(list, c.fn), !c.wantAll)
			assert.Equal(t, lists.Any(list, c.fn), c.wantAny)
			assert.Equal(t, lists.None(list, c.fn), !c.wantAny)
			// OOP: list.go
			assert.Equal(t, listObject.All(c.fn), c.wantAll)
			assert.Equal(t, listObject.NotAll(c.fn), !c.wantAll)
			assert.Equal(t, listObject.Any(c.fn), c.wantAny)
			assert.Equal(t, listObject.None(c.fn), !c.wantAny)
		})
	}
}

func TestReduce(t *testing.T) {

	var liste []int

	summe := lists.Reduce(liste, func(a, b int) int {
		return a + b
	})
	assert.Equal(t, summe, 0)

	liste = []int{}
	summe = lists.Reduce(liste, func(a, b int) int {
		return a + b
	})
	assert.Equal(t, summe, 0)

	liste = []int{3}
	summe = lists.Reduce(liste, func(a, b int) int {
		return a + b
	})
	assert.Equal(t, summe, 3)

	liste = []int{1, 2, 3, 4, 5, 6, 7}

	summe = lists.Reduce(liste, func(a, b int) int {
		return a + b
	})
	assert.Equal(t, summe, 28)

	// OOP: list.go
	listObject := lists.New(1, 2, 3, 4, 5, 6, 7)
	summe = listObject.Reduce(func(a, b int) int {
		return a + b
	})
	assert.Equal(t, summe, 28)
}

func TestFirst(t *testing.T) {

	var liste []int
	got, ok := lists.First(liste, func(value int) bool {
		return value%2 == 0
	})
	assert.Equal(t, got, 0)
	assert.Equal(t, ok, false)

	liste = []int{1, 3, 5, 6, 7, 8}

	got, ok = lists.First(liste, func(value int) bool {
		return value%2 == 0
	})
	assert.Equal(t, got, 6)
	assert.Equal(t, ok, true)

	got, ok = lists.First(liste, func(value int) bool {
		return value > 10
	})
	assert.Equal(t, got, 0)
	assert.Equal(t, ok, false)
}

func TestHead(t *testing.T) {

	liste := []int{1, 2, 3, 4, 5}
	assert.EqualList(t, lists.Head(liste, 3), []int{1, 2, 3})
	assert.EqualList(t, lists.Head(liste, 2), []int{1, 2})
	assert.EqualList(t, lists.Head(liste, 1), []int{1})
	assert.EqualList(t, lists.Head(liste, 0), []int{})
	assert.EqualList(t, lists.Head(liste, -1), []int{})
	assert.EqualList(t, lists.Head(liste, 5), []int{1, 2, 3, 4, 5})
	assert.EqualList(t, lists.Head(liste, 6), []int{1, 2, 3, 4, 5})
	assert.EqualList(t, lists.Head(liste, 9), []int{1, 2, 3, 4, 5})
}

func TestTail(t *testing.T) {

	liste := []int{1, 2, 3, 4, 5}
	assert.EqualList(t, lists.Tail(liste, 3), []int{3, 4, 5})
	assert.EqualList(t, lists.Tail(liste, 2), []int{4, 5})
	assert.EqualList(t, lists.Tail(liste, 1), []int{5})
	assert.EqualList(t, lists.Tail(liste, 0), []int{})
	assert.EqualList(t, lists.Tail(liste, -1), []int{})
	assert.EqualList(t, lists.Tail(liste, 5), []int{1, 2, 3, 4, 5})
	assert.EqualList(t, lists.Tail(liste, 6), []int{1, 2, 3, 4, 5})
	assert.EqualList(t, lists.Tail(liste, 9), []int{1, 2, 3, 4, 5})
}
