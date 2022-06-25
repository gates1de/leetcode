package main
import (
	"fmt"
	"sort"
	"container/heap"
)

type CourseHeap [][]int

func (h CourseHeap) Len() int           { return len(h) }
func (h CourseHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h CourseHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h CourseHeap) Peek() []int        { return h[0] }

func (h *CourseHeap) Push(x interface{}) {
    *h = append(*h, x.([]int))
}

func (h *CourseHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0 : n - 1]
    return x
}

func scheduleCourse(courses [][]int) int {
    if len(courses) == 0 {
        return 0
    }

    sort.Slice(courses, func (a, b int) bool {
        return courses[a][1] < courses[b][1]
    })

    h := &CourseHeap{}
    heap.Init(h)
    currentTotalTime := int(0)
    if courses[0][0] <= courses[0][1] {
        heap.Push(h, courses[0])
        currentTotalTime += courses[0][0]
    }
    for _, course := range courses[1:] {
        if course[0] > course[1] {
            continue
        }
        currentTotalTime += course[0]
        heap.Push(h, course)

        if currentTotalTime > course[1] {
            longestDurationCourse := heap.Pop(h).([]int)
            currentTotalTime -= longestDurationCourse[0]
        }
    }
    return h.Len()
}

// Time Limit Exceeded
func ngSolution(courses [][]int) int {
	filteredCourses := [][]int{}
	for _, course := range courses {
		if course[0] > course[1] {
			continue
		}
		filteredCourses = append(filteredCourses, course)
	}

	sort.Slice(filteredCourses, func (a, b int) bool {
		return filteredCourses[a][1] < filteredCourses[b][1]
	})

	result := int(0)
	for i, course := range filteredCourses {
		helper(course[0], 1, filteredCourses[i + 1:], &result)
	}
	return result
}

func helper(currentDay int, count int, courses [][]int, result *int) {
	if *result < count {
		*result = count
	}

	if len(courses) == 0 {
		return
	}

	for i, course := range courses {
		if currentDay + course[0] > course[1] {
			continue
		}
		helper(currentDay + course[0], count + 1, courses[i + 1:], result)
	}
}

func main() {
	// result: 3
	// courses := [][]int{{100,200},{200,1300},{1000,1250},{2000,3200}}

	// result: 1
	// courses := [][]int{{1, 2}}

	// result: 0
	// courses := [][]int{{3, 2},{4, 3}}

	// result: 4
	// courses := [][]int{{2,5},{2,19},{1,8},{1,3}}

	// result: 4
	// courses := [][]int{{7,17},{3,12},{10,20},{9,10},{5,20},{10,19},{4,18}}

	// result: 3
	// courses := [][]int{{9,14},{7,12},{1,11},{4,7}}

	// result: 3
	courses := [][]int{{1,19},{2,2},{1,17}}

	// result: 
	// courses := [][]int{}

	result := scheduleCourse(courses)
	fmt.Printf("result = %v\n", result)
}

