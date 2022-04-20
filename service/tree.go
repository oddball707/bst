package service

import (
	"errors"
	"fmt"
)

type Tree interface {
	DeepestNodes() (deepest []int, depth int)
}

type BstTree struct {
	root *BstNode
}

func NewTree(input []int) (*BstTree, error) {
	if len(input) < 1 {
		return nil, errors.New("Tree must have at least one element")
	}
	root := NewNode(input[0])
	for i := 1; i < len(input); i++ {
		fmt.Printf("Adding %v...\n", input[i])
		root.AddChild(input[i]);
	}

	return &BstTree{
		root: root,
	}, nil
}

func (t *BstTree) DeepestNodes() (deepest []int, depth int) {
	depth = 0
	return t.root.DeepestNodes(depth)
}