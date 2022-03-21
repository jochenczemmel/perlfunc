package assert

import "testing"

func Equals[C comparable](t *testing.T, got, want C) {
	t.Helper()
	if got != want {
		t.Errorf("ERROR: got %v, want %v", got, want)
	}
}

func NotEquals[C comparable](t *testing.T, got, want C) {
	t.Helper()
	if got == want {
		t.Errorf("ERROR: got %v, want %v", got, want)
	}
}

func EqualsList[L ~[]C, C comparable](t *testing.T, got, want L) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("ERROR: length: got %d, want %d", len(got), len(want))
	}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("ERROR: got %v, want %v", got[i], w)
		}
	}
}
