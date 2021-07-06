package main

import (
	"log"
)

// A Binary Tree Node
type Node struct {
	data  int
	left  *Node
	right *Node
}

//new node
func newNode(data int) (temp *Node) {
	temp = &Node{}
	temp.data = data
	temp.left = nil
	temp.right = nil
	return
}

// leaf node print
func printLeaf(root *Node) {
	// base condition
	if root == nil {
		return
	}

	// check node is leaf or not
	if root.left == nil && root.right == nil {
		log.Print(root.data)
		return
	}
	// check left is nil
	if root.left != nil {
		printLeaf(root.left)
	}
	// check right is nil
	if root.right != nil {
		printLeaf(root.right)
	}
}


// bottom view of tree
func printBottomView(root *Node) {
  // base condition
	if root == nil {
		return
	}
}


func main() {
	root := newNode(1)
	root.left = newNode(2)
	root.right = newNode(3)
	root.left.left = newNode(4)

	printLeaf(root)
}

