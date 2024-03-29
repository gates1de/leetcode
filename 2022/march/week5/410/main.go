package main
import (
	"fmt"
	"math"
	"sort"
)

func splitArray(nums []int, m int) int {
    length := len(nums)
    dp := make([][]int, length + 1)
    for i := 0; i <= length; i++ {
        dp[i] = make([]int, m + 1)
        for j := 0; j <= m; j++ {
            dp[i][j] = math.MaxInt32
        }
    }

    sums := make([]int, length + 1)
    for i := 0; i < length; i++ {
        sums[i + 1] = sums[i] + nums[i]
    }

    dp[0][0] = 0
    for i := 1; i <= length; i++ {
        for j := 1; j <= m; j++ {
            for k := 0; k < i; k++ {
                dp[i][j] = min(dp[i][j], max(dp[k][j - 1], sums[i] - sums[k]))
            }
        }
    }

    return dp[length][m]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Wrong Answer
func ngSolution(nums []int, m int) int {
	n := len(nums)
	sum := rangeSum(nums, 0, len(nums) - 1)
	result := sum
	if m == 1 {
		return result
	} else if len(nums) == m {
		sort.Ints(nums)
		return nums[len(nums) - 1]
	}

	for i := 0; i < len(nums) - m + 1; i++ {
		for j := 0; j < n - i; j++ {
			// fmt.Printf("range (%v, %v)\n", j, j + i)
			if m == 2 && j > 0 && j + i != n - 1 {
				continue
			}

			ran := rangeSum(nums, j, j + i)
			// fmt.Println(sum, ran, sum - ran)
			if ran <= sum - ran && sum - ran < result {
				result = sum - ran
			}
		}
	}

	return result
}

func rangeSum(nums []int, start int, end int) int {
	result := int(0)
	if start < 0 || len(nums) <= start || start > end {
		return result
	}

	for i := start; i <= end; i++ {
		result += nums[i]
	}

	return result
}

func main() {
	// result: 18
	// nums := []int{7, 2, 5, 10, 8}
	// m := int(2)

	// result: 9
	// nums := []int{1, 2, 3, 4, 5}
	// m := int(2)

	// result: 4
	// nums := []int{1, 4, 4}
	// m := int(3)

	// result: 47
	nums := []int{4,9,0,3,5,1,6,47,2,3,8,9,1,6,4,0,2,9,3,8,1,4}
	m := int(5)

	// result: 
	// nums := []int{}
	// m := int(0)

	result := splitArray(nums, m)
	fmt.Printf("result = %v\n", result)
}

