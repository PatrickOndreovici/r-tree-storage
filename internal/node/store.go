package node

type NodeStore interface {
	// Get retrieves the node with the given ID.
	Get(id NodeID) (*Node, error)

	// Put persists a node, creating or overwriting it.
	Put(node *Node) error

	// Delete removes a node permanently.
	Delete(id NodeID) error

	// NewID allocates a fresh, unique NodeID.
	NewID() (NodeID, error)

	// Root returns the NodeID of the current root node.
	// Returns InvalidID when the tree is empty.
	Root() (NodeID, error)

	// SetRoot updates the root pointer.
	SetRoot(id NodeID) error

	// Close flushes and releases all resources.
	Close() error
}
