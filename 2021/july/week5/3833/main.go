package main
import (
	"fmt"
)

func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }

    leftMax := make([]int, len(height))
    leftMax[0] = height[0]

    rightMax := make([]int, len(height))
    rightMax[len(height) - 1] = height[len(height) - 1]

    for i := 1; i < len(height); i++ {
        leftMax[i] = max(height[i], leftMax[i - 1])
    }

    for i := len(height) - 2; i >= 0; i-- {
        rightMax[i] = max(height[i], rightMax[i + 1])
    }

    ans := 0
    for i := 0; i < len(height); i++ {
        ans += min(leftMax[i], rightMax[i]) - height[i]
    }
    return ans
}

// Slow & Use more memory
func mySolution(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	copyHeight := make([]int, len(height))
	copy(copyHeight, height)
	currentMax := height[0]
	currentMaxIndex := int(0)
	isDown := false
	valleyStartIndex := int(0)
	for i := 1; i < len(height); i++ {
		pre := height[i - 1]
		num := height[i]
		if pre >= num {
			if !isDown {
				for j := valleyStartIndex; j < i; j++ {
					newHeight := min(height[valleyStartIndex], pre)
					if copyHeight[j] > newHeight {
						continue
					}
					copyHeight[j] = newHeight
				}
				isDown = true
				valleyStartIndex = i - 1
			}
			// fmt.Printf("valleyStartIndex = %v\n", valleyStartIndex)
		} else {
			for j := currentMaxIndex; j < i; j++ {
				newHeight := min(currentMax, num)
				if copyHeight[j] > newHeight {
					continue
				}
				copyHeight[j] = newHeight
			}

			isDown = false
			if currentMax < num {
				currentMax = num
				currentMaxIndex = i
			}
		}

		// fmt.Printf("copyHeight = %v\n", copyHeight)
	}

	if !isDown {
		for j := valleyStartIndex; j < len(height); j++ {
			newHeight := min(height[valleyStartIndex], height[len(height) - 1])
			if copyHeight[j] > newHeight {
				continue
			}
			copyHeight[j] = newHeight
		}
	}
	// fmt.Printf("copyHeight = %v\n", copyHeight)

	result := int(0)
	for i, num := range height {
		copyNum := copyHeight[i]
		result += copyNum - num
	}
	return result
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 6
	// height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	// result: 9
	// height := []int{4, 2, 0, 3, 2, 5}

	// result: 33
	// height := []int{0,1,0,2,1,0,1,3,2,1,2,14,9,0,3,8,2,1,7,4,2,1}

	// result: 0
	// height := []int{}

	// result: 83
	// height := []int{39,8,0,1,26,5,2,1,6,5,0,2,3}

	// result: 12
	height := []int{2, 8, 5, 5, 6, 1, 7, 4, 5}

	// result: 
	// height := []int{}

	result := trap(height)
	fmt.Printf("result = %v\n", result)
}

