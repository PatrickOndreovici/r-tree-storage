package geo

import "math"

type Rect struct {
	Minx, Miny float64
	Maxx, Maxy float64
}

func (r Rect) Area() float64 {
	return (r.Maxx - r.Minx) * (r.Maxy - r.Miny)
}

func (r Rect) Intersects(r2 Rect) bool {
	if r.Maxx < r2.Minx || r.Minx > r2.Maxx {
		return false
	}
	if r.Maxy < r2.Miny || r.Miny > r2.Maxy {
		return false
	}
	return true
}

func (r Rect) Union(r2 Rect) Rect {
	return Rect{
		Minx: math.Min(r.Minx, r2.Minx),
		Miny: math.Min(r.Miny, r2.Miny),
		Maxx: math.Max(r.Maxx, r2.Maxx),
		Maxy: math.Max(r.Maxy, r2.Maxy),
	}
}

func (r Rect) Enlargement(r2 Rect) float64 {
	return r.Union(r2).Area() - r.Area()
}

func (r Rect) MaxX() float64 { return r.Maxx }
func (r Rect) MinX() float64 { return r.Minx }
func (r Rect) MaxY() float64 { return r.Maxy }
func (r Rect) MinY() float64 { return r.Miny }
