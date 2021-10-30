package main
import (
	"fmt"
)

func sortColors(nums []int)  {
	quickSort(0, len(nums) - 1, nums)
}

func quickSort(l int, r int, nums []int) {
	// fmt.Printf("l = %v, r = %v\n", l, r)
	if l >= r {
		return
	}

	i := l
	j := r
	p := int(1)

	for true {
		for nums[i] < p {
			i++
			if i > r {
				return
			}
		}

		for nums[j] > p {
			j--
			if j < l {
				return
			}
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	quickSort(l, i - 1, nums)
	quickSort(j + 1, r, nums)
}

func main() {
	// result: [0,0,1,1,2,2]
	// nums := []int{2, 0, 2, 1, 1, 0}

	// result: [0,1,2]
	// nums := []int{2, 0, 1}

	// result: [0]
	// nums := []int{0}

	// result: [1]
	// nums := []int{1}

	// result: [0,0,0,0,0,0,0,0,0,1,1,1,1,2,2,2,2,2,2]
	nums := []int{2,1,0,2,0,1,2,0,1,0,2,1,0,2,0,0,2,0,0}

	// result: 
	// nums := []int{}

	sortColors(nums)
	fmt.Printf("result = %v\n", nums)
}

