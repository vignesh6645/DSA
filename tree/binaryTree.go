package tree

import (
	"cmp"
	"fmt"
	"math/rand"
	"time"
)

// Node represents a single Node in the BST
type Node[T cmp.Ordered] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

// BST represents the binary search tree
type BST[T cmp.Ordered] struct {
	Root *Node[T]
}

func InitBTree() {
	bst := BST[int]{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		num := r.Intn(100)
		bst.Insert(num)
	}
	bst.InOrder()
}

func (b *BST[T]) Insert(val T) {
	b.Root = insertNode(b.Root, val)
}

func insertNode[T cmp.Ordered](node *Node[T], val T) *Node[T] {
	if node == nil {
		return &Node[T]{value: val}
	}

	if val < node.value {
		node.left = insertNode(node.left, val)
	} else if val > node.value {
		node.right = insertNode(node.right, val)
	}
	return node
}

/* Traversal */

func (b *BST[T]) InOrder() {
	fmt.Println("In-Order: ")
	inOrder(b.Root)
	fmt.Println()
}

func inOrder[T cmp.Ordered](node *Node[T]) {
	if node == nil {
		return
	}
	inOrder(node.left)
	fmt.Printf("%v -> ", node.value)
	inOrder(node.right)
}
