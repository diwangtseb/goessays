package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"testing"
)

func TestConsistentHash(t *testing.T) {
	ch := NewConsistentHash(100)

	ch.AddNode(1, "192.168.0.1")
	ch.AddNode(2, "192.168.0.2")
	ch.AddNode(3, "192.168.0.3")

	node1 := ch.GetNodeByKey("key1")
	fmt.Println(node1.Id, node1.Ip)

	node2 := ch.GetNodeByKey("key2")
	fmt.Println(node2.Id, node2.Ip)

	node3 := ch.GetNodeByKey("key3")
	fmt.Println(node3.Id, node3.Ip)

	ch.RemoveNode("192.168.0.1")
	ch.RemoveNode("192.168.0.2")

	node3 = ch.GetNodeByKey("key3")
	fmt.Println(node3.Id, node3.Ip)
	node3 = ch.GetNodeByKey("key2")
	fmt.Println(node3.Id, node3.Ip)
	node3 = ch.GetNodeByKey("key1")
	fmt.Println(node3.Id, node3.Ip)
}

type HashCircle []uint32

func (hc HashCircle) Len() int {
	return len(hc)
}

func (hc HashCircle) Less(i, j int) bool {
	return hc[i] < hc[j]
}

func (hc HashCircle) Swap(i, j int) {
	hc[i], hc[j] = hc[j], hc[i]
}

type Node struct {
	Id       int
	Ip       string
	Hashcode uint32
}

type Nodes []*Node

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Less(i, j int) bool {
	return n[i].Hashcode < n[j].Hashcode
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

type ConsistentHash struct {
	Circle HashCircle
	Nodes  Nodes
	N      int
}

func NewConsistentHash(n int) *ConsistentHash {
	return &ConsistentHash{
		Circle: make([]uint32, 0),
		Nodes:  make([]*Node, 0),
		N:      n,
	}
}

func (ch *ConsistentHash) AddNode(id int, ip string) bool {
	for _, node := range ch.Nodes {
		if node.Id == id || node.Ip == ip {
			return false
		}
	}

	node := &Node{Id: id, Ip: ip, Hashcode: crc32.ChecksumIEEE([]byte(ip))}
	ch.Nodes = append(ch.Nodes, node)
	ch.Circle = append(ch.Circle, node.Hashcode)
	sort.Sort(ch.Circle)
	return true
}

func (ch *ConsistentHash) RemoveNode(ip string) bool {
	for i, node := range ch.Nodes {
		if node.Ip == ip {
			n := len(ch.Nodes)
			ch.Nodes[i] = ch.Nodes[n-1]
			ch.Nodes = ch.Nodes[:n-1]

			m := len(ch.Circle)
			var h uint32
			for index, value := range ch.Circle {
				if value == node.Hashcode {
					h = value
					ch.Circle[index] = ch.Circle[m-1]
					ch.Circle = ch.Circle[:m-1]
					break
				}
			}

			if m != 1 {
				j := sort.Search(len(ch.Circle), func(i int) bool { return ch.Circle[i] >= h })
				if j == len(ch.Circle) {
					j = 0
				}
				copy(ch.Circle[j+1:], ch.Circle[j:])
				ch.Circle[j] = h
			}

			return true
		}
	}

	return false
}

func (ch *ConsistentHash) GetNodeByKey(key string) *Node {
	hashcode := crc32.ChecksumIEEE([]byte(key))

	i := sort.Search(len(ch.Circle), func(i int) bool { return ch.Circle[i] >= hashcode })

	if i == len(ch.Circle) {
		i = 0
	}

	return ch.Nodes[i]
}
