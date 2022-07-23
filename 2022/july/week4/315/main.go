package main
import (
	"fmt"
)

func countSmaller(nums []int) []int {
    indexes := make([]int, 0, len(nums))
    for i := range nums {
        indexes = append(indexes, i)
    }
    inversions := make([]int, len(nums), len(nums))
    merge(0, len(nums)-1, nums, indexes, inversions)
    return inversions
}

func merge(left int, right int, nums []int, indexes []int, inversions []int) {
    if left >= right {
        return
    }

    mid := (left + right) / 2
    merge(left, mid, nums, indexes, inversions)
    merge(mid+1, right, nums, indexes, inversions)

    mergeHelper(left, right, nums, indexes, inversions)
}

func mergeHelper(left int, right int, nums []int, indexes []int, inversions []int) {
    start, end := left, right
    mid := (left + right) / 2
    leftMid, rightMid := mid, mid+1

    newIndexes := make([]int, 0)
    smaller := int(0)

    for left <= leftMid && rightMid <= right {
        if nums[indexes[left]] <= nums[indexes[rightMid]] {
            newIndexes = append(newIndexes, indexes[left])
            inversions[indexes[left]] += smaller
            left++
        } else {
            newIndexes = append(newIndexes, indexes[rightMid])
            rightMid++
            smaller++
        }
    }

    for left <= leftMid {
        inversions[indexes[left]] += smaller
        newIndexes = append(newIndexes, indexes[left])
        left++
    }

    for rightMid <= right {
        newIndexes = append(newIndexes, indexes[rightMid])
        rightMid++
    }

    copy(indexes[start:end+1], newIndexes)
}

func main() {
	// result: [2,1,1,0]
	// nums := []int{5,2,6,1}

	// result: [0]
	// nums := []int{-1}

	// result: [0,0]
	nums := []int{-1,-1}

	// result: 
	// nums := []int{}

	result := countSmaller(nums)
	fmt.Printf("result = %v\n", result)
}

