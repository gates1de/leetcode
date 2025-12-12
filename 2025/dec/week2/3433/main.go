package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	sort.Slice(events, func(i, j int) bool {
		timeA, _ := strconv.Atoi(events[i][1])
		timeB, _ := strconv.Atoi(events[j][1])
		if timeA != timeB {
			return timeA < timeB
		}
		return events[i][0] != "MESSAGE" && events[j][0] == "MESSAGE"
	})

	result := make([]int, numberOfUsers)
	nextOnlineTime := make([]int, numberOfUsers)

	for _, event := range events {
		curTime, _ := strconv.Atoi(event[1])
		eventType := event[0]

		if eventType == "MESSAGE" {
			target := event[2]
			switch target {
			case "ALL":
				for i := range numberOfUsers {
					result[i]++
				}
			case "HERE":
				for i := range numberOfUsers {
					if nextOnlineTime[i] <= curTime {
						result[i]++
					}
				}
			default:
				users := strings.Split(target, " ")
				for _, user := range users {
					index, _ := strconv.Atoi(user[2:])
					result[index]++
				}
			}
		} else {
			index, _ := strconv.Atoi(event[2])
			nextOnlineTime[index] = curTime + 60
		}
	}

	return result
}

func main() {
	// result: [2,2]
	// numberOfUsers := int(2)
	// events := [][]string{{"MESSAGE","10","id1 id0"},{"OFFLINE","11","0"},{"MESSAGE","71","HERE"}}

	// result: [2,2]
	// numberOfUsers := int(2)
	// events := [][]string{{"MESSAGE","10","id1 id0"},{"OFFLINE","11","0"},{"MESSAGE","12","ALL"}}

	// result: [0,1]
	numberOfUsers := int(2)
	events := [][]string{{"OFFLINE","10","0"},{"MESSAGE","12","HERE"}}

	// result: []
	// numberOfUsers := int(0)
	// events := [][]string{}

	result := countMentions(numberOfUsers, events)
	fmt.Printf("result = %v\n", result)
}

