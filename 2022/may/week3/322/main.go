package main
import (
	"fmt"
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
    
    for i := 1; i <= amount; i++ {
        minCoin := math.MaxInt32
        for _, coin := range coins {
            if i - coin >= 0 && dp[i - coin] != -1 {
                minCoin = min(minCoin, dp[i - coin] + 1)
            }
        }
        
        if minCoin == math.MaxInt32 {
            dp[i] = -1
        } else {
            dp[i] = minCoin
        }
    }
    
    return dp[amount]
}

func min(a int, b int) int {
    if a > b {
        return b
    }
    return a
}

// Time Limit Exceeded
func ngSolution(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 1 {
		if amount % coins[0] == 0 {
			return amount / coins[0]
		}
		return -1
	}

	m := map[int]bool{}
	for _, coin := range coins {
		if m[coin] || coin > amount {
			continue
		}
		m[coin] = true
	}

	uniqueCoins := make([]int, len(m))
	i := int(0)
	for coin, _ := range m {
		uniqueCoins[i] = coin
		i++
	}
	sort.Ints(uniqueCoins)

	tmp := make([]int, len(uniqueCoins))
	copy(tmp, uniqueCoins)
	maxLcm := lcmArr(tmp)
	maxCoin := uniqueCoins[len(uniqueCoins) - 1]
	defaultCount := (amount / maxLcm) * (maxLcm / maxCoin)

	result := math.MaxInt32
	dfs(0, 0, uniqueCoins, amount - (defaultCount * maxCoin), &result, map[int]int{})

	if result == math.MaxInt32 {
		return -1
	}

	return defaultCount + result
}

func gcd(a int, b int) int {
	if a % b == 0 {
		return b
	}
	return gcd(b, a % b)
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

func lcmArr(nums []int) int {
	for i := len(nums) - 2; i >= 0; i-- {
        nums[i] = lcm(nums[i], nums[i + 1])
    }
    return nums[0]
}

func dfs(current int, count int, coins []int, amount int, result *int, m map[int]int) {
	// fmt.Println(current, count, amount, *result)
	if current == amount {
		if *result > count {
			*result = count
		}
		return
	}

	if m[current] > 0 && m[current] < count {
		return
	}
	m[current] = count

	for i := len(coins) - 1; i >= 0; i-- {
		coin := coins[i]
		if current + coin > amount {
			continue
		}

		dfs(current + coin, count + 1, coins, amount, result, m)
	}
}

func main() {
	// result: 3
	// coins := []int{1, 2, 5}
	// amount := int(11)

	// result: -1
	// coins := []int{2}
	// amount := int(3)

	// result: 0
	// coins := []int{1}
	// amount := int(0)

	// result: 238
	// coins := []int{3,40,1,7,4,6,2}
	// amount := int(9382)

	// result: 159
	// coins := []int{3,40,1,7,4,6,2,60,5}
	// amount := int(9389)

	// result: 22
	// coins := []int{3,4,1,7,4,6,2,6,5}
	// amount := int(149)

	// result: 20
	coins := []int{186, 419, 83, 408}
	amount := int(6249)

	// result: 
	// coins := []int{}
	// amount := int(0)

	result := coinChange(coins, amount)
	fmt.Printf("result = %v\n", result)
}

