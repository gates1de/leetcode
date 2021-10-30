package main
import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
    res := [][]int{}

    if len(nums) < 3 {
        return res
    }

    sort.Ints(nums)
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        secondIdx := i + 1
        thirdIdx := len(nums) - 1
        for secondIdx < thirdIdx {
            sum := nums[i] + nums[secondIdx] + nums[thirdIdx]
            if sum < 0 {
                secondIdx++
            } else if sum > 0 {
                thirdIdx--
            } else {
                res = append(res, []int{nums[i], nums[secondIdx], nums[thirdIdx]})
                for secondIdx < thirdIdx && nums[secondIdx] == nums[secondIdx + 1] {
                    secondIdx++
                }
                for secondIdx < thirdIdx && nums[thirdIdx] == nums[thirdIdx - 1] {
                    thirdIdx--
                }
                secondIdx++
                thirdIdx--
            }
        }
    }

    return res
}

// Slow & Use more memory
func mySolution(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums) - 1]

	added := map[[3]int]bool{}
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if -sum < min || max < -sum {
				continue
			}

			for k := j + 1; k < len(nums); k++ {
				// fmt.Printf("i = %v, j = %v, k = %v\n", nums[i], nums[j], nums[k])
				if nums[k] > -sum {
					break
				} else if nums[k] < -sum {
					continue
				}

				if added[[3]int{nums[i], nums[j], nums[k]}] {
					break
				}
				result = append(result, []int{nums[i], nums[j], nums[k]})
				added[[3]int{nums[i], nums[j], nums[k]}] = true
			}
		}
	}

	return result
}

func main() {
	// result: [[-1,-1,2],[-1,0,1]]
	// nums := []int{-1, 0, 1, 2, -1, -4}

	// result: []
	// nums := []int{}

	// result: []
	// nums := []int{0}

	// result: 
	nums := []int{2,-3,1,6,-5,7,-8,-3,0,4,5,-1,8,-4,4,5,-2,7,-11,4,8}

	// result: 
	// nums := []int{}

	result := threeSum(nums)
	fmt.Printf("result = %v\n", result)
}

