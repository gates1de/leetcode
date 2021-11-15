package main
import (
	"fmt"
	"math"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)
    max := int(1)
    l := len(nums)
    dp := make([]int, len(nums))
    for i := range dp {
		dp[i] = 1
	}
    for i := 1; i < l; i++ {
        for j := 0; j < i; j++ {
            if nums[i] % nums[j] == 0 && dp[j] + 1 > dp[i] {
                dp[i] = dp[j] + 1
                if dp[i] > max {
					max = dp[i]
				}
            }
        }
    }

    result := []int{}
    prev := int(-1)
    for i := l - 1; i >= 0; i-- {
        if dp[i] == max && (prev == -1 || prev % nums[i] == 0) {
            result = append(result, nums[i])
            prev = nums[i]
            max--
        }
    }
    return result
}

// Wrong Answer
const maxNum = 2 * 1000000000
func ngSolution(nums []int) []int {
	m := map[int][]int{}
	isExistsOne := false
	for _, num := range nums {
		if num == 1 {
			isExistsOne = true
			continue
		}
		p := prime(num)
		if m[p] == nil {
			m[p] = []int{}
		}
		m[p] = append(m[p], num)
	}

	// fmt.Printf("m = %v\n", m)

	longestPrime := int(0)
	longestPrimes := []int{}
	for p, arr := range m {
		if len(longestPrimes) < len(arr) {
			longestPrime = p
			longestPrimes = arr
		}
	}

	result := []int{}
	if len(longestPrimes) == 1 {
		result = longestPrimes
		if isExistsOne {
			result = append([]int{1}, result...)
		}
		return result
	}

	// fmt.Printf("longestPrimes = %v\n", longestPrimes)

	isOKMap := map[int]bool{}
	for _, n := range longestPrimes {
		isOKMap[n] = true
	}
	for _, n1 := range longestPrimes {
		if n1 == longestPrime {
			continue
		}
		if !isOKMap[n1] {
			continue
		}
		for _, n2 := range longestPrimes {
			if n1 % n2 != 0 && n2 % n1 != 0 {
				isOKMap[n2] = false
				continue
			}
		}
	}

	for key, isOK := range isOKMap {
		if !isOK {
			continue
		}
		result = append(result, key)
	}

	if isExistsOne {
		result = append([]int{1}, result...)
	}

	return result
}

func prime(n int) int {
    if n <= 2 {
        return n
    }
	if n % 2 == 0 {
		isTwo := true
		for i := n; i > 1; i /= 2 {
			if i % 2 != 0 {
				isTwo = false
				break
			}
		}
		if isTwo {
			return 2
		}
	}

    root := int(math.Sqrt(float64(n)))
	// fmt.Printf("n = %v, root = %v\n", n, root)
    i := int(3)
    for i <= root {
        if n % i == 0 {
			return i
        }
        i += 2
    }
    return n
}

func main() {
	// result: [1,2]
	// nums := []int{1, 2, 3}

	// result: [1,2,4,8]
	// nums := []int{1, 2, 4, 8}

	// result: [1,3,6]
	// nums := []int{1, 3, 4, 5, 6, 9}

	// result: [100,1000,10000]
	// nums := []int{100, 1000, 10000, 111111, 2222222}

	// result: [3,6,18]
	nums := []int{3, 6, 9, 18, 27}

	// result: 
	// nums := []int{}

	result := largestDivisibleSubset(nums)
	fmt.Printf("result = %v\n", result)
}

