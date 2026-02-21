package main

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	result := make([]string, 0)

	for h := range 12 {
		for m := range 60 {
			if bits.OnesCount8(uint8(h)) + bits.OnesCount8(uint8(m)) == turnedOn {
				result = append(result, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return result
}

func main() {
	// result: ["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
	// turnedOn := int(1)

	// result: 
	turnedOn := int(9)

	// result: 
	// turnedOn := int(0)

	result := readBinaryWatch(turnedOn)
	fmt.Printf("result = %v\n", result)
}

