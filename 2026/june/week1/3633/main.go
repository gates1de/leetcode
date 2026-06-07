package main

import (
	"fmt"
)

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	if len(landStartTime) == 0 || len(waterStartTime) == 0 {
		return int(0)
	}

	result := int(^uint(0) >> 1)
	for i := range landStartTime {
		landFinish := landStartTime[i] + landDuration[i]
		for j := range waterStartTime {
			waterFinish := max(landFinish, waterStartTime[j]) + waterDuration[j]
			if waterFinish < result {
				result = waterFinish
			}
		}
	}

	for i := range waterStartTime {
		waterFinish := waterStartTime[i] + waterDuration[i]
		for j := range landStartTime {
			landFinish := max(waterFinish, landStartTime[j]) + landDuration[j]
			if landFinish < result {
				result = landFinish
			}
		}
	}

	return result
}

func main() {
	// result: 9
	// landStartTime := []int{2,8}
	// landDuration := []int{4,1}
	// waterStartTime := []int{6}
	// waterDuration := []int{3}

	// result: 14
	landStartTime := []int{5}
	landDuration := []int{3}
	waterStartTime := []int{1}
	waterDuration := []int{10}

	// result: 
	// landStartTime := []int{}
	// landDuration := []int{}
	// waterStartTime := []int{}
	// waterDuration := []int{}

	result := earliestFinishTime(landStartTime, landDuration, waterStartTime, waterDuration)
	fmt.Printf("result = %v\n", result)
}
