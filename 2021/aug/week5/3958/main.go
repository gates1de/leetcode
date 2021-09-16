package main
import (
	"fmt"
)

func findMin(nums []int) int {
    l := 0
    r := len(nums) - 1
    for l < r {
        m := (r + l) / 2
        if nums[l] < nums[r] {
            return nums[l]
        }

        if m > 0 && nums[m] < nums[m - 1] {
            return nums[m]
        }

        if nums[r] < nums[m] {
            l = m + 1
        } else {
            r = m
        }
    }
    return nums[l]
}

// Slow & Use more memory
func mySolution(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] < nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	return search(nums)
}

func search(nums []int) int {
   high := len(nums) - 1
   low := 0
   if nums[low] < nums[high] {
	   return nums[low]
   }

   var mid int
   for low < high {
      mid = (high + low) / 2
	  // fmt.Printf("nums[low] = %v, nums[mid] = %v, nums[high] = %v\n", nums[low], nums[mid], nums[high])
	  if mid == 0 || mid == len(nums) - 1 {
		  return nums[mid]
	  }
	  if nums[mid - 1] > nums[mid] {
		  return nums[mid]
	  }
	  if nums[mid] > nums[mid + 1] {
		  return nums[mid + 1]
	  }

	  if nums[low] < nums[mid] {
		  low = mid
	  } else {
		  high = mid
	  }
   }
   return 0
}

func main() {
	// result: 1
	// nums := []int{3, 4, 5, 1, 2}

	// result: 0
	// nums := []int{4, 5, 6, 7, 0, 1, 2}

	// result: 11
	// nums := []int{11, 13, 15, 17}

	// result: -3300
	// nums := []int{400,500,600,700,800,900,1000,1100,2200,-3300,-440,-85,-78,-77}

	// result: 1
	nums := []int{2, 1}

	// result: 
	// nums := []int{}

	result := findMin(nums)
	fmt.Printf("result = %v\n", result)
}

