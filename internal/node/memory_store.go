package node

import (
	"errors"
	"sync"
)

type MemoryStore struct {
	mu     sync.RWMutex
	nodes  map[NodeID]*Node
	nextID NodeID
	root   NodeID
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		nodes:  make(map[NodeID]*Node),
		nextID: 1,
		root:   0,
	}
}

func (m *MemoryStore) Get(id NodeID) (*Node, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	node, ok := m.nodes[id]
	if !ok {
		return nil, errors.New("node not found")
	}
	return node, nil
}

func (m *MemoryStore) Put(node *Node) error {
	if node == nil {
		return errors.New("node is nil")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.nodes[node.ID] = node
	return nil
}

func (m *MemoryStore) Delete(id NodeID) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.nodes[id]; !ok {
		return errors.New("node not found")
	}

	delete(m.nodes, id)
	return nil
}

func (m *MemoryStore) NewID() (NodeID, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id := m.nextID
	m.nextID++
	return id, nil
}

func (m *MemoryStore) Root() (NodeID, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.root == 0 {
		return 0, errors.New("root not set")
	}
	return m.root, nil
}

func (m *MemoryStore) SetRoot(id NodeID) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.nodes[id]; !ok {
		return errors.New("root node does not exist")
	}

	m.root = id
	return nil
}

func (m *MemoryStore) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.nodes = make(map[NodeID]*Node)
	m.root = 0
	m.nextID = 1

	return nil
}

func (m *MemoryStore) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.nodes)
}
