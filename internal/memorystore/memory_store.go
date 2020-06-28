package memorystore

import (
	"fmt"

	"github.com/justinhbaker3/todo-list/internal/list"
)

// MemoryStore implements the Store interface
// Stores lists in memory
type MemoryStore struct {
	lists map[string]*list.List
}

func New() *MemoryStore {
	return &MemoryStore{
		lists: map[string]*list.List{},
	}
}

func (m *MemoryStore) Get(title string) (*list.List, error) {
	l, ok := m.lists[title]
	if !ok {
		return nil, fmt.Errorf("list does not exist")
	}
	return l, nil
}

func (m *MemoryStore) Upsert(list *list.List) {
	m.lists[list.Title] = list
}

func (m *MemoryStore) Delete(title string) {
	delete(m.lists, title)
}
