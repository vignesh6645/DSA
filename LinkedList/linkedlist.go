package linkedlist

// Node represents a node in the linked list
type Node struct {
	Value interface{}
	Next  *Node
}

// LinkedList defines the interface for a linked list data structure
type LinkedList interface {
	// Insert adds a new node with the given value at the end of the list
	Insert(value interface{})
	
	// InsertAt adds a new node with the given value at the specified index
	InsertAt(index int, value interface{}) error
	
	// Delete removes the first node with the given value
	Delete(value interface{}) bool
	
	// DeleteAt removes the node at the specified index
	DeleteAt(index int) error
	
	// Get returns the value at the specified index
	Get(index int) (interface{}, error)
	
	// Size returns the number of nodes in the list
	Size() int
	
	// IsEmpty returns true if the list is empty
	IsEmpty() bool
	
	// Clear removes all nodes from the list
	Clear()
	
	// ToSlice converts the linked list to a slice
	ToSlice() []interface{}
}

