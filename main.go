package main

import (
	"github.com/go-btree/btree"
)

func main() {
	tree := btree.CreateBTree(2)

	data := []int{5, 3, 6, 4, 6, 10, 34, 1, 2}

	tree.Add(data[0])
	btree.PrintBTree(tree)
}
