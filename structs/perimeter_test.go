package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{4.0, 12.0}
	got := Perimeter(rectangle)
	want := 32.0

	if got != want {
		t.Errorf("want %.2f but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{4.0, 12.0}
	got := Area(rectangle)
	want := 48.0

	if got != want {
		t.Errorf("want %.2f but got %.2f", want, got)
	}
}
