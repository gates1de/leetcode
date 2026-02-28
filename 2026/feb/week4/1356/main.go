package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
    sort.Slice(arr, func(a, b int) bool {
        aCount := getBitOnesCount(arr[a])
        bCount := getBitOnesCount(arr[b])
        if aCount == bCount {
            return arr[a] < arr[b]
        }
        return aCount < bCount
    })
    return arr
}

func getBitOnesCount(val int) (count int) {
    for val != 0 {
        count += val & 0x1
        val >>= 1
    }
    return
}

func main() {
	// result: [0,1,2,4,8,3,5,6,7]
	// arr := []int{0,1,2,3,4,5,6,7,8}

	// result: [1,2,4,8,16,32,64,128,256,512,1024]
	arr := []int{1024,512,256,128,64,32,16,8,4,2,1}

	// result: 
	// arr := []int{}

	result := sortByBits(arr)
	fmt.Printf("result = %v\n", result)
}

