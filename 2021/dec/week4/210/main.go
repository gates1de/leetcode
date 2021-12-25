package main
import (
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
    edges := make([][]int, numCourses)
    indegree := make([]int, numCourses)

    for _, v := range prerequisites {
        edges[v[1]] = append(edges[v[1]], v[0])
        indegree[v[0]]++
    }

    queue, res := []int{}, []int{}
    for i, v := range indegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        numCourses--
        res = append(res, cur)
        for _, v := range edges[cur] {
            indegree[v]--
            if indegree[v] == 0 {
				queue = append(queue, v)
			}
        }
    }

    if numCourses != 0 {
		return []int{}
	}

    return res
}

// Time Limit Exceeded
func ngSolution(numCourses int, prerequisites [][]int) []int {
	if len(prerequisites) == 0 {
		result := make([]int, numCourses)
		for i := 0; i < numCourses; i++ {
			result[i] = i
		}
		return result
	}

	m := map[int]map[int]bool{}
	for i := 0; i < numCourses; i++ {
		m[i] = map[int]bool{}
	}
	for _, pre := range prerequisites {
		course := pre[0]
		req := pre[1]
		m[course][req] = true
	}

	result := []int{}
	helper(m, &result)
	return result
}

func helper(m map[int]map[int]bool, courses *[]int) {
	fmt.Printf("m = %v\n", m)
	if len(m) == 0 {
		return
	}

	isChanged := false

	for course, reqs := range m {
		if len(reqs) == 0 {
			delete(m, course)
			*courses = append(*courses, course)
			isChanged = true
			continue
		}

		for _, c := range *courses {
			if reqs[c] {
				isChanged = true
				delete(reqs, c)
			}
		}
	}

	if !isChanged {
		*courses = []int{}
		return
	}

	helper(m, courses)
}

func main() {
	// result: [0,1]
	// numCourses := int(2)
	// prerequisites := [][]int{{1,0}}

	// result: [0,2,1,3]
	// numCourses := int(4)
	// prerequisites := [][]int{{1,0},{2,0},{3,1},{3,2}}

	// result: [0]
	// numCourses := int(1)
	// prerequisites := [][]int{}

	// result: [0,4,2,1,3]
	// numCourses := int(5)
	// prerequisites := [][]int{{1,0},{2,0},{3,1},{3,2},{4,0}}

	// result: []
	numCourses := int(2)
	prerequisites := [][]int{{0,1},{1,0}}

	// result: 
	// numCourses := int(0)
	// prerequisites := [][]int{}

	result := findOrder(numCourses, prerequisites)
	fmt.Printf("result = %v\n", result)
}

