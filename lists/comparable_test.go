package lists_test

import (
	"testing"

	"github.com/jochenczemmel/perlfunc/assert"
	"github.com/jochenczemmel/perlfunc/lists"
)

func TestHashesUniq(t *testing.T) {
	data := []string{
		"Osaka",
		"Fukuoka",
		"Kyoto",
		"Osaka",
		"Fukuoka",
		"Kyoto",
		"Osaka",
		"Fukuoka",
		"Fukuoka",
	}

	got := lists.Uniq(lists.Sort(data))
	t.Logf("DEBUG: %v", got)
	want := []string{
		"Fukuoka",
		"Kyoto",
		"Osaka",
	}
	assert.EqualsList(t, got, want)
}
