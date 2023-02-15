package main
import (
	"fmt"
)

func addToArrayForm(num []int, k int) []int {
	kNums := make([]int, 0, 5)
	for k > 0 {
		digitNum := k % 10
		k /= 10
		kNums = append(kNums, 0)
		copy(kNums[1:], kNums)
		kNums[0] = digitNum
	}

	for len(num) > len(kNums) {
		kNums = append(kNums, 0)
		copy(kNums[1:], kNums)
		kNums[0] = 0
	}
	for len(kNums) > len(num) {
		num = append(num, 0)
		copy(num[1:], num)
		num[0] = 0
	}

	carry := int(0)
	for i := len(kNums) - 1; i >= 0; i-- {
		originalNum := num[i]
		kNum := kNums[i]
		addedNum := originalNum + kNum + carry
		if addedNum >= 10 {
			carry = 1
			addedNum -= 10
		} else {
			carry = 0
		}
		num[i] = addedNum
	}

	if carry == 1 {
		num = append(num, 0)
		copy(num[1:], num)
		num[0] = 1
	}

	return num
}

func main() {
	// result: [1,2,3,4]
	// num := []int{1,2,0,0}
	// k := int(34)

	// result: [4,5,5]
	// num := []int{2,7,4}
	// k := int(181)

	// result: [1,0,2,1]
	num := []int{2,1,5}
	k := int(806)

	// result: 
	// num := []int{}
	// k := int(0)

	result := addToArrayForm(num, k)
	fmt.Printf("result = %v\n", result)
}

