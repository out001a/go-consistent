package consistent

import (
	"hash/crc32"
	"strconv"
)

// virtual nodes' number per node
const VIRT_NUM = (1 << 6) - 1
// max hash value (capacity of the circle)
const CAPACITY  = 1 << 18

type Node struct {
	value  string
	bucket *bucket
}

type VirtualNode struct {
	node   *Node
	serial uint16
	hash   uint32
	hashed bool
}

type bucket struct {
	node   *Node
	vNodes [VIRT_NUM]*VirtualNode
}

func (node *Node) Bucket() *bucket {
	if node.bucket == nil {
		bucket := new(bucket)
		bucket.node = node
		for i := 0; i < len(bucket.vNodes); i++ {
			v := &VirtualNode{
				node:   node,
				serial: uint16(i),
			}
			v.Hash()
			bucket.vNodes[i] = v
		}
		node.bucket = bucket
	}
	return node.bucket
}

func (node *Node) Hash(serial uint16) uint32 {
	//md5Sum := md5.Sum([]byte(node.value + "#" + strconv.Itoa(int(serial))))
	//h, _ := strconv.ParseUint(hex.EncodeToString(md5Sum[:])[0:8], 16, 64)
	//return uint32(h % CAPACITY)
	return crc32.ChecksumIEEE([]byte(node.value + "#" + strconv.Itoa(int(serial)))) % CAPACITY
}

func (v *VirtualNode) Hash() uint32 {
	if !v.hashed {
		//v.hash = crc32.ChecksumIEEE([]byte(v.node.value + "#" + strconv.Itoa(int(v.serial))))
		v.hash = v.node.Hash(v.serial)
		v.hashed = true
	}
	return v.hash
}

func (v *VirtualNode) Value() string {
	return v.node.value
}

func (v0 *VirtualNode) Equal(v1 *VirtualNode) bool {
	return v0.hash == v1.hash
}

func (v0 *VirtualNode) Less(v1 *VirtualNode) bool {
	return v0.hash < v1.hash
}
