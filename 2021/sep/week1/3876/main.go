package main
import (
	"fmt"
	"sort"
)

func slowestKey(releaseTimes []int, keysPressed string) byte {
	maxTime := releaseTimes[0]
	minKeys := []rune{rune(keysPressed[0])}
	for i := 1; i < len(releaseTimes); i++ {
		releaseTime := releaseTimes[i] - releaseTimes[i - 1]
		if maxTime < releaseTime {
			maxTime = releaseTime
			minKeys = []rune{rune(keysPressed[i])}
		} else if maxTime == releaseTime {
			minKeys = append(minKeys, rune(keysPressed[i]))
		}
	}
	fmt.Printf("maxTime = %v, minKeys = %v\n", maxTime, minKeys)
	sort.Slice(minKeys, func (a, b int) bool { return minKeys[a] > minKeys[b] })
	return byte(minKeys[0])
}

func main() {
	// result: c
	// releaseTimes := []int{9, 29, 49, 50}
	// keyPressed := "cbcd"

	// result: a
	// releaseTimes := []int{12, 23, 36, 46, 62}
	// keyPressed := "spuda"

	// result: 
	releaseTimes := []int{8, 16, 20, 24, 26, 34, 40}
	keyPressed := "abcdefg"

	// result: 
	// releaseTimes := []int{}
	// keyPressed := ""

	// result: 
	// releaseTimes := []int{}
	// keyPressed := ""

	result := slowestKey(releaseTimes, keyPressed)
	fmt.Printf("result = %v\n", result)
}

