package structsmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(3.0, 5.0)
	want := 16.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
