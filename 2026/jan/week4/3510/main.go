package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	Value int64
	Left  int
	Prev  *Node
	Next  *Node
}

type Item struct {
	First  *Node
	Second *Node
	Cost   int64
	Index  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Cost == pq[j].Cost {
		return pq[i].First.Left < pq[j].First.Left
	}
	return pq[i].Cost < pq[j].Cost
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n - 1]
	item.Index = -1
	*pq = old[0 : n - 1]
	return item
}

func minimumPairRemoval(nums []int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	merged := make([]bool, len(nums))
	decreaseCount := int(0)
	count := int(0)
	head := &Node{Value: int64(nums[0]), Left: 0}
	current := head

	for i := 1; i < len(nums); i++ {
		newNode := &Node{Value: int64(nums[i]), Left: i}
		current.Next = newNode
		newNode.Prev = current

		heap.Push(pq, &Item{
			First:  current,
			Second: newNode,
			Cost:   current.Value + newNode.Value,
		})
		if nums[i - 1] > nums[i] {
			decreaseCount++
		}

		current = newNode
	}

	for decreaseCount > 0 {
		item := heap.Pop(pq).(*Item)
		first := item.First
		second := item.Second
		cost := item.Cost

		if merged[first.Left] || merged[second.Left] || first.Value + second.Value != cost {
			continue
		}

		count++
		if first.Value > second.Value {
			decreaseCount--
		}

		prevNode := first.Prev
		nextNode := second.Next
		first.Next = nextNode
		if nextNode != nil {
			nextNode.Prev = first
		}

		if prevNode != nil {
			if prevNode.Value > first.Value && prevNode.Value <= cost {
				decreaseCount--
			} else if prevNode.Value <= first.Value && prevNode.Value > cost {
				decreaseCount++
			}

			heap.Push(pq, &Item{
				First:  prevNode,
				Second: first,
				Cost:   prevNode.Value + cost,
			})
		}

		if nextNode != nil {
			if second.Value > nextNode.Value && cost <= nextNode.Value {
				decreaseCount--
			} else if second.Value <= nextNode.Value && cost > nextNode.Value {
				decreaseCount++
			}

			heap.Push(pq, &Item{
				First:  first,
				Second: nextNode,
				Cost:   cost + nextNode.Value,
			})
		}

		first.Value = cost
		merged[second.Left] = true
	}

	return count
}

func main() {
	// result: 2
	// nums := []int{5,2,3,1}

	// result: 0
	nums := []int{1,2,2}

	// result: 
	// nums := []int{}

	result := minimumPairRemoval(nums)
	fmt.Printf("result = %v\n", result)
}

