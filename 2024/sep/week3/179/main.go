package main
import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Chars []string

func (this Chars) Len() int {
	return len(this)
}

func (this Chars) Less(i, j int) bool {
	return this[i] + this[j] > this[j] + this[i]
}

func (this Chars) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	flag := true

	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
		flag = flag && (nums[i] == 0)
	}

	if flag {
		return "0"
	}

	sort.Sort(Chars(strs))
	return strings.Join(strs, "")
}

func main() {
	// result: "210"
	// nums := []int{10,2}

	// result: "9534330"
	nums := []int{3,30,34,5,9}

	// result: ""
	// nums := []int{}

	result := largestNumber(nums)
	fmt.Printf("result = %v\n", result)
}

