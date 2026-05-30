package main

import (
	"fmt"
)

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if n == 0 || s[0] != '0' || s[n - 1] != '0' {
		return false
	}

	reachable := make([]bool, n)
	reachable[0] = true

	windowCount := int(0)
	for i := 1; i < n; i++ {
		add := i - minJump
		if add >= 0 && reachable[add] {
			windowCount++
		}

		remove := i - maxJump - 1
		if remove >= 0 && reachable[remove] {
			windowCount--
		}

		reachable[i] = s[i] == '0' && windowCount > 0
	}

	return reachable[n - 1]
}

func main() {
	// result: true
	// s := "011010"
	// minJump := int(2)
	// maxJump := int(3)

	// result: false
	s := "01101110"
	minJump := int(2)
	maxJump := int(3)

	// result:
	// s := ""
	// minJump := int(0)
	// maxJump := int(0)

	result := canReach(s, minJump, maxJump)
	fmt.Printf("result = %v\n", result)
}
