package main
import (
	"fmt"
)

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	if startFuel > target {
		return 0
	}

	dp := make([]int, len(stations) + 1)
	dp[0] = startFuel

	for i, station := range stations {
		location := station[0]
		getFuel := station[1]
		for j := i + 1; j > 0; j-- {
			if location <= dp[j - 1] {
				dp[j] = max(dp[j], dp[j - 1] + getFuel)
			}
		}
	}

	for i, v := range dp {
		if v >= target {
			return i
		}
	}

	return -1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// Wrong Answer
func ngSolution(target int, startFuel int, stations [][]int) int {
	if target <= startFuel {
		return 0
	}
	if len(stations) == 0 || startFuel < stations[0][0] {
		return -1
	}

	dp := map[int]int{}
	dp[target] = 500

	helper(0, startFuel, 0, target, stations, dp)

	fmt.Println(dp)
	if dp[target] == 500 {
		return -1
	}
	return dp[target]
}

func helper(position int, fuel int, count int, target int, stations [][]int, dp map[int]int) {
	if position + fuel >= target {
		if count < dp[target] {
			dp[target] = count
		}
		return
	}

	if val, ok := dp[position]; ok {
		if count < val {
			dp[position] = count
		} else {
			return
		}
	} else {
		dp[position] = count
	}

	for i, station := range stations {
		nextPosition := station[0]
		getFuel := station[1]
		distance := nextPosition - position
		if fuel < distance {
			break
		}
		helper(nextPosition, fuel - distance + getFuel, count + 1, target, stations[i + 1:], dp)
	}
}

func main() {
	// result: 0
	// target := int(1)
	// startFuel := int(1)
	// stations := [][]int{}

	// result: -1
	// target := int(100)
	// startFuel := int(1)
	// stations := [][]int{{10,100}}

	// result: 2
	// target := int(100)
	// startFuel := int(10)
	// stations := [][]int{{10,60},{20,30},{30,30},{60,40}}

	// result: 
	target := int(1000000000)
	startFuel := int(1000)
	stations := [][]int{{1,1000000000}}

	// result: 
	// target := int(0)
	// startFuel := int(0)
	// stations := [][]int{}

	result := minRefuelStops(target, startFuel, stations)
	fmt.Printf("result = %v\n", result)
}

