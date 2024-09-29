package main
import (
	"fmt"
)

type AllOne struct {
	Head *Node
	Tail *Node
	Map map[string]*Node
}

type Node struct {
	Freq int
	Prev *Node
	Next *Node
	Keys map[string]bool
}

func Constructor() AllOne {
	allOne := AllOne{Head: &Node{}, Tail: &Node{}, Map: make(map[string]*Node)}
	allOne.Head.Next = allOne.Tail
	allOne.Tail.Prev = allOne.Head
	return allOne
}

func (this *AllOne) Inc(key string)  {
	if this.Map[key] != nil {
		node := this.Map[key]
		freq := node.Freq
		delete(node.Keys, key)

		nextNode := node.Next
		if nextNode == this.Tail || nextNode.Freq != freq + 1 {
			newNode := &Node{Freq: freq + 1, Keys: make(map[string]bool)}
			newNode.Keys[key] = true
			newNode.Prev = node
			newNode.Next = nextNode
			node.Next = newNode
			nextNode.Prev = newNode
			this.Map[key] = newNode
		} else {
			nextNode.Keys[key] = true
			this.Map[key] = nextNode
		}

		if len(node.Keys) == 0 {
			removeNode(node)
		}
	} else {
		firstNode := this.Head.Next
		if firstNode == this.Tail || firstNode.Freq > 1 {
			newNode := &Node{Freq: 1, Keys: make(map[string]bool)}
			newNode.Keys[key] = true
			newNode.Prev = this.Head
			newNode.Next = firstNode
			this.Head.Next = newNode
			firstNode.Prev = newNode
			this.Map[key] = newNode
		} else {
			firstNode.Keys[key] = true
			this.Map[key] = firstNode
		}
	}
}

func (this *AllOne) Dec(key string)  {
	if this.Map[key] == nil {
		return
	}

	node := this.Map[key]
	delete(node.Keys, key)
	freq := node.Freq

	if freq == 1 {
		delete(this.Map, key)
	} else {
		prevNode := node.Prev

		if prevNode == this.Head || prevNode.Freq != freq - 1 {
			newNode := &Node{Freq: freq - 1, Keys: make(map[string]bool)}
			newNode.Keys[key] = true
			newNode.Prev = prevNode
			newNode.Next = node
			prevNode.Next = newNode
			node.Prev = newNode
			this.Map[key] = newNode
		} else {
			prevNode.Keys[key] = true
			this.Map[key] = prevNode
		}
	}


	if len(node.Keys) == 0 {
		removeNode(node)
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.Tail.Prev == this.Head {
		return ""
	}

	for key, _ := range this.Tail.Prev.Keys {
		return key
	}

	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.Head.Next == this.Tail {
		return ""
	}

	for key, _ := range this.Head.Next.Keys {
		return key
	}

	return ""
}

func removeNode(node *Node) {
	prevNode := node.Prev
	nextNode := node.Next
	prevNode.Next = nextNode
	nextNode.Prev = prevNode
}

func main() {
	// result: [null, null, null, "hello", "hello", null, "hello", "leet"]
	operations := [][]string{{"0", "hello"}, {"0", "hello"}, {"2", ""}, {"3", ""}, {"0", "leet"}, {"2", ""}, {"3", ""}}

	obj := Constructor()
	for _, ope := range operations {
		function := ope[0]
		key := ope[1]

		if function == "0" {
			obj.Inc(key)
		} else if function == "1" {
			obj.Dec(key)
		} else if function == "2" {
			fmt.Printf("obj.GetMaxKey() = %s\n", obj.GetMaxKey())
		} else if function == "3" {
			fmt.Printf("obj.GetMinKey() = %s\n", obj.GetMinKey())
		}
	}
}

