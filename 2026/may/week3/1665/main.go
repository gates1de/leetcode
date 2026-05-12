package main

import (
	"fmt"
)

func minimumEffort(tasks [][]int) int {
	sortTasks(tasks)

	result := int(0)
	energy := int(0)
	for _, task := range tasks {
		actual, minimum := task[0], task[1]
		if energy < minimum {
			result += minimum - energy
			energy = minimum
		}
		energy -= actual
	}

	return result
}

func sortTasks(tasks [][]int) {
	n := len(tasks)
	for i := n / 2 - 1; i >= 0; i-- {
		siftDown(tasks, i, n)
	}
	for end := n - 1; end > 0; end-- {
		tasks[0], tasks[end] = tasks[end], tasks[0]
		siftDown(tasks, 0, end)
	}
}

func siftDown(tasks [][]int, root int, end int) {
	for {
		child := root * 2 + 1
		if child >= end {
			return
		}

		if child+1 < end && taskDiff(tasks[child+1]) < taskDiff(tasks[child]) {
			child++
		}

		if taskDiff(tasks[root]) <= taskDiff(tasks[child]) {
			return
		}

		tasks[root], tasks[child] = tasks[child], tasks[root]
		root = child
	}
}

func taskDiff(task []int) int {
	return task[1] - task[0]
}

func main() {
	// result: 8
	// tasks := [][]int{{1,2},{2,4},{4,8}}

	// result: 32
	// tasks := [][]int{{1,3},{2,4},{10,11},{10,12},{8,9}}

	// result: 27
	tasks := [][]int{{1, 7}, {2, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 12}}

	// result:
	// tasks := [][]int{}
	result := minimumEffort(tasks)
	fmt.Printf("result = %v\n", result)
}
