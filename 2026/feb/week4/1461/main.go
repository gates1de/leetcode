package main

import (
	"fmt"
)

func hasAllCodes(s string, k int) bool {
	req := 1 << k
	seen := make([]bool, req)
	mask := req - 1
	hash := int(0)

	arr := []byte(s)
	for i := range arr {
		hash = ((hash << 1) & mask) | (int(arr[i]) & 1)

		if i >= k - 1 && !seen[hash] {
			seen[hash] = true
			req--
			if req == 0 {
				return true
			}
		}
	}

	return false
}

func main() {
	// result: true
	// s := "00110110"
	// k := int(2)

	// result: true
	// s := "0110"
	// k := int(1)

	// result: false
	s := "0110"
	k := int(2)

	// result: 
	// s := ""
	// k := int(0)

	result := hasAllCodes(s, k)
	fmt.Printf("result = %v\n", result)
}

