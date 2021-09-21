package main
import (
	"fmt"
	"sort"
)

func canReorderDoubled(arr []int) bool {
    m := make(map[int]int)
    sort.SliceStable(arr, func(i, j int) bool {
        return abs(arr[i]) < abs(arr[j])
    })
    for _, num := range arr {
        if num % 2 == 0 && m[num / 2] > 0 {
            m[num / 2]--
            if m[num / 2] == 0 {
                delete(m, num / 2)
            }
        } else {
            m[num]++
        }
    }
    return len(m) == 0
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

// Slow
func mySolution(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	sort.Ints(arr)
	negatives := []int{}
	zeros := []int{}
	positives := []int{}
	for _, num := range arr {
		if num == 0 {
			zeros = append(zeros, num)
		} else if num < 0 {
			negatives = append(negatives, num)
		} else {
			positives = append(positives, num)
		}
	}

	if len(zeros) % 2 != 0 ||
		len(negatives) % 2 != 0 ||
		len(positives) % 2 != 0 {
		return false
	}

	if len(negatives) > 0 {
		for i := 0; i < len(negatives); i++ {
			if negatives[i] == 0 {
				continue
			}
			for j := i + 1; j < len(negatives); j++ {
				if negatives[j] * 2 == negatives[i] {
					negatives[i] = 0
					negatives[j] = 0
				}
			}
		}
	}

	if !isAllZero(negatives) {
		return false
	}

	if len(positives) > 0 {
		for i := 0; i < len(positives); i++ {
			if positives[i] == 0 {
				continue
			}
			for j := i + 1; j < len(positives); j++ {
				if positives[i] * 2 == positives[j] {
					positives[i] = 0
					positives[j] = 0
				}
			}
		}
	}

	return isAllZero(positives)
}

func isAllZero(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func main() {
	// result: false
	// arr := []int{3, 1, 3, 6}

	// result: false
	// arr := []int{2, 1, 2, 6}

	// result: true
	// arr := []int{4, -2, 2, -4}

	// result: false
	// arr := []int{1, 2, 4, 16, 8, 4}

	// result: true
	// arr := []int{2, 4, 0, 0, 16, 32}

	// result: false
	// arr := []int{-5, -2}

	// result: true
	arr := []int{2, 1, 2, 1, 1, 1, 2, 2}

	// result: 
	// arr := []int{}

	result := canReorderDoubled(arr)
	fmt.Printf("result = %v\n", result)
}

