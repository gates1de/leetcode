package main
import (
	"fmt"
)

func trap(height []int) int {
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

    result := int(0)
    for i, num := range height {
        copyNum := copyHeight[i]
        result += copyNum - num
    }
    return result
}

func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}

func main() {
	// result: 6
	// height := []int{0,1,0,2,1,0,1,3,2,1,2,1}

	// result: 9
	height := []int{4,2,0,3,2,5}

	// result: 
	// height := []int{}

	result := trap(height)
	fmt.Printf("result = %v\n", result)
}

