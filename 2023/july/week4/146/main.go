package main
import (
	"fmt"
)


type LRUCache struct {
    Capacity int
    Head *Node
    Tail *Node
    Values map[int]*Node
}

type Node struct {
    Key int
    Value int
    Prev *Node
    Next *Node
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        Capacity: capacity,
		Head: nil,
		Tail: nil,
        Values: make(map[int]*Node),
    }
}

func (this *LRUCache) Get(key int) int { 
    if _, ok := this.Values[key]; !ok {
		return -1
	}

	node := this.Values[key]
    this.MoveToLast(node)
    return node.Value
}

func (this *LRUCache) MoveToLast(node *Node) {
    if node == this.Tail {
		return
	}

    if node == this.Head {
        this.Head = this.Head.Next
        this.Head.Prev = nil
    } else {
        node.Prev.Next = node.Next
        node.Next.Prev = node.Prev
    }

    this.Tail.Next = node
    node.Prev = this.Tail
    this.Tail = this.Tail.Next
    this.Tail.Next = nil    
}

func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.Values[key]; ok {
        this.Values[key].Value = value
        this.MoveToLast(this.Values[key])
        return
    }

    if len(this.Values) < this.Capacity {
        this.Append(key, value)
        return
    }

    node := this.Head
    this.MoveToLast(node)
    delete(this.Values, node.Key)

    this.Values[key] = node
    node.Key = key
    node.Value = value
}

func (this *LRUCache) Append(key, value int) {
    node := &Node{
        Key: key,
        Value: value,
		Prev: nil,
		Next: nil,
    }

    if this.Tail == nil {
        this.Tail = node
        this.Head = node
    } else {
        this.Tail.Next = node
        node.Prev = this.Tail
        this.Tail = node
    }
    this.Values[key] = node 
}

func main() {
	// result: [null, null, null, 1, null, -1, null, -1, 3, 4]
	capacity := int(2)
	operations := [][]int{{0,1,1},{0,2,2},{1,1},{0,3,3},{1,2},{0,4,4},{1,1},{1,3},{1,4}}

	// result: 
	// capacity := int(0)
	// operations := [][]int{{0,0,0},{1,0}}

	obj := Constructor(capacity)
	for _, ope := range operations {
		if ope[0] == 0 {
			obj.Put(ope[1],ope[2])
		} else if ope[0] == 1 {
			fmt.Printf("obj.Get(%v) = %v\n", ope[1], obj.Get(ope[1]))
		}
	}
}

