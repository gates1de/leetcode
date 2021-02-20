package main
import (
    localHeap "./heap"
    "container/heap"
    "fmt"
	"math"
    "sort"
)

func minimumDeviation(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	minNum := math.MaxInt64
	numsHeap := &localHeap.MaxHeap{}
	heap.Init(numsHeap)
	for _, num := range nums {
		if num % 2 == 1 {
			heap.Push(numsHeap, num * 2)
			minNum = min(minNum, num * 2)
		} else {
			heap.Push(numsHeap, num)
			minNum = min(minNum, num)
		}
	}

	result := math.MaxInt64
	for numsHeap.Len() > 0 {
		num := heap.Pop(numsHeap).(int)
		result = min(result, num - minNum)
		// fmt.Printf("result = %v, num = %v, minNum = %v, abs = %v\n", result, num, minNum, abs(num, minNum))

		if num % 2 == 0 {
			heap.Push(numsHeap, num / 2)
			minNum = min(minNum, num / 2)
			continue
		}
		break
	}
	return result
}

// Wrong Answer
func ngSolution(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Slice(nums, func (i, j int) bool { return nums[i] < nums[j] })
	minNum := nums[0]
	maxNum := nums[len(nums) - 1]
	fmt.Printf("1: nums = %v, max = %v\n", nums, maxNum)
	for i, num := range nums {
		if num % 2 == 1 && num * 2 <= maxNum {
			nums[i] = num * 2
		}
	}
	sort.Slice(nums, func (i, j int) bool { return nums[i] < nums[j] })
	minNum = nums[0]
	maxNum = nums[len(nums) - 1]
	fmt.Printf("2: nums = %v, max = %v, min = %v\n", nums, maxNum, minNum)

	minEven := int(-1)
	minOdd := int(-1)
	for _, num := range nums {
		if num == maxNum {
			break
		}
		if num % 2 == 0 && minEven == -1 {
			minEven = num
		} else if num % 2 == 1 && minOdd == -1 {
			minOdd = num
		}
		if minEven > 0 && minOdd > 0 {
			break
		}
	}

	fmt.Printf("minEven = %v, minOdd = %v\n", minEven, minOdd)

	if minEven < 0 && minOdd >= 0 {
		return min(abs(maxNum, minOdd), abs(maxNum, minOdd * 2))
	} else if minEven >= 0 && minOdd < 0 {
		return abs(maxNum, minEven)
	} else if minEven < 0 && minOdd < 0 {
		return 0
	}

	minAbs := findMinAbs([]int{minOdd, minEven, maxNum})
	doubleOddMinAbs := findMinAbs([]int{minOdd * 2, minEven, maxNum})
	return min(minAbs, doubleOddMinAbs)
	// return nums[len(nums) - 1] - nums[0]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func abs(a int, b int) int {
    if a - b < 0 {
        return b - a
    }
    return a - b
}

func findMinAbs(nums []int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		return nums[0]
	}
	sort.Slice(nums, func (i, j int) bool { return nums[i] < nums[j] })
	return nums[len(nums) - 1] - nums[0]
}

func remove(i int, nums []int) []int {
    nums = append(nums[:i], nums[i + 1:]...)
    result := copySlice(nums)
    return result
}

func copySlice(target []int) []int {
    result := make([]int, len(target))
    copy(result, target)
    return result
}

func main() {
	// nums := []int{1, 2, 3, 4}
	// nums := []int{4, 1, 5, 20, 3}
	// nums := []int{2, 10, 8}
    // nums := []int{3, 5}
	// nums := []int{4, 9, 4, 5}
	// nums := []int{10, 4, 3}
	// nums := []int{8, 1, 2, 1}
	nums := []int{399, 908, 648, 357, 693, 502, 331, 649, 596, 698}

	result := minimumDeviation(nums)
	fmt.Printf("result = %v\n", result)
}

