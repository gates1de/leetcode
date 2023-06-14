package main
import (
	"fmt"
)

func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return []string{}
    }

    result := []string{}
    start := nums[0]
    pre := nums[0]
    for _, num := range nums[1:] {
        if num - pre == 1 {
            pre = num
            continue
        }

        if start == pre {
            result = append(result, fmt.Sprintf("%v", start))
        } else {
            result = append(result, fmt.Sprintf("%v->%v", start, pre))
        }
        start = num
        pre = num
    }

    if start == pre {
        result = append(result, fmt.Sprintf("%v", start))
    } else {
        result = append(result, fmt.Sprintf("%v->%v", start, pre))
    }
    return result
}

func main() {
	// result: ["0->2","4->5","7"]
	// nums := []int{0,1,2,4,5,7}

	// result: ["0","2->4","6","8->9"]
	nums := []int{0,2,3,4,6,8,9}

	// result: 
	// nums := []int{}

	result := summaryRanges(nums)
	fmt.Printf("result = %v\n", result)
}

