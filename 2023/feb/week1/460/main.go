package main
import (
	"fmt"
	"container/list"
)


type LFUCache struct {
	Nodes map[int]*list.Element
	Lists map[int]*list.List
	Capacity int
	Min int
}

type Node struct {
	Key int
	Value int
	Freq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Nodes: make(map[int]*list.Element),
		Lists: make(map[int]*list.List),
		Capacity: capacity,
		Min: 0,
	}
}

func (this *LFUCache) Get(key int) int {
	v, ok := this.Nodes[key]
	if !ok {
		return -1
	}

	n := v.Value.(*Node)
	this.Lists[n.Freq].Remove(v)
	n.Freq++

	if _, ok := this.Lists[n.Freq]; !ok {
		this.Lists[n.Freq] = list.New()
	}

	newlist := this.Lists[n.Freq]
    newNode := newlist.PushBack(n)
    this.Nodes[key] = newNode
	if n.Freq - 1 == this.Min && this.Lists[n.Freq - 1].Len() == 0 {
		this.Min++
	}

	return n.Value
}

func (this *LFUCache) Put(key int, value int)  {
	if this.Capacity == 0 {
		return
	}

	if v, ok := this.Nodes[key]; ok {
		n := v.Value.(*Node)
		n.Value = value
		this.Get(key)
		return
	}

	if this.Capacity == len(this.Nodes) {
		list := this.Lists[this.Min]
		frontNode := list.Front()
		delete(this.Nodes, frontNode.Value.(*Node).Key)
		list.Remove(frontNode)
	}

	this.Min = 1
	n := &Node{Key: key, Value: value, Freq: 1}

	if _, ok := this.Lists[1]; !ok {
		this.Lists[1] = list.New()
	}

	list1 := this.Lists[1]
	newNode := list1.PushBack(n)
	this.Nodes[key] = newNode
}

func main() {
	// result: [null, null, null, 1, null, -1, 3, null, -1, 3, 4]
	capacity := int(2)
	operations := [][]int{
		{1,1,1},
		{1,2,2},
		{0,1},
		{1,3,3},
		{0,2},
		{0,3},
		{1,4,4},
		{0,1},
		{0,3},
		{0,4},
	}

	// result: 
	// capacity := int(0)
	// operations := [][]int{}

	obj := Constructor(capacity)
	for _, ope := range operations {
		if ope[0] == 0 {
			fmt.Printf("obj.Get(%v) = %v\n", ope[1], obj.Get(ope[1]))
		} else {
			obj.Put(ope[1], ope[2])
		}
	}
}

