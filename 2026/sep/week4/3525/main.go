package main

import (
	"fmt"
)

type segment struct {
	product int
	prefix  [5]int
}

func mergeSegments(left segment, right segment, k int) segment {
	merged := segment{product: left.product * right.product % k}
	merged.prefix = left.prefix
	for remainder, count := range right.prefix[:k] {
		merged.prefix[left.product*remainder%k] += count
	}
	return merged
}

func resultArray(nums []int, k int, queries [][]int) []int {
	if k <= 0 {
		return []int{}
	}

	size := 1
	for size < len(nums) {
		size <<= 1
	}
	tree := make([]segment, size*2)
	identity := segment{product: 1 % k}
	for i := range tree {
		tree[i] = identity
	}

	for i, value := range nums {
		product := value % k
		tree[size+i] = segment{product: product}
		tree[size+i].prefix[product] = 1
	}
	for i := size - 1; i > 0; i-- {
		tree[i] = mergeSegments(tree[i*2], tree[i*2+1], k)
	}

	result := make([]int, len(queries))
	for queryIndex, query := range queries {
		index, value, start, remainder := query[0], query[1], query[2], query[3]
		position := size + index
		product := value % k
		tree[position] = segment{product: product}
		tree[position].prefix[product] = 1
		for position >>= 1; position > 0; position >>= 1 {
			tree[position] = mergeSegments(tree[position*2], tree[position*2+1], k)
		}

		left := size + start
		right := size + len(nums)
		leftResult := identity
		rightResult := identity
		for left < right {
			if left&1 == 1 {
				leftResult = mergeSegments(leftResult, tree[left], k)
				left++
			}
			if right&1 == 1 {
				right--
				rightResult = mergeSegments(tree[right], rightResult, k)
			}
			left >>= 1
			right >>= 1
		}

		merged := mergeSegments(leftResult, rightResult, k)
		result[queryIndex] = merged.prefix[remainder]
	}

	return result
}

func main() {
	// result: [2,2,2]
	// nums := []int{1,2,3,4,5}
	// k := int(3)
	// queries := [][]int{{2,2,0,2},{3,3,3,0},{0,1,0,1}}

	// result: [1,0]
	nums := []int{1,2,4,8,16,32}
	k := int(4)
	queries := [][]int{{0,2,0,2},{0,2,0,1}}

	// result: [5]
	// nums := []int{1,1,2,1,1}
	// k := int(2)
	// queries := [][]int{{2,1,0,1}}

	result := resultArray(nums, k, queries)
	fmt.Printf("result = %v\n", result)
}
