package split

import (
	"r-tree/internal/geo"
	"r-tree/internal/node"
	"testing"
)

func TestLinearSplit_Basic(t *testing.T) {
	entries := []node.Entry{
		{Rectangle: geo.Rect{Minx: 0, Miny: 0, Maxx: 1, Maxy: 1}},
		{Rectangle: geo.Rect{Minx: 10, Miny: 10, Maxx: 11, Maxy: 11}},
		{Rectangle: geo.Rect{Minx: 0.5, Miny: 0.5, Maxx: 1.5, Maxy: 1.5}},
		{Rectangle: geo.Rect{Minx: 9, Miny: 9, Maxx: 10, Maxy: 10}},
	}

	s := LinearSplit{}
	g1, g2 := s.Split(entries, 1)

	if len(g1) == 0 || len(g2) == 0 {
		t.Fatal("expected non-empty groups")
	}
	if len(g1)+len(g2) != len(entries) {
		t.Fatalf("expected %d total entries, got %d", len(entries), len(g1)+len(g2))
	}
}

func TestLinearSplit_MinEntries(t *testing.T) {
	entries := []node.Entry{
		{Rectangle: geo.Rect{Minx: 0, Miny: 0, Maxx: 1, Maxy: 1}},
		{Rectangle: geo.Rect{Minx: 1, Miny: 1, Maxx: 2, Maxy: 2}},
		{Rectangle: geo.Rect{Minx: 2, Miny: 2, Maxx: 3, Maxy: 3}},
		{Rectangle: geo.Rect{Minx: 3, Miny: 3, Maxx: 4, Maxy: 4}},
		{Rectangle: geo.Rect{Minx: 4, Miny: 4, Maxx: 5, Maxy: 5}},
	}

	minEntries := 2
	s := LinearSplit{}
	g1, g2 := s.Split(entries, minEntries)

	if len(g1) < minEntries {
		t.Errorf("group1 has %d entries, want at least %d", len(g1), minEntries)
	}
	if len(g2) < minEntries {
		t.Errorf("group2 has %d entries, want at least %d", len(g2), minEntries)
	}
}

func TestLinearSplit_AllEntriesDistributed(t *testing.T) {
	entries := []node.Entry{
		{Rectangle: geo.Rect{Minx: 0, Miny: 0, Maxx: 2, Maxy: 2}},
		{Rectangle: geo.Rect{Minx: 1, Miny: 1, Maxx: 3, Maxy: 3}},
		{Rectangle: geo.Rect{Minx: 5, Miny: 5, Maxx: 7, Maxy: 7}},
		{Rectangle: geo.Rect{Minx: 6, Miny: 6, Maxx: 8, Maxy: 8}},
		{Rectangle: geo.Rect{Minx: 3, Miny: 3, Maxx: 4, Maxy: 4}},
	}

	s := LinearSplit{}
	g1, g2 := s.Split(entries, 2)

	if len(g1)+len(g2) != len(entries) {
		t.Fatalf("lost entries during split: got %d, want %d", len(g1)+len(g2), len(entries))
	}
}
