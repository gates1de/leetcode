package main
import (
	"fmt"
	list "container/list"
	"sort"
)

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	n := len(tasks)
	m := len(workers)
	sort.Ints(tasks)
	sort.Ints(workers)

	left := int(1)
	right := min(m, n)
	result := int(0)

	for left <= right {
		mid := (left + right) / 2
		if check(mid, tasks, workers, pills, strength) {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func check(mid int, tasks []int, workers []int, pills int, strength int) bool {
	m := len(workers)
	queue := list.New()
	pointer := m - 1

	for i := mid - 1; i >= 0; i-- {
		for pointer >= m - mid && workers[pointer] + strength >= tasks[i] {
			queue.PushFront(workers[pointer])
			pointer--
		}

		if queue.Len() == 0 {
			return false
		}

		if queue.Back().Value.(int) >= tasks[i] {
			queue.Remove(queue.Back())
		} else {
			if pills == 0 {
				return false
			}
			pills--
			queue.Remove(queue.Front())
		}
	}

	return true
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// tasks := []int{3,2,1}
	// workers := []int{0,3,3}
	// pills := int(1)
	// strength := int(1)

	// result: 1
	// tasks := []int{5,4}
	// workers := []int{0,0,0}
	// pills := int(1)
	// strength := int(5)

	// result: 2
	tasks := []int{10,15,30}
	workers := []int{0,10,10,10,10}
	pills := int(3)
	strength := int(10)

	// result: 
	// tasks := []int{}
	// workers := []int{}
	// pills := int(0)
	// strength := int(0)

	result := maxTaskAssign(tasks, workers, pills, strength)
	fmt.Printf("result = %v\n", result)
}

