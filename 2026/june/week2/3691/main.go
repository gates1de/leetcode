package main

import (
	"fmt"
)

// You are given an integer array nums of length n and an integer k.
//
// You must select exactly k distinct non-empty subarrays nums[l..r] of nums. Subarrays may overlap, but the exact same subarray (same l and r) cannot be chosen more than once.
//
// The value of a subarray nums[l..r] is defined as: max(nums[l..r]) - min(nums[l..r]).
//
// The total value is the sum of the values of all chosen subarrays.
//
// Return the maximum possible total value you can achieve.
//
// Constraints:
//
// 1 <= n == nums.length <= 5 * 10​​​​​​​4
// 0 <= nums[i] <= 109
// 1 <= k <= min(105, n * (n + 1) / 2)


func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)
	if n == 0 || k == 0 {
		return int64(0)
	}

	segmentTree := NewSegmentTree(nums)
	heap := make([]Interval, 0, k*2+1)
	seen := make(map[uint64]struct{}, k*2+1)

	root := Interval{
		Left:  int(0),
		Right: n - 1,
		Value: int64(segmentTree.QueryMax(int(0), n-1) - segmentTree.QueryMin(int(0), n-1)),
	}
	pushInterval(&heap, root)
	seen[intervalKey(root.Left, root.Right)] = struct{}{}

	result := int64(0)
	for count := int(0); count < k && len(heap) > 0; count++ {
		current := popInterval(&heap)
		result += current.Value

		if current.Left < current.Right {
			leftChild := Interval{
				Left:  current.Left + 1,
				Right: current.Right,
				Value: int64(segmentTree.QueryMax(current.Left + 1, current.Right) - segmentTree.QueryMin(current.Left + 1, current.Right)),
			}
			leftKey := intervalKey(leftChild.Left, leftChild.Right)
			if _, exists := seen[leftKey]; !exists {
				seen[leftKey] = struct{}{}
				pushInterval(&heap, leftChild)
			}

			rightChild := Interval{
				Left:  current.Left,
				Right: current.Right - 1,
				Value: int64(segmentTree.QueryMax(current.Left, current.Right - 1) - segmentTree.QueryMin(current.Left, current.Right - 1)),
			}
			rightKey := intervalKey(rightChild.Left, rightChild.Right)
			if _, exists := seen[rightKey]; !exists {
				seen[rightKey] = struct{}{}
				pushInterval(&heap, rightChild)
			}
		}
	}

	return result
}

type Interval struct {
	Left  int
	Right int
	Value int64
}

type SegmentTree struct {
	Size    int
	MaxTree []int
	MinTree []int
}

func NewSegmentTree(nums []int) *SegmentTree {
	size := int(1)
	n := len(nums)
	for size < n {
		size <<= 1
	}

	maxTree := make([]int, size*2)
	minTree := make([]int, size*2)
	inf := int(^uint(0) >> 1)
	for i := range minTree {
		minTree[i] = inf
	}

	for i, num := range nums {
		index := size + i
		maxTree[index] = num
		minTree[index] = num
	}

	for index := size - 1; index > 0; index-- {
		left := index << 1
		right := left | 1

		if maxTree[left] > maxTree[right] {
			maxTree[index] = maxTree[left]
		} else {
			maxTree[index] = maxTree[right]
		}

		if minTree[left] < minTree[right] {
			minTree[index] = minTree[left]
		} else {
			minTree[index] = minTree[right]
		}
	}

	return &SegmentTree{
		Size:    size,
		MaxTree: maxTree,
		MinTree: minTree,
	}
}

func (s *SegmentTree) QueryMax(left int, right int) int {
	left += s.Size
	right += s.Size
	result := int(0)

	for left <= right {
		if left&1 == 1 {
			if s.MaxTree[left] > result {
				result = s.MaxTree[left]
			}
			left++
		}
		if right&1 == 0 {
			if s.MaxTree[right] > result {
				result = s.MaxTree[right]
			}
			right--
		}
		left >>= 1
		right >>= 1
	}

	return result
}

func (s *SegmentTree) QueryMin(left int, right int) int {
	left += s.Size
	right += s.Size
	result := int(^uint(0) >> 1)

	for left <= right {
		if left&1 == 1 {
			if s.MinTree[left] < result {
				result = s.MinTree[left]
			}
			left++
		}
		if right&1 == 0 {
			if s.MinTree[right] < result {
				result = s.MinTree[right]
			}
			right--
		}
		left >>= 1
		right >>= 1
	}

	return result
}

func intervalKey(left int, right int) uint64 {
	return (uint64(left) << 32) | uint64(uint32(right))
}

func pushInterval(heap *[]Interval, value Interval) {
	*heap = append(*heap, value)
	index := len(*heap) - 1

	for index > 0 {
		parent := (index - 1) / 2
		if (*heap)[parent].Value >= (*heap)[index].Value {
			break
		}
		(*heap)[parent], (*heap)[index] = (*heap)[index], (*heap)[parent]
		index = parent
	}
}

func popInterval(heap *[]Interval) Interval {
	result := (*heap)[0]
	last := len(*heap) - 1

	if last == 0 {
		*heap = (*heap)[:0]
		return result
	}

	(*heap)[0] = (*heap)[last]
	*heap = (*heap)[:last]

	index := int(0)
	for {
		left := index*2 + 1
		right := left + 1
		largest := index

		if left < len(*heap) && (*heap)[left].Value > (*heap)[largest].Value {
			largest = left
		}
		if right < len(*heap) && (*heap)[right].Value > (*heap)[largest].Value {
			largest = right
		}
		if largest == index {
			break
		}

		(*heap)[index], (*heap)[largest] = (*heap)[largest], (*heap)[index]
		index = largest
	}

	return result
}

func main() {
	// result: 4
	// nums := []int{1,3,2}
	// k := int(2)

	// result: 12
	nums := []int{4,2,5,1}
	k := int(3)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maxTotalValue(nums, k)
	fmt.Printf("result = %v\n", result)
}
