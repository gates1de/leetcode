package main
import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
    if len(nums1) > len(nums2){
        nums1, nums2 = nums2, nums1
    }

    key := make(map[int]int)
    var list []int
    for _,val:=range nums1{
        key[val]++
    }

    for _,val:=range nums2{
        if _,found := key[val]; found{
            if key[val] > 0{
                list = append(list, val)
            }
            key[val]--
        }
    }
    return list
}

// Slow & Use more memory
func mySolution(nums1 []int, nums2 []int) []int {
	result := []int{}
	m1 := map[int]int{}
	m2 := map[int]int{}
	for _, n := range nums1 {
		m1[n]++
	}
	for _, n := range nums2 {
		m2[n]++
	}

	longM := m1
	shortM := m2
	if len(m1) < len(m2) {
		longM = m2
		shortM = m1
	}

	// fmt.Printf("longM = %v, shortM = %v\n", longM, shortM)
	for n, count1 := range longM {
		count2 := shortM[n]
		if count1 == 0 || count2 == 0 {
			continue
		}
		loopCount := min(count1, count2)
		for i := 0; i < loopCount; i++ {
			result = append(result, n)
		}
	}
	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: [2,2]
	// nums1 := []int{1, 2, 2, 1}
	// nums2 := []int{2, 2}

	// result: [4, 9]
	// nums1 := []int{4, 9, 5}
	// nums2 := []int{9, 4, 9, 8, 4}

	// result: [4,1,2]
	nums1 := []int{4, 1, 2}
	nums2 := []int{2, 1, 1, 4, 5}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := intersect(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

