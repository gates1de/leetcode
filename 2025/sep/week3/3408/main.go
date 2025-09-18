package main

import (
	heap "container/heap"
	"fmt"
)

type TaskHeap []Task

func (h TaskHeap) Len() int { return len(h) }

func (h TaskHeap) Less(i, j int) bool {
    if h[i].Priority == h[j].Priority {
        return h[i].TaskId > h[j].TaskId
    }
    return h[i].Priority > h[j].Priority
}

func (h TaskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *TaskHeap) Push(x any) {
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n-1]
	return x
}

type Task struct {
	UserId   int
	TaskId   int
	Priority int
}

type TaskManager struct {
	Heap TaskHeap
	Index map[int]Task
}

func Constructor(tasks [][]int) TaskManager {
	result := TaskManager{
		Heap: make(TaskHeap, 0, len(tasks)),
		Index: make(map[int]Task, len(tasks)),
	}
	for _,task := range tasks {
		t := Task{task[0], task[1], task[2]}
		result.Heap.Push(t)
		result.Index[t.TaskId] = t
	}

	heap.Init(&result.Heap)
	return result
}

func (this *TaskManager) Add(userId int, taskId int, priority int)  {
	t := Task{userId, taskId, priority}
	heap.Push(&this.Heap,t)
	this.Index[t.TaskId] = t
}

func (this *TaskManager) Edit(taskId int, newPriority int)  {
	t, _ := this.Index[taskId]
	t.Priority = newPriority
	this.Index[taskId] = t
	heap.Push(&this.Heap, t)
}

func (this *TaskManager) Rmv(taskId int)  {
	delete(this.Index, taskId)
}

func (this *TaskManager) ExecTop() int {
	if this.Heap.Len() <= 0 {
		return -1
	}

	pop := heap.Pop(&this.Heap).(Task)
	t, exists := this.Index[pop.TaskId]

	for (!exists || t.Priority != pop.Priority) && this.Heap.Len() > 0 {
		pop = heap.Pop(&this.Heap).(Task)
		t, exists = this.Index[pop.TaskId]
	}

	if exists && t.Priority == pop.Priority {
		delete(this.Index,pop.TaskId)
		return t.UserId
	}

	return -1
}

func main() {
	//result: [null, null, null, 3, null, null, 5]
	tasks := [][]int{{1,101,10},{2,102,20},{3,103,15}}
	operations := [][][]int{
		{{1}, {4,104,5}},
		{{2}, {102, 8}},
		{{4}, {}},
		{{3}, {101}},
		{{1}, {5,105,15}},
		{{4}, {}},
	}

	//result: [null, null, null, 3, null, null, 5]
	// tasks := [][]int{}
	// operations := [][][][]int{
	// 	{1, {{}}},
	// 	{2, {{}}},
	// 	{3, {{}}},
	// 	{4, {{}}},
	// }

	obj := Constructor(tasks)

	for _, ope := range operations {
		if ope[0][0] == 1 {
			obj.Add(ope[1][0], ope[1][1], ope[1][2])
		} else if ope[0][0] == 2 {
			obj.Edit(ope[1][0], ope[1][1])
		} else if ope[0][0] == 3 {
			obj.Rmv(ope[1][0]);
		} else if ope[0][0] == 4 {
			fmt.Printf("obj.ExecTop() = %v\n", obj.ExecTop())
		}
	}
}

