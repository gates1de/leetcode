package main
import (
	"fmt"
)

func isPossible(nums []int) bool {
	subs := [][]int{}

    TOP:
    for _, num := range nums {
        for i := len(subs) - 1; i >= 0; i-- {
            if subs[i][len(subs[i]) - 1] + 1 == num {
                subs[i] = append(subs[i], num)
                continue TOP
            }
        }

        subs = append(subs, []int{num})
    }
    
    for _, sub := range subs {
        if len(sub) < 3 {
            return false
        }
    }

    return true
}

// Wrong Answer
func ngSolution(nums []int) bool {
	subs := [][]int{{nums[0]}}
	for _, num := range nums[1:] {
		isAdded := false
		for i, sub := range subs {
			if len(sub) < 3 && sub[len(sub) - 1] == num - 1 {
				subs[i] = append(subs[i], num)
				isAdded = true
				break
			}
		}

		if !isAdded {
			newSub := []int{num}
			subs = append(subs, newSub)
		}
	}

	remains := []int{}
	for i, sub := range subs {
		if len(sub) < 3 {
			for _, num := range sub {
				remains = append(remains, num)
			}
			subs[i] = []int{}
		}
	}

	TOP:
	for _, num := range remains {
		for i, sub := range subs {
			if len(sub) == 0 {
				continue
			}

			if sub[len(sub) - 1] == num - 1 {
				subs[i] = append(subs[i], num)
				continue TOP
			}
		}

		fmt.Println(num, subs)
		return false
	}

	for _, sub := range subs {
		if 0 < len(sub) && len(sub) < 3 {
			fmt.Println(sub)
			return false
		}
	}

	fmt.Println(subs)
	return true
}

func main() {
	// result: true
	// nums := []int{1,2,3,3,4,5}

	// result: true
	// nums := []int{1,2,3,3,4,4,5,5}

	// result: false
	// nums := []int{1,2,3,4,4,5}

	// result: true
	nums := []int{1,2,3,4,5,5,6,7}

	// result: 
	// nums := []int{}

	result := isPossible(nums)
	fmt.Printf("result = %v\n", result)
}

