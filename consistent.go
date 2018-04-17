package consistent

import (
	"sync"
)

type Consistent struct {
	sync.RWMutex
	circle *Circle
}

func (c *Consistent) GetNode(index int) *Node {
	return (*c.circle)[index].node;
}

func (c *Consistent) AddNode(node *Node) {
	c.Lock()
	defer c.Unlock()
	bucket := node.Bucket()
	for _, v := range bucket.vNodes {
		c.circle.Add(v)
	}
}

func (c *Consistent) RemoveNode(node *Node) {
	c.Lock()
	defer c.Unlock()
	bucket := node.Bucket()
	for _, v := range bucket.vNodes {
		c.circle.Remove(v)
	}
}

func (c *Consistent) Lookup(node *Node) int {
	return c.circle.Lookup(&VirtualNode{node:node})
}
