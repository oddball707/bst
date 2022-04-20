package service

import (
	"fmt"
)

type Node interface {
	getValue() int
	getRight() *BstNode
	getLeft() *BstNode
	AddChild(value int) error
	DeepestNodes(depth int) (deepest []int, finalDepth int)
}

type BstNode struct {
	value int
	rightChild *BstNode
	leftChild *BstNode
}

func NewNode(value int) *BstNode {
	return &BstNode{
		value: 		value, 
		rightChild: nil, 
		leftChild:  nil,
	}
}

func (n *BstNode) getValue() int {
	return n.value
}

func (n *BstNode) getRight() *BstNode {
	if n.rightChild != nil {
		return n.rightChild
	} else {
		return nil
	}
}

func (n *BstNode) getLeft() *BstNode {
	if n.rightChild != nil {
		return n.leftChild
	} else {
		return nil
	}
}

func (n *BstNode) AddChild(value int) error {
	if value > n.getValue() {
		fmt.Printf("   Adding right\n")
		if n.rightChild == nil {
			n.rightChild = NewNode(value)
		} else {
			n.rightChild.AddChild(value)
		}
	} else {
		fmt.Printf("   Adding left\n")	
		if n.leftChild == nil {
			n.leftChild = NewNode(value)
		} else {
			n.leftChild.AddChild(value)
		}
	}

	return nil
}

func (n *BstNode) DeepestNodes(depth int) (deepest []int, finalDepth int) {
	var rightDeepest, leftDeepest []int
	var rightDepth, leftDepth int
	
	right := n.getRight()
	left := n.getLeft()
	if(right == nil && left == nil) {
		return []int{n.getValue()}, depth
	}
	if(right != nil) {
		rightDeepest, rightDepth = right.DeepestNodes(depth+1)
	}
	if(left != nil) {
		leftDeepest, leftDepth = left.DeepestNodes(depth+1)
	}
	if(rightDepth > leftDepth) {
		return rightDeepest, rightDepth
	} else if(leftDepth > rightDepth) {
		return leftDeepest, leftDepth
	}

	return append(rightDeepest, leftDeepest...), leftDepth
}