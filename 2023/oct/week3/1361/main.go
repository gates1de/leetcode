package main
import (
	"fmt"
)

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	root := findRoot(n, leftChild, rightChild)
	if root == -1 {
		return false
	}

	seen := make(map[int]bool)
	stack := make([]int, 0, n)
	seen[root] = true
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		children := []int{leftChild[node], rightChild[node]}
		for _, child := range children {
			if child != -1 {
				if seen[child] {
					return false
				}

				stack = append(stack, child)
				seen[child] = true
			}
		}
	}

	return len(seen) == n
}

func findRoot(n int, leftChild []int, rightChild []int) int {
	children := make(map[int]bool)
	for _, node := range leftChild {
		children[node] = true
	}
	for _, node := range rightChild {
		children[node] = true
	}

	for i := 0; i < n; i++ {
		if !children[i] {
			return i
		}
    }
    
    return -1
}

func main() {
	// result: true
	// n := int(4)
	// leftChild := []int{1,-1,3,-1}
	// rightChild := []int{2,-1,-1,-1}

	// result: false
	// n := int(4)
	// leftChild := []int{1,-1,3,-1}
	// rightChild := []int{2,3,-1,-1}

	// result: false
	n := int(2)
	leftChild := []int{1,0}
	rightChild := []int{-1,-1}

	// result: 
	// n := int(0)
	// leftChild := []int{}
	// rightChild := []int{}

	result := validateBinaryTreeNodes(n, leftChild, rightChild)
	fmt.Printf("result = %v\n", result)
}

