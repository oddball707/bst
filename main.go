package main

import (
	s "github.com/oddball707/bst/service"
	"fmt"
	"os"
)

func main() {
	testData := []int{20, 4, 8, 21, 3, 7, 10}

	tree, err := s.NewTree(testData)
	if err != nil {
		fmt.Println("Error creating tree")
		os.Exit(1)
	}
	deepest, depth := tree.DeepestNodes()

	fmt.Printf("depth: %d\n", depth)
	fmt.Printf("deepest: %d\n", deepest)
}