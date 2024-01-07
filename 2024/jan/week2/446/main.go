package main
import (
	"fmt"
)

func numberOfArithmeticSlices(nums []int) int {
    if len(nums) < 3 {
        return 0
    }

    result := int(0)
    m := map[int]map[int]int{}
    for i := 0; i < len(nums); i++ {
        m[i] = map[int]int{}
        for j := 0; j < i; j++ {
            diff := nums[i] - nums[j]
            c1 := m[i][diff]
            c2 := m[j][diff]
            result += c2
            m[i][diff] = c1 + c2 + 1
        }
    }

    return result
}

func main() {
	// result: 7
	// nums := []int{2,4,6,8,10}

	// result: 16
	nums := []int{7,7,7,7,7}

	// result: 
	// nums := []int{}

	result := numberOfArithmeticSlices(nums)
	fmt.Printf("result = %v\n", result)
}

