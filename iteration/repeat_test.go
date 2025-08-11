package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 9)
	want := "aaaaaaaaa"

	if got != want {
		t.Errorf("%q got %q want", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 100)
	}
}
