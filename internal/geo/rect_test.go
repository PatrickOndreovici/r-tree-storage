package geo

import "testing"

func TestRectArea(t *testing.T) {
	r := Rect{Minx: 0, Miny: 0, Maxx: 2, Maxy: 3}
	expected := 6.0

	if r.Area() != expected {
		t.Errorf("expected area %v, got %v", expected, r.Area())
	}
}

func TestRectIntersects(t *testing.T) {
	r1 := Rect{0, 0, 2, 2}
	r2 := Rect{1, 1, 3, 3}
	r3 := Rect{3, 3, 4, 4}

	if !r1.Intersects(r2) {
		t.Error("expected r1 to intersect r2")
	}
	if r1.Intersects(r3) {
		t.Error("expected r1 NOT to intersect r3")
	}
}

func TestRectUnion(t *testing.T) {
	r1 := Rect{0, 0, 2, 2}
	r2 := Rect{1, 1, 3, 3}

	union := r1.Union(r2)

	expected := Rect{0, 0, 3, 3}
	if union != expected {
		t.Errorf("expected %v, got %v", expected, union)
	}
}

func TestRectEnlargement(t *testing.T) {
	r1 := Rect{0, 0, 2, 2}
	r2 := Rect{1, 1, 3, 3}

	enlargement := r1.Enlargement(r2)
	expected := 5.0 // union area (9) - original (4)

	if enlargement != expected {
		t.Errorf("expected %v, got %v", expected, enlargement)
	}
}
