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

	for i, s := range stations {
		currentMile := s[0]
		fuel := s[1]
		for j := i + 1; j > 0; j-- {
			if currentMile <= dp[j - 1] {
				dp[j] = max(dp[j], dp[j - 1] + fuel)
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
	if a > b {
		return a
	}
	return b
}

// Time Limit Exceeded (Infinite Loop)
func mySolution(target int, startFuel int, stations [][]int) int {
	if startFuel >= target {
		return 0
	}
	if len(stations) == 0 {
		return -1
	}
	return helper(startFuel, 0, 0, stations[0], target, stations[1:])
}

func helper(currentFuel int, count int, currentMile int, now []int, target int, stations [][]int) int {
	// fmt.Printf("currentFuel = %v, now = %v, count = %v\n", currentFuel, now, count)
	if currentFuel + currentMile >= target {
		return count
	}

	if currentFuel < now[0] - currentMile {
		return -1
	}

	expendFuel := now[0] - currentMile
	stopFuel := currentFuel - expendFuel + now[1]
	nonStopFuel := currentFuel - expendFuel
	currentMile = now[0]
	if len(stations) == 0 {
		if stopFuel + currentMile >= target {
			return count + 1
		}
		return -1
	}

	stop := helper(stopFuel, count + 1, currentMile, stations[0], target, stations[1:])
	nonStop := helper(nonStopFuel, count, currentMile, stations[0], target, stations[1:])

	if stop == -1 {
		return nonStop
	} else if nonStop == -1 {
		return stop
	}

	if nonStop < stop {
		return nonStop
	}
	return stop
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 0
	// target := int(1)
	// startFuel := int(1)
	// stations := [][]int{}

	// result: -1
	// target := int(100)
	// startFuel := int(1)
	// stations := [][]int{{10, 100}}

	// result: 2
	// target := int(100)
	// startFuel := int(10)
	// stations := [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}

	// result: 4
	target := int(1000)
	startFuel := int(10)
	stations := [][]int{{7,217},{10,211},{17,146},{21,232},{25,310},{48,175},{53,23},{63,158},{71,292},{96,85},{100,302},{102,295},{116,110},{122,46},{131,20},{132,19},{141,13},{163,85},{169,263},{179,233},{191,169},{215,163},{224,231},{228,282},{256,115},{259,3},{266,245},{283,331},{299,21},{310,224},{315,188},{328,147},{345,74},{350,49},{379,79},{387,276},{391,92},{405,174},{428,307},{446,205},{448,226},{452,275},{475,325},{492,310},{496,94},{499,313},{500,315},{511,137},{515,180},{519,6},{533,206},{536,262},{553,326},{561,103},{564,115},{582,161},{593,236},{599,216},{611,141},{625,137},{626,231},{628,66},{646,197},{665,103},{668,8},{691,329},{699,246},{703,94},{724,277},{729,75},{735,23},{740,228},{761,73},{770,120},{773,82},{774,297},{780,184},{791,239},{801,85},{805,156},{837,157},{844,259},{849,2},{860,115},{874,311},{877,172},{881,109},{888,321},{894,302},{899,321},{908,76},{916,241},{924,301},{933,56},{960,29},{979,319},{983,325},{988,190},{995,299},{996,97}}

	// result: 
	// target := int(0)
	// startFuel := int(0)
	// stations := [][]int{}

	result := minRefuelStops(target, startFuel, stations)
	fmt.Printf("result = %v\n", result)
}

