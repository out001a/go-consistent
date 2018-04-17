package consistent

type Circle []*VirtualNode

func (c *Circle) Add(v *VirtualNode) int {
	index := c.locate(v)
	if index < len(*c) && (*c)[index].Equal(v) {
		return -1
	}
	if index < len(*c) {
		rear := append([]*VirtualNode{},(*c)[index:]...)
		*c = append((*c)[:index], v)
		*c = append(*c, rear...)
	} else {
		*c = append(*c, v)
	}
	return index
}

func (c *Circle) Remove(v *VirtualNode) int {
	index := c.locate(v)
	if index < len(*c) && (*c)[index].Equal(v) && (*c)[index].Value() == v.Value() {
		*c = append((*c)[:index], (*c)[index+1:]...)
		return index
	}
	return -1
}

func (c *Circle) Lookup(v *VirtualNode) (index int) {
	if len(*c) > 0 {
		index = c.locate(v)
		if index >= len(*c) {
			index = 0
		}
	} else {
		index = -1
	}
	return
}

func (c *Circle) locate(v *VirtualNode) int {
	v.Hash()
	l, h := 0, len(*c) - 1

	for l <= h {
		//m := (l + h) / 2
		m := l + (h - l) / 2
		vm := (*c)[m]
		if v.Equal(vm) {
			return m
		}
		if v.Less(vm) {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return (l + h + 1) / 2
}