package consistent

import (
	"sync"
)

type Consistent struct {
	sync.RWMutex
	Circle *Circle
}

func NewConsistent() *Consistent {
	return &Consistent{Circle:&Circle{}}
}

func (c *Consistent) Add(value string) {
	c.AddNode(&Node{Value:value})
}

func (c *Consistent) Remove(value string) {
	c.RemoveNode(&Node{Value:value})
}

func (c *Consistent) Lookup(value string) (string, bool) {
	if node := c.GetNode(c.LookupByValue(value)); node != nil {
		return node.Value, true
	}
	return "", false
}

func (c *Consistent) AddNode(node *Node) {
	c.Lock()
	defer c.Unlock()
	bucket := node.Bucket()
	for _, v := range bucket.vNodes {
		c.Circle.Add(v)
	}
}

func (c *Consistent) RemoveNode(node *Node) {
	c.Lock()
	defer c.Unlock()
	bucket := node.Bucket()
	for _, v := range bucket.vNodes {
		c.Circle.Remove(v)
	}
}

func (c *Consistent) GetNode(index int) *Node {
	if index < 0 {
		return nil
	}
	return (*c.Circle)[index].node
}

func (c *Consistent) LookupByValue(value string) int {
	return c.LookupByNode(&Node{Value:value})
}

func (c *Consistent) LookupByNode(node *Node) int {
	return c.Circle.Lookup(&VirtualNode{node:node})
}
