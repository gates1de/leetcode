package main
import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
    var result [][]int
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length - 3; i++ {
		if i != 0 && nums[i] == nums[i - 1] {
			continue
		}
		if nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3] > target {
			break
		}
		if nums[i] + nums[length - 1] + nums[length - 2] + nums[length - 2] < target {
			continue
		}
		for j := i + 1; j < length - 2; j++ {
			if j != i + 1 && nums[j] == nums[j - 1] {
				continue
			}
			if nums[i] + nums[j] + nums[j + 1] + nums[j + 2] > target {
				break
			}
			if nums[i] + nums[j] + nums[length - 1] + nums[length - 2] < target {
				continue
			}
			head, end := j + 1, length - 1
			for head < end {
				tempt := nums[i] + nums[j] + nums[head] + nums[end]
				if tempt == target {
					var lengthArray []int
					lengthArray = append(lengthArray, nums[i])
					lengthArray = append(lengthArray, nums[j])
					lengthArray = append(lengthArray, nums[head])
					lengthArray = append(lengthArray, nums[end])
					result = append(result, lengthArray)
					head++
					for head < end && nums[head] == nums[head - 1] {
						head++
					}
				} else if tempt > target {
					end--
				} else {
					head++
				}
			}
		}
	}
	return result
}

// Slow & More use memory
func mySolution(nums []int, target int) [][]int {
	result := [][]int{}
	if len(nums) < 4 {
		return result
	}

	m := map[int]int{}
	sortedNums := []int{}
	for _, num := range nums {
		if m[num] == 4 {
			continue
		}
		m[num]++
	}
	for num, count := range m {
		for i := 0; i < count; i++ {
			sortedNums = append(sortedNums, num)
		}
	}

	sort.Slice(sortedNums, func (a, b int) bool { return sortedNums[a] < sortedNums[b] })

	preNum1 := sortedNums[0] - 1
	for i := 0; i < len(sortedNums) - 3; i++ {
		num1 := sortedNums[i]
		if target > 0 && num1 > target {
			break
		}
		if preNum1 == num1 {
			continue
		}
		preNum1 = num1

		preNum2 := sortedNums[i + 1] - 1
		for j := i + 1; j < len(sortedNums) - 2; j++ {
			num2 := sortedNums[j]
			if target > 0 && num1 + num2 > target {
				break
			}
			if preNum2 == num2 {
				continue
			}
			preNum2 = num2

			preNum3 := sortedNums[j + 1] - 1
			for k := j + 1; k < len(sortedNums) - 1; k++ {
				num3 := sortedNums[k]
				if target > 0 && num1 + num2 + num3 > target {
					break
				}
				if preNum3 == num3 {
					continue
				}
				preNum3 = num3

				preNum4 := sortedNums[k + 1] - 1
				for l := k + 1; l < len(sortedNums); l++ {
					num4 := sortedNums[l]
					if preNum4 == num4 {
						continue
					}
					preNum4 = num4
					sum := num1 + num2 + num3 + num4
					// fmt.Printf("[%v, %v, %v, %v]\n", num1, num2, num3, num4)
					if sum == target {
						result = append(result, []int{num1, num2, num3, num4})
						break
					}
				}
			}
		}
	}
	return result
}

func main() {
	// result: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	// nums := []int{1, 0, -1, 0, -2, 2}
	// target := int(0)

	// result: [[2,2,2,2]]
	// nums := []int{2, 2, 2, 2, 2}
	// target := int(8)

	// result: [[0,0,0,0]]
	// nums := []int{0, 0, 0, 0}
	// target := int(0)

	// result: [[-5,-4,-3,1]]
	nums := []int{1, -2, -5, -4, -3, 3, 3, 5}
	target := int(-11)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := fourSum(nums, target)
	fmt.Printf("result = %v\n", result)
}

