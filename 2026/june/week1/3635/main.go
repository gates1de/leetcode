package main

import (
	"fmt"
)

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	if len(landStartTime) == 0 || len(waterStartTime) == 0 {
		return int(0)
	}

	minLandFinish := landStartTime[0] + landDuration[0]
	for i := 1; i < len(landStartTime); i++ {
		finish := landStartTime[i] + landDuration[i]
		if finish < minLandFinish {
			minLandFinish = finish
		}
	}

	minWaterFinish := waterStartTime[0] + waterDuration[0]
	for i := 1; i < len(waterStartTime); i++ {
		finish := waterStartTime[i] + waterDuration[i]
		if finish < minWaterFinish {
			minWaterFinish = finish
		}
	}

	result := int(^uint(0) >> 1)
	for i := range waterStartTime {
		finish := max(minLandFinish, waterStartTime[i]) + waterDuration[i]
		if finish < result {
			result = finish
		}
	}

	for i := range landStartTime {
		finish := max(minWaterFinish, landStartTime[i]) + landDuration[i]
		if finish < result {
			result = finish
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
