package split

import "r-tree/internal/node"

type Splitter interface {
	Split(entries []node.Entry, minEntries int) ([]node.Entry, []node.Entry)
}
