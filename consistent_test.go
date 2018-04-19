package consistent

import (
	"testing"
	"strconv"
	"log"
)

var consistent = NewConsistent()

var hosts = []string{"127.0.0.1", "192.168.1.100", "192.168.1.101"}

func TestConsistent_LookupEmpty(t *testing.T) {
	s, ok := consistent.Lookup("123")
	log.Println(s, ok)
}

func TestConsistent_AddNode(t *testing.T) {
	logCircle()
	nodes := [3]*Node{}
	for i, node := range nodes {
		// rand.Seed(time.Now().UnixNano() + int64(1000 * i))
		// host := hosts[rand.Intn(len(hosts))]
		host := hosts[0]
		port := 80 + i
		node = &Node{Value:host + ":" + strconv.Itoa(port)}
		//log.Println(node.value)
		consistent.AddNode(node)
		//log.Println(len(*consistent.circle))
	}
	log.Printf("count: %d\n", len(*consistent.Circle))
	logCircle()
}

func TestConsistent_RemoveNode(t *testing.T) {
	return
	host := hosts[0]
	node := &Node{Value:host + ":" + strconv.Itoa(81)}
	consistent.RemoveNode(node)
	//log.Println(len(*consistent.circle))
	log.Printf("count: %d\n", len(*consistent.Circle))
}

func TestConsistent_LookupByNode(t *testing.T) {
	n1, n2, n3, n4, n5, n6, n7 :=  &Node{Value:"abc"}, &Node{Value:"def"}, &Node{Value:"ghi"}, &Node{Value:"jkl"}, &Node{Value:"mno"}, &Node{Value:"pqr"}, &Node{Value:"stu"}
	i1, i2, i3, i4, i5, i6, i7 := consistent.LookupByNode(n1), consistent.LookupByNode(n2),consistent.LookupByNode(n3), consistent.LookupByNode(n4), consistent.LookupByNode(n5), consistent.LookupByNode(n6), consistent.LookupByNode(n7)
	log.Printf("locate: %d\t%s\t%d\t%d\n", i1, consistent.GetNode(i1).Value, (*consistent.Circle)[i1].hash, n1.Hash(0))
	log.Printf("locate: %d\t%s\t%d\t%d\n", i2, consistent.GetNode(i2).Value, (*consistent.Circle)[i2].hash, n2.Hash(0))
	log.Printf("locate: %d\t%s\t%d\t%d\n", i3, consistent.GetNode(i3).Value, (*consistent.Circle)[i3].hash, n3.Hash(0))
	log.Printf("locate: %d\t%s\t%d\t%d\n", i4, consistent.GetNode(i4).Value, (*consistent.Circle)[i4].hash, n4.Hash(0))
	log.Printf("locate: %d\t%s\t%d\t%d\n", i5, consistent.GetNode(i5).Value, (*consistent.Circle)[i5].hash, n5.Hash(0))
	log.Printf("locate: %d\t%s\t%d\t%d\n", i6, consistent.GetNode(i6).Value, (*consistent.Circle)[i6].hash, n6.Hash(0))
	log.Printf("locate: %d\t%s\t%d\t%d\n", i7, consistent.GetNode(i7).Value, (*consistent.Circle)[i7].hash, n7.Hash(0))

	log.Println("--------")
	counts := make(map[string]int, 26)
	for i := 0; i < 26; i++ {
		s := string(65 + i) + strconv.Itoa(1000 + i)
		t, _ := consistent.Lookup(s)
		counts[t]++
		log.Printf("%s\t--> %s\n", s, t)
	}
	log.Println(counts)
}

func logCircle() {
	for k, v := range *consistent.Circle {
		log.Printf("[%d] %s\t%d\t%d\t%d\n", k, v.node.Value, v.serial, v.hash, len(strconv.Itoa(int(v.hash))))
	}
	log.Println("========")
}