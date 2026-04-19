package node

import "r-tree/internal/geo"

type NodeID uint64

type Node struct {
	ID       NodeID
	IsLeaf   bool
	ParentID NodeID
	Entries  []Entry
}

type Entry struct {
	Rectangle geo.Rect
	ChildID   NodeID
	Data      []byte
}
