package main
import (
	"fmt"
)

type MountainArray struct {
	arr []int
}

func (this *MountainArray) get(index int) int {
	return this.arr[index]
}

func (this *MountainArray) length() int {
	return len(this.arr)
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	length := mountainArr.length()
	low := int(1)
	high := length - 2

	for low != high {
		testIndex := (low + high) / 2

		if mountainArr.get(testIndex) < mountainArr.get(testIndex + 1) {
			low = testIndex + 1
		} else {
			high = testIndex
		}
	}

	peakIndex := low
	low = 0
	high = peakIndex

	for low != high {
		testIndex := (low + high) / 2

		if mountainArr.get(testIndex) < target {
			low = testIndex + 1
		} else {
			high = testIndex
		}
	}

	if mountainArr.get(low) == target {
		return low
	}

	low = peakIndex + 1
	high = length - 1

	for low != high {
		testIndex := (low + high) / 2

		if mountainArr.get(testIndex) > target {
			low = testIndex + 1
		} else {
			high = testIndex
		}
	}

	if mountainArr.get(low) == target {
		return low
	}

	return -1
}

func main() {
	// result: 2
	// target := int(3)
	// mountainArr := &MountainArray{arr: []int{1,2,3,4,5,3,1}}

	// result: -1
	target := int(3)
	mountainArr := &MountainArray{arr: []int{0,1,2,4,2,1}}

	// result: 
	// target := int(0)
	// mountainArr := &MountainArray{arr: []int{}}

	result := findInMountainArray(target, mountainArr)
	fmt.Printf("result = %v\n", result)
}

