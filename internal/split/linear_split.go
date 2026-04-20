package split

import (
	"math"
	"r-tree/internal/node"
)

type LinearSplit struct{}

func (l LinearSplit) Split(entries []node.Entry, minEntries int) ([]node.Entry, []node.Entry) {
	seed1, seed2 := l.pickSeeds(entries)

	group1 := make([]node.Entry, 1)
	group2 := make([]node.Entry, 1)

	group1[0] = entries[seed1]
	group2[0] = entries[seed2]

	group1Rectangle := group1[0].Rectangle
	group2Rectangle := group2[0].Rectangle

	if seed1 > seed2 {
		entries = append(entries[:seed1], entries[seed1+1:]...)
		entries = append(entries[:seed2], entries[seed2+1:]...)
	} else {
		entries = append(entries[:seed2], entries[seed2+1:]...)
		entries = append(entries[:seed1], entries[seed1+1:]...)
	}

	for i := 0; i < len(entries); i++ {

		if len(group1)+len(entries)-i == minEntries {
			group1 = append(group1, entries[i:]...)
			break
		}
		if len(group2)+len(entries)-i == minEntries {
			group2 = append(group2, entries[i:]...)
			break
		}
		
		group1Enlargement := entries[i].Rectangle.Enlargement(group1Rectangle)
		group2Enlargement := entries[i].Rectangle.Enlargement(group2Rectangle)

		if group1Enlargement > group2Enlargement {
			group2 = append(group2, entries[i])
			group2Rectangle = group2Rectangle.Union(entries[i].Rectangle)
		} else if group1Enlargement < group2Enlargement {
			group1 = append(group1, entries[i])
			group1Rectangle = group1Rectangle.Union(entries[i].Rectangle)
		} else {
			if len(group1) < len(group2) {
				group1 = append(group1, entries[i])
				group1Rectangle = group1Rectangle.Union(entries[i].Rectangle)
			} else {
				group2 = append(group2, entries[i])
				group2Rectangle = group2Rectangle.Union(entries[i].Rectangle)
			}
		}
	}

	return group1, group2
}

func (l LinearSplit) pickSeeds(entries []node.Entry) (int, int) {
	xMaxLowIdx, xMinHighIdx := 0, 0
	xMaxLow := entries[0].Rectangle.Minx
	xMinHigh := entries[0].Rectangle.Maxx
	xAbsMin := entries[0].Rectangle.Minx
	xAbsMax := entries[0].Rectangle.Maxx

	yMaxLowIdx, yMinHighIdx := 0, 0
	yMaxLow := entries[0].Rectangle.Miny
	yMinHigh := entries[0].Rectangle.Maxy
	yAbsMin := entries[0].Rectangle.Miny
	yAbsMax := entries[0].Rectangle.Maxy

	for i := 1; i < len(entries); i++ {
		r := entries[i].Rectangle

		if r.Minx > xMaxLow {
			xMaxLow = r.Minx
			xMaxLowIdx = i
		}
		if r.Maxx < xMinHigh {
			xMinHigh = r.Maxx
			xMinHighIdx = i
		}
		if r.Minx < xAbsMin {
			xAbsMin = r.Minx
		}
		if r.Maxx > xAbsMax {
			xAbsMax = r.Maxx
		}

		if r.Miny > yMaxLow {
			yMaxLow = r.Miny
			yMaxLowIdx = i
		}
		if r.Maxy < yMinHigh {
			yMinHigh = r.Maxy
			yMinHighIdx = i
		}
		if r.Miny < yAbsMin {
			yAbsMin = r.Miny
		}
		if r.Maxy > yAbsMax {
			yAbsMax = r.Maxy
		}
	}

	// Normalize by each axis's total span so a large coordinate range
	// doesn't make one axis win purely due to scale.
	seed1, seed2 := xMinHighIdx, xMaxLowIdx

	xSpan := xAbsMax - xAbsMin
	ySpan := yAbsMax - yAbsMin

	xSep, ySep := 0.0, 0.0
	if xSpan > 0 {
		xSep = math.Abs(xMaxLow-xMinHigh) / xSpan
	}
	if ySpan > 0 {
		ySep = math.Abs(yMaxLow-yMinHigh) / ySpan
	}

	if ySep > xSep {
		seed1, seed2 = yMinHighIdx, yMaxLowIdx
	}

	// Ensure seeds are distinct
	if seed1 == seed2 {
		if seed2 == 0 {
			seed2 = 1
		} else {
			seed1 = 0
		}
	}

	return seed1, seed2
}
