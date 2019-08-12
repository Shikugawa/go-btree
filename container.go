package container

type Block struct {
	value int
	left  *Container
	right *Container
}

type Container struct {
	size   int
	blocks []Block
	parent *Container
}

func (c *Container) CreateRightContainer() *Container {
	container := &Container{
		size: c.size,
	}

	for i := c.size/2 + 1; i < c.size; i += 1 {
		blc := Block{
			value: c.blocks[i].value,
			left:  nil,
			right: nil,
		}

		container.blocks = append(container.blocks, blc)
	}

	return container
}

func (c *Container) CreateLeftContainer() *Container {
	container := &Container{
		size: c.size,
	}

	for i := 0; i < c.size/2; i += 1 {
		blc := Block{
			value: c.blocks[i].value,
			left:  nil,
			right: nil,
		}

		container.blocks = append(container.blocks, blc)
	}

	return container
}

func (c *Container) ReConstructContainer(mid *Block) {
	var leftContainer, rightContainer *Container

	ch := make([]chan bool, 2)

	go func() {
		leftContainer = c.CreateRightContainer()
		ch[0] <- true
	}()

	go func() {
		rightContainer = c.CreateLeftContainer()
		ch[1] <- true
	}()

	for i := 0; i < 2; i += 1 {
		<-ch[i]
	}

	blk := &Block{
		value: mid.value,
		left:  leftContainer,
		right: rightContainer,
	}
	c.parent.AddNode(*blk)
}

func (c *Container) AddNode(b Block) {
	if len(c.blocks) == c.size {
		mid := c.blocks[c.size/2]
		c.ReConstructContainer(&mid)
	}

	for i := 0; i < len(c.blocks); i += 1 {
		if c.blocks[i].value < b.value && c.blocks[i+1].value > b.value {
			if c.blocks[i].right.IsLeaf() {
				c.blocks[i].right.AddLeaf(b)
			} else {
				c.blocks[i].right.AddNode(b)
			}
		}
	}
}

func (c *Container) AddLeaf(b Block) {
	if len(c.blocks) == c.size {
		mid := c.blocks[c.size/2]
		c.ReConstructContainer(&mid)
	}

	for i := 0; i < len(c.blocks); i += 1 {
		if c.blocks[i].value > b.value && c.blocks[i+1].value < b.value {
			for j := len(c.blocks) + 1; j < i+1; i -= 1 {
				c.blocks[j] = c.blocks[j-1]
			}
			c.blocks[i+1] = b
		}
	}
}

func (c *Container) IsLeaf() bool {
	if c.blocks[0].left == nil {
		return true
	} else {
		return false
	}
}
