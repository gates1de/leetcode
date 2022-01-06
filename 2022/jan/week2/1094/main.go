package main
import (
	"fmt"
)

func carPooling(trips [][]int, capacity int) bool {
	mTrip := map[int][][]int{}
	for _, trip := range trips {
		from := trip[1]
		if mTrip[from] == nil {
			mTrip[from] = [][]int{}
		}
		mTrip[from] = append(mTrip[from], trip)
	}

	currentCap := int(0)
	drops := map[int]int{}
	for i := 0; i <= 1000; i++ {
		if drops[i] > 0 {
			currentCap -= drops[i]
		}

		startTrips := mTrip[i]
		if len(startTrips) == 0 {
			continue
		}

		for _, trip := range startTrips {
			currentCap += trip[0]
			drops[trip[2]] += trip[0]
		}

		if currentCap > capacity {
			return false
		}
	}

	return true
}

func main() {
	// result: false
	// trips := [][]int{{2,1,5},{3,3,7}}
	// capacity := int(4)

	// result: true
	// trips := [][]int{{2,1,5},{3,3,7}}
	// capacity := int(5)

	// result: true
	// trips := [][]int{{2,1,3},{3,3,7}}
	// capacity := int(4)

	// result: false
	// trips := [][]int{{100,1,1000}}
	// capacity := int(99)

	// result: false
	// trips := [][]int{{3,5,6},{2,4,7},{7,1,9}}
	// capacity := int(10)

	// result: false
	trips := [][]int{{8,2,3},{4,1,3},{1,3,6},{8,4,6},{4,4,8}}
	capacity := int(12)

	// result: 
	// trips := [][]int{}
	// capacity := int(0)

	result := carPooling(trips, capacity)
	fmt.Printf("result = %v\n", result)
}

