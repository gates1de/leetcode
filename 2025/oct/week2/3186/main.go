package main

import (
	"fmt"
	"sort"
)

func maximumTotalDamage(power []int) int64 {
	counter := make(map[int]int)
	for _, p := range power {
		counter[p]++
	}

	keys := make([]int, 0, len(counter))
	for power := range counter {
		keys = append(keys, power)
	}

	sort.Ints(keys)
	arr := [][2]int{{-1000000000, 0}}
	for _, power := range keys {
		arr = append(arr, [2]int{power, counter[power]})
	}

	n := len(arr)
	f := make([]int64, n)
	maxNum := int64(0)
	result := int64(0)
	j := int(1)

	for i := 1; i < n; i++ {
		for j < i && arr[j][0] < arr[i][0] - 2 {
			if f[j] > maxNum {
				maxNum = f[j]
			}
			j++
		}

		f[i] = maxNum + int64(arr[i][0])*int64(arr[i][1])
		if f[i] > result {
			result = f[i]
		}
	}

	return result
}

func main() {
	// result: 6
	// power := []int{1,1,3,4}

	// result: 13
	power := []int{7,1,6,6}

	// result: 
	// power := []int{}

	result := maximumTotalDamage(power)
	fmt.Printf("result = %v\n", result)
}

