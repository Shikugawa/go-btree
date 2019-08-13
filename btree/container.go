package btree

type Block struct {
	Value            int
	Index            int
	CurrentContainer *Container
	Left             *Container
	Right            *Container
}

type Container struct {
	MaxSize     int
	CurrentSize int
	Blocks      []*Block
	Parent      *Block
}

func InitContainer(maxSize int) *Container {
	container := &Container{
		MaxSize: maxSize,
	}
	container.Blocks = make([]*Block, container.MaxSize)
	return container
}

func (c *Container) createRightContainer() *Container {
	container := InitContainer(c.MaxSize)

	for i := c.MaxSize/2 + 1; i < c.MaxSize; i++ {
		container.Blocks[i-c.MaxSize/2+1] = &Block{
			Value:            c.Blocks[i].Value,
			Index:            i - c.MaxSize/2 + 1,
			CurrentContainer: container,
			Left:             nil,
			Right:            nil,
		}
	}

	return container
}

func (c *Container) createLeftContainer() *Container {
	container := InitContainer(c.MaxSize)

	for i := 0; i < c.MaxSize/2; i++ {
		container.Blocks[i] = &Block{
			Value:            c.Blocks[i].Value,
			Index:            i,
			CurrentContainer: container,
			Left:             nil,
			Right:            nil,
		}
	}

	return container
}

func (c *Container) constructTree(b *Block) {
	mid := c.Blocks[c.MaxSize/2]
	leftContainer := c.createLeftContainer()
	rightContainer := c.createLeftContainer()

	if c.isRootContainer() {
		parentContainer := InitContainer(c.MaxSize)
		parentContainer.AddNode(&Block{
			Value:            mid.Value,
			Index:            -1,
			CurrentContainer: parentContainer,
			Left:             leftContainer,
			Right:            rightContainer,
		})
	} else {
		parentContainer := c.Parent.CurrentContainer
		blk := &Block{
			Value:            mid.Value,
			Index:            -1,
			CurrentContainer: parentContainer,
			Left:             nil,
			Right:            nil,
		}
		parentContainer.AddNode(blk)
		parentContainer.Blocks[blk.Index].Right = rightContainer
		parentContainer.Blocks[blk.Index].Left = leftContainer
		parentContainer.Blocks[blk.Index-1].Right = leftContainer
	}
}

func (c *Container) AddNode(b *Block) {
	if c.CurrentSize == c.MaxSize {
		c.constructTree(b)
		return
	}

	if c.isLeaf() {
		c.addLeaf(b)
	} else {
		c.addMiddle(b)
	}
}

func (c *Container) addMiddle(b *Block) {
	if c.CurrentSize == 0 {
		panic("Middle Container should have more than 1 value")
	}

	if c.CurrentSize == 1 {
		if c.Blocks[0].Value < b.Value {
			c.Blocks[0].Right.AddNode(b)
		} else {
			c.Blocks[0].Left.AddNode(b)
		}
		return
	}

	if c.Blocks[c.CurrentSize-1].Value < b.Value {
		c.Blocks[c.CurrentSize-1].Right.AddNode(b)
		return
	}

	for i := 0; i < c.CurrentSize-1; i++ {
		if c.Blocks[i].Value < b.Value && c.Blocks[i+1].Value > b.Value {
			c.Blocks[i].Right.AddNode(b)
			return
		}
	}
}

func (c *Container) addLeaf(b *Block) {
	defer func() {
		c.CurrentSize++
	}()

	b.CurrentContainer = c

	if c.CurrentSize == 0 {
		c.Blocks[0] = b
		return
	}

	if c.Blocks[0].Value > b.Value {
		for i := 0; i < c.CurrentSize-1; i++ {
			swapBlock(i, i+1, c)
		}
		c.Blocks[0] = b
		return
	}

	if c.Blocks[0].Value < b.Value && c.CurrentSize == 1 {
		c.Blocks[1] = b
		return
	}

	for i := 0; i < c.CurrentSize-1; i++ {
		if c.Blocks[i].Value < b.Value && c.Blocks[i+1].Value > b.Value {
			for j := c.CurrentSize - 1; j <= i+1; i-- {
				swapBlock(j-1, j, c)
			}
			c.Blocks[i] = b
			b.Index = i
			return
		}
	}
}

func (c *Container) isLeaf() bool {
	if c.Blocks[0].Left == nil {
		return true
	} else {
		return false
	}
}

func (c *Container) isRootContainer() bool {
	if c.Parent == nil {
		return true
	} else {
		return false
	}
}
