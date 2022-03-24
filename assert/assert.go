package assert

import "testing"

// Equals checks if two values are equal.
func Equals[C comparable](t *testing.T, got, want C) {
	t.Helper()
	if got != want {
		t.Errorf("ERROR: got %v, want %v", got, want)
	}
}

// NotEquals checks if two values are not equal.
func NotEquals[C comparable](t *testing.T, got, want C) {
	t.Helper()
	if got == want {
		t.Errorf("ERROR: got %v, want %v", got, want)
	}
}

// EqualsList checks if two lists are equal.
// If the lenght is different, the elements are not compared.
func EqualsList[L ~[]C, C comparable](t *testing.T, got, want L) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("ERROR: length: got %d, want %d", len(got), len(want))
	}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("ERROR: %d: got %v, want %v", i, got[i], w)
		}
	}
}

// EqualsMap checks if two maps are equal.
// If the lenght is different, the elements are not compared.
func EqualsMap[M ~map[K]V, K, V comparable](t *testing.T,
	got, want M) {

	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("ERROR: length: got %d, want %d", len(got), len(want))
	}
	for k, w := range want {
		if got[k] != w {
			t.Errorf("ERROR: %v: got %v, want %v", k, got[k], w)
		}
	}
}
