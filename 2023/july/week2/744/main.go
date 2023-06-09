package main
import (
	"fmt"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	if n == 0 {
		return 0
	}

	targetIndex := binarySearch(0, letters, target)
	if targetIndex == n {
		return letters[0]
	}

	return letters[targetIndex]
}

func binarySearch(index int, letters []byte, target byte) int {
	n := len(letters)
	left := int(0)
	right := n - 1

	for left <= right {
		mid := (left + right) / 2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	// result: "c"
	// letters := []byte{'c','f','j'}
	// target := byte('a')

	// result: "f"
	// letters := []byte{'c','f','j'}
	// target := byte('c')

	// result: "f"
	// letters := []byte{'c','f','j'}
	// target := byte('d')

	// result: "x"
	letters := []byte{'x','x','y','y'}
	target := byte('z')

	// result: 
	// letters := []byte{}
	// target := byte('')

	result := nextGreatestLetter(letters, target)
	fmt.Printf("result = %v (%v)\n", result, string(result))
}

