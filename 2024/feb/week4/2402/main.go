package main
import (
	"fmt"
	"math"
	"sort"
)

func mostBooked(n int, meetings [][]int) int {
	meetingCount := make([]int, n)
	roomAvailabilityTime := make([]int64, n)
	sort.Slice(meetings, func (a, b int) bool {
		return meetings[a][0] < meetings[b][0]
	})

	for _, meeting := range meetings {
		start := int64(meeting[0])
		end := int64(meeting[1])
		minRoomAvailabilityTime := int64(math.MaxInt64)
		minAvailabilityTimeRoom := int64(0)
		isFoundUnusedRoom := false

		for i, _ := range meetingCount {
			if roomAvailabilityTime[i] <= start {
				isFoundUnusedRoom = true
				meetingCount[i]++
				roomAvailabilityTime[i] = end
				break
			}

			if minRoomAvailabilityTime > roomAvailabilityTime[i] {
				minRoomAvailabilityTime = roomAvailabilityTime[i]
				minAvailabilityTimeRoom = int64(i)
			}
		}

		if !isFoundUnusedRoom {
			roomAvailabilityTime[minAvailabilityTimeRoom] += end - start
			meetingCount[minAvailabilityTimeRoom]++
		}
	}

	maxMeetingCount := int(0)
	result := int(0)

	for i, count := range meetingCount {
		if count > maxMeetingCount {
			maxMeetingCount = meetingCount[i]
			result = i
		}
	}

	return result
}

func main() {
	// result: 0
	// n := int(2)
	// meetings := [][]int{{0,10},{1,5},{2,7},{3,4}}

	// result: 1
	n := int(3)
	meetings := [][]int{{1,20},{2,10},{3,5},{4,9},{6,8}}

	// result: 
	// n := int(0)
	// meetings := [][]int{}

	result := mostBooked(n, meetings)
	fmt.Printf("result = %v\n", result)
}

