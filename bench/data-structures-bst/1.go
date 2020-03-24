package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// Node a single node that composes the tree
type Node struct {
	value int
	left  *Node //left
	right *Node //right
}

// ItemBinarySearchTree the binary search tree of Items
type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

// Insert inserts the Item t in the tree
func (bst *ItemBinarySearchTree) Insert(value int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &Node{value, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

// internal function to find the correct place for a node in a tree
func insertNode(node, newNode *Node) {
	if newNode.value < node.value {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// InOrderTraverse visits all nodes with in-order traversing
func (bst *ItemBinarySearchTree) InOrderTraverse(f func(int)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.root, f)
}

// internal recursive function to traverse in order
func inOrderTraverse(n *Node, f func(int)) {
	if n != nil {
		inOrderTraverse(n.left, f)
		f(n.value)
		inOrderTraverse(n.right, f)
	}
}

// PreOrderTraverse visits all nodes with pre-order traversing
func (bst *ItemBinarySearchTree) PreOrderTraverse(f func(int)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrderTraverse(bst.root, f)
}

// internal recursive function to traverse pre order
func preOrderTraverse(n *Node, f func(int)) {
	if n != nil {
		f(n.value)
		preOrderTraverse(n.left, f)
		preOrderTraverse(n.right, f)
	}
}

// PostOrderTraverse visits all nodes with post-order traversing
func (bst *ItemBinarySearchTree) PostOrderTraverse(f func(int)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	postOrderTraverse(bst.root, f)
}

// internal recursive function to traverse post order
func postOrderTraverse(n *Node, f func(int)) {
	if n != nil {
		postOrderTraverse(n.left, f)
		postOrderTraverse(n.right, f)
		f(n.value)
	}
}

// Min returns the Item with min value stored in the tree
func (bst *ItemBinarySearchTree) Min() *int {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	n := bst.root
	if n == nil {
		return nil
	}
	for {
		if n.left == nil {
			return &n.value
		}
		n = n.left
	}
}

// Max returns the Item with max value stored in the tree
func (bst *ItemBinarySearchTree) Max() *int {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	n := bst.root
	if n == nil {
		return nil
	}
	for {
		if n.right == nil {
			return &n.value
		}
		n = n.right
	}
}

// Search returns true if the Item t exists in the tree
func (bst *ItemBinarySearchTree) Search(value int) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return search(bst.root, value)
}

// internal recursive function to search an item in the tree
func search(n *Node, value int) bool {
	if n == nil {
		return false
	}
	if value < n.value {
		return search(n.left, value)
	}
	if value > n.value {
		return search(n.right, value)
	}
	return true
}

// Remove removes the Item with vakue `value` from the tree
func (bst *ItemBinarySearchTree) Remove(value int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	remove(bst.root, value)
}

// internal recursive function to remove an item
func remove(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.value {
		node.left = remove(node.left, value)
		return node
	}
	if value > node.value {
		node.right = remove(node.right, value)
		return node
	}
	// value == node.value
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	leftmostrightside := node.right
	for {
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	node.value = leftmostrightside.value
	node.right = remove(node.right, node.value)
	return node
}

// String prints a visual representation of the tree
func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.value)
		stringify(n.right, level)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var bst ItemBinarySearchTree

	var n = 1000000
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	for i := 0; i < n; i++ {
		bst.Insert(rand.Int())
	}
}
