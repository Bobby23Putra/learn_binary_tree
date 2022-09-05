package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(keyValue int) {
	if n.Key <= keyValue {
		if n.right == nil {
			n.right = &Node{Key: keyValue}
			fmt.Println("right " + strconv.Itoa(keyValue))
		} else {
			n.right.Insert(keyValue)
		}
	} else {
		if n.left == nil {
			n.left = &Node{Key: keyValue}
			fmt.Println("left " + strconv.Itoa(keyValue))
		} else {
			n.left.Insert(keyValue)
		}
	}
}

func (n *Node) Search(keyValue int) bool {
	if n == nil {
		return false
	}

	if n.Key < keyValue {
		return n.right.Search(keyValue)
	} else if n.Key > keyValue {
		return n.left.Search(keyValue)
	}

	return true
}

func (n *Node) MinKey() int {
	if n == nil {
		return 0
	}

	if n.left != nil {
		return n.left.MinKey()
	}

	return n.Key

}

func (n *Node) MaxKey() int {
	if n == nil {
		return 0
	}

	if n.right != nil {
		return n.right.MaxKey()
	}

	return n.Key

}

func main() {
	tree := &Node{Key: 150}
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(100)
	tree.Insert(50)
	fmt.Println(tree.MinKey())
	fmt.Println(tree.MaxKey())
}
