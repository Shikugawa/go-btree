package btree

import "fmt"

func swapBlock(first int, second int, c *Container) {
	c.Blocks[first], c.Blocks[second] = c.Blocks[second], c.Blocks[first]
}

func PrintContainer(c *Container, depth int) {
	fmt.Println(depth)

	for i := 0; i < c.CurrentSize; i++ {
		fmt.Print(string(c.Blocks[i].Value) + " ")
	}

	fmt.Println()

	if !c.isLeaf() {
		for i := 0; i < c.CurrentSize; i++ {
			PrintContainer(c.Blocks[i].Left, depth+1)
			PrintContainer(c.Blocks[i].Right, depth+1)
		}
	}
}
