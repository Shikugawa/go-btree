package btree

type BTree struct {
	Container *Container
	K         int
}

func (t *BTree) Add(n int) {
	block := &Block{
		Value: n,
	}
	t.Container.AddNode(block)
}

func CreateBTree(size int) *BTree {
	tree := &BTree{
		Container: InitContainer(2 * size),
	}

	return tree
}

func PrintBTree(t *BTree) {
	PrintContainer(t.Container, 1)
}
