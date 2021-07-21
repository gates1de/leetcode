package main
import (
	"fmt"
)

// REF: https://leetcode.com/problems/three-equal-parts/discuss/468725/Go-O(n)-solution-with-detailed-explanation
func threeEqualParts(arr []int) []int {
	if len(arr) < 3 {
		return []int{-1, -1}
	}
	oneCount := int(0)
	for i := range arr {
		if arr[i] == 1 {
			oneCount++
		}
	}
	if oneCount == 0 {
		return []int{0, 2}
	}
	if oneCount % 3 != 0 {
		return []int{-1, -1}
	}

	counter := oneCount / 3
	r := len(arr) - 1
	rCounter := int(0)
	for rCounter < counter {
		if arr[r] == 1 {
			rCounter++
		}
		r--
	}
	lp := len(arr) - r - 1
	l := int(0)
	for arr[l] == 0 {
		l++
	}

	for i := 0; i < lp; i++ {
		if arr[l + i] != arr[r + 1 + i] {
			return []int{-1, -1}
		}
	}
	m := l + lp
	for arr[m] == 0 {
		m++
	}

	for i := 0; i < lp; i++ {
		if arr[m + i] != arr[r + 1 + i] {
			return []int{-1, -1}
		}
	}
	return []int{l + lp - 1, m + lp}
}

// Time Limit Exceeded
func ngSolution(arr []int) []int {
	result := []int{-1, -1}
	i := int(0)
	j := len(arr) - 1
	for i + 1 < j {
		if isOK(i, j, arr) {
			result[0] = i
			result[1] = j
			return result
		}

		j--
		if i + 1 == j {
			j = len(arr) - 1
			i++
		}
	}
	return result
}

func isOK(i int, j int, arr []int) bool {
	if i >= j {
		return false
	}
	first := arr[:i + 1]
	second := arr[i + 1:j]
	third := arr[j:]
	length := len(first)
	// fmt.Printf("first = %v, second = %v, third = %v\n", first, second, third)
	if length < len(second) {
		length = len(second)
	}
	if length < len(third) {
		length = len(third)
	}

	for k := 0; k <= length; k++ {
		firstVal := int(-1)
		secondVal := int(-1)
		thirdVal := int(-1)
		if k < len(first) {
			firstVal = first[len(first) - 1 - k]
		}
		if k < len(second) {
			secondVal = second[len(second) - 1 - k]
		}
		if k < len(third) {
			thirdVal = third[len(third) - 1 - k]
		}
		// fmt.Printf("fistVal = %v, secondVal = %v, thirdVal = %v\n", firstVal, secondVal, thirdVal)

		if (firstVal == -1 && (secondVal == 1 || thirdVal == 1)) ||
			(secondVal == -1 && (firstVal == 1 || thirdVal == 1)) ||
			(thirdVal == -1 && (firstVal == 1 || secondVal == 1)) ||
			(firstVal >= 0 && secondVal >= 0 && thirdVal >= 0 && (firstVal != secondVal || firstVal != thirdVal)) {
			return false
		}
	}
	return true
}

func main() {
	// result: [0, 3]
	// arr := []int{1, 0, 1, 0, 1}

	// result: [-1, -1]
	// arr := []int{1, 1, 0, 1, 1}

	// result: [0, 2]
	// arr := []int{1, 1, 0, 0, 1}

	// result: [-1, -1]
	arr := []int{1, 0, 1}

	// result: 
	// arr := []int{}

	result := threeEqualParts(arr)
	fmt.Printf("result = %v\n", result)
}

