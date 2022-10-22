package main
import (
	"fmt"
	"math"
)

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
    if n < d {
        return -1
    }

    memo := make([][]int, d)
    for i := 0; i < d; i++ {
        memo[i] = make([]int, n)
		for j, _ := range memo[i] {
            memo[i][j] = -1
        }
    }
    
    return helper(0, jobDifficulty, d - 1, memo)
}

func helper(start int, jobDifficulty []int, d int, memo [][]int) int {
    maxDifficulty := int(0)
    
    if memo[d][start] != -1 {
        return memo[d][start]
    }
    
    if d == 0 {
        for i := start; i < len(jobDifficulty); i++ {
            maxDifficulty = max(maxDifficulty, jobDifficulty[i])
        }
        return maxDifficulty
    }
    
    minDifficulty := math.MaxInt32
    
    for i := start; i < len(jobDifficulty) - d; i++ {
        maxDifficulty = max(maxDifficulty, jobDifficulty[i])
        minDifficulty = min(minDifficulty, maxDifficulty + helper(i + 1, jobDifficulty, d - 1, memo))
    }
    
    memo[d][start] = minDifficulty
    
    return minDifficulty
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}

// Time Limit Exceeded
func ngSolution(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	} else if n == d {
		result := int(0)
		for _, num := range jobDifficulty {
			result += num
		}
		return result
	}

	result := math.MaxInt32
	dp(0, 0, jobDifficulty, d, &result)
	return result
}

func dp(index int, sum int, jd []int, d int, result *int) {
	// fmt.Println(index, sum, d, *result)
	if index == len(jd) {
		if sum < *result {
			*result = sum
		}
		return
	}

	if sum >= *result {
		return
	}

	maxDifficulty := int(0)

	if d == 1 {
		for _, num := range jd[index:] {
			if num > maxDifficulty {
				maxDifficulty = num
			}
		}

		if sum + maxDifficulty < *result || *result == math.MaxInt32 {
			*result = sum + maxDifficulty
		}
		return
	}

	for i := index; i <= len(jd) - d; i++ {
		num := jd[i]
		if num > maxDifficulty {
			maxDifficulty = num
		}

		dp(i + 1, sum + maxDifficulty, jd, d - 1, result)
	}
}

func main() {
	// result: 7
	// jobDifficulty := []int{6,5,4,3,2,1}
	// d := int(2)

	// result: -1
	// jobDifficulty := []int{9,9,9}
	// d := int(4)

	// result: 3
	// jobDifficulty := []int{1,1,1}
	// d := int(3)

	// result: 16
	jobDifficulty := []int{6,5,4,3,2,1,9}
	d := int(3)

	// result: 
	// jobDifficulty := []int{}
	// d := int(0)

	result := minDifficulty(jobDifficulty, d)
	fmt.Printf("result = %v\n", result)
}

