package main

import (
	"fmt"
)

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	arriveAliceDay := toDay(arriveAlice)
	leaveAliceDay := toDay(leaveAlice)
	arriveBobDay := toDay(arriveBob)
	leaveBobDay := toDay(leaveBob)

	start := max(arriveAliceDay, arriveBobDay)
	end := min(leaveAliceDay, leaveBobDay)

	result := end - start + 1
	if result < 0 {
		return 0
	}

	return result
}

func toDay(date string) int {
	daysBeforeMonth := [...]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	month := int(date[0]-'0')*10 + int(date[1]-'0')
	day := int(date[3]-'0')*10 + int(date[4]-'0')

	return daysBeforeMonth[month-1] + day
}

func main() {
	// result: 3
	// arriveAlice := "08-15"
	// leaveAlice := "08-18"
	// arriveBob := "08-16"
	// leaveBob := "08-19"

	// result: 0
	arriveAlice := "10-01"
	leaveAlice := "10-31"
	arriveBob := "11-01"
	leaveBob := "12-31"

	// result: 0
	// arriveAlice := ""
	// leaveAlice := ""
	// arriveBob := ""
	// leaveBob := ""

	result := countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob)
	fmt.Printf("result = %v\n", result)
}
