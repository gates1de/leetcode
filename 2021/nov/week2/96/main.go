package main
import (
	"fmt"
)

func numTrees(n int) int {
	result := int(0)
	for i := 1; i <= n; i++ {
		m := map[int]bool{}
		m[i] = true
		// fmt.Printf("---------- root: %v ----------\n", i)
		helper(i, i, n - 1, n, "", m, &result)
	}
	return result
}

func helper(root int, parent int, remain int, n int, dir string, m map[int]bool, result *int) {
	// fmt.Printf("parent = %v, remain = %v, dir = %v, m = %v\n", parent, remain, dir, m)
	if remain == 0 {
		*result++
		return
	}

	for i := 1; i <= n; i++ {
		if m[i] || ((dir == "left" && root < i) || (dir == "right" && i < root)) {
			continue
		}
		if i < parent {
			newM := copyMap(m)
			newM[i] = true
			if dir == "" {
				helper(root, i, remain - 1, n, "left", newM, result)
				helper(root, parent, remain - 1, n, "right", newM, result)
			} else {
				helper(root, i, remain - 1, n, dir, newM, result)
			}
		}
		if parent < i {
			newM := copyMap(m)
			newM[i] = true
			if dir == "" {
				helper(root, i, remain - 1, n, "right", newM, result)
			} else {
				helper(root, i, remain - 1, n, dir, newM, result)
			}
		}
	}
}

func copyMap(m map[int]bool) map[int]bool {
    result := map[int]bool{}
    for k, v := range m {
        result[k] = v
    }
    return result
}

func main() {
	// result: 5
	// n := int(3)

	// result: 1
	// n := int(1)

	// result: 42
	n := int(5)

	// result: 
	// n := int(0)

	result := numTrees(n)
	fmt.Printf("result = %v\n", result)
}

