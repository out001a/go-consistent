package consistent

import (
	"testing"
	"strconv"
	"log"
)

var consistent = &Consistent{circle:&Circle{}}

var hosts = []string{"127.0.0.1", "192.168.1.100", "192.168.1.101"}

func TestConsistent_AddNode(t *testing.T) {
	logCircle()
	nodes := [3]*Node{}
	for i, node := range nodes {
		// rand.Seed(time.Now().UnixNano() + int64(1000 * i))
		// host := hosts[rand.Intn(len(hosts))]
		host := hosts[0]
		port := 80 + i
		node = &Node{value:host + ":" + strconv.Itoa(port)}
		//log.Println(node.value)
		consistent.AddNode(node)
		//log.Println(len(*consistent.circle))
	}
	log.Printf("count: %d\n", len(*consistent.circle))
	logCircle()
}

func TestConsistent_RemoveNode(t *testing.T) {
	return
	host := hosts[0]
	node := &Node{value:host + ":" + strconv.Itoa(81)}
	consistent.RemoveNode(node)
	//log.Println(len(*consistent.circle))
	log.Printf("count: %d\n", len(*consistent.circle))
}

func TestConsistent_Lookup(t *testing.T) {
	n1, n2, n3, n4, n5, n6, n7 :=  &Node{value:"abc"}, &Node{value:"def"}, &Node{value:"ghi"}, &Node{value:"jkl"}, &Node{value:"mno"}, &Node{value:"pqr"}, &Node{value:"stu"}
	i1, i2, i3, i4, i5, i6, i7 := consistent.Lookup(n1), consistent.Lookup(n2),consistent.Lookup(n3), consistent.Lookup(n4), consistent.Lookup(n5), consistent.Lookup(n6), consistent.Lookup(n7)
	log.Printf("locate: %d\t%s\t%d\t%d\n", i1, consistent.GetNode(i1).value, (*consistent.circle)[i1].hash, n1.Hash(0))
	log.Printf("locate: %d\t%s\t%d\t%d\n", i2, consistent.GetNode(i2).value, (*consistent.circle)[i2].hash, n2.Hash(0))
	log.Printf("locate: %d\t%s\t%d\t%d\n", i3, consistent.GetNode(i3).value, (*consistent.circle)[i3].hash, n3.Hash(0))
	log.Printf("locate: %d\t%s\t%d\t%d\n", i4, consistent.GetNode(i4).value, (*consistent.circle)[i4].hash, n4.Hash(0))
	log.Printf("locate: %d\t%s\t%d\t%d\n", i5, consistent.GetNode(i5).value, (*consistent.circle)[i5].hash, n5.Hash(0))
	log.Printf("locate: %d\t%s\t%d\t%d\n", i6, consistent.GetNode(i6).value, (*consistent.circle)[i6].hash, n6.Hash(0))
	log.Printf("locate: %d\t%s\t%d\t%d\n", i7, consistent.GetNode(i7).value, (*consistent.circle)[i7].hash, n7.Hash(0))
}

func logCircle() {
	for k, v := range *consistent.circle {
		log.Printf("[%d] %s\t%d\t%d\t%d\n", k, v.node.value, v.serial, v.hash, len(strconv.Itoa(int(v.hash))))
	}
	log.Println("========")
}