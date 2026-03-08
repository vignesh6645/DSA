package tree

import (
	"cmp"
	"fmt"
)

// Node represents a single Node in the BST
type Node[T cmp.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// BST represents the binary search tree
type BST[T cmp.Ordered] struct {
	Root *Node[T]
}

func InitBTree() {
	bst := BST[int]{}

	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := []int{1, 10, 2, 5, 100, 10, 10, 25, 3, 7, 200, 4, 0, 9}
	for i := 0; i < len(nums)-1; i++ {
		//num := r.Intn(100)
		bst.Insert(nums[i])
	}
	bst.InOrder()
	fmt.Println("IsPresent: ", bst.Search(200))
	bst.Delete(5)

	bst.InOrder()
}

func (b *BST[T]) Insert(val T) {
	b.Root = insertNode(b.Root, val)
}

func insertNode[T cmp.Ordered](node *Node[T], val T) *Node[T] {
	if node == nil {
		return &Node[T]{Value: val}
	}

	if val < node.Value {
		node.Left = insertNode(node.Left, val)
	} else if val > node.Value {
		node.Right = insertNode(node.Right, val)
	}
	return node
}

func (b BST[T]) Search(val T) bool {
	return searchNode(b.Root, val)
}

func searchNode[T cmp.Ordered](node *Node[T], val T) bool {
	if node == nil {
		return false
	}
	if node.Value == val {
		return true
	}

	if val < node.Value {
		return searchNode(node.Left, val)
	}
	return searchNode(node.Right, val)
}

func (b *BST[T]) Delete(val T) {
	b.Root = deleteNode(b.Root, val)
}

func deleteNode[T cmp.Ordered](node *Node[T], val T) *Node[T] {
	if node == nil {
		return nil
	}

	if val < node.Value {
		node.Left = deleteNode(node.Left, val)
	} else if val > node.Value {
		node.Right = deleteNode(node.Right, val)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		}

		if node.Left == nil {
			return node.Right
		}

		if node.Right == nil {
			return node.Left
		}

		successor := minNode(node.Right)
		node.Value = successor.Value
		node.Right = deleteNode(node.Right, successor.Value)

	}
	return node
}

func minNode[T cmp.Ordered](node *Node[T]) *Node[T] {
	for node.Left != nil {
		node = node.Left
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
	inOrder(node.Left)
	fmt.Printf("%v > ", node.Value)
	inOrder(node.Right)
}
