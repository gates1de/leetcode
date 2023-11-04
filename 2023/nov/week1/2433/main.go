package main
import (
	"fmt"
)

func findArray(pref []int) []int {
	n := len(pref)
	for i := n - 1; i > 0; i-- {
		pref[i] = pref[i] ^ pref[i - 1]
	}
	return pref
}

func main() {
	// result: [5,7,2,3,2]
	// pref := []int{5,2,0,3,1}

	// result: [13]
	pref := []int{13}

	// result: []
	// pref := []int{}

	result := findArray(pref)
	fmt.Printf("result = %v\n", result)
}

