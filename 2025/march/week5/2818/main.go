package main
import (
	"fmt"
	"math"
	heap "container/heap"
)

const modulo = int(1e9 + 7)

type Heap [][]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] > h[j][0]
}
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Peek() []int        { return h[0] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func maximumScore(nums []int, k int) int {
	n := len(nums)
	primeScores := make([]int, n)

	for i, num := range nums {
		for factor := 2; int(factor) <= int(math.Sqrt(float64(num))); factor++ {
			if num % factor == 0 {
				primeScores[i]++
				for num % factor == 0 {
					num /= factor
				}
			}
		}
		
		if num >= 2 {
			primeScores[i]++
		}
	}

	nextDominant := make([]int, n)
	prevDominant := make([]int, n)
	for i, _ := range nextDominant {
		nextDominant[i] = n
	}
	for i, _ := range prevDominant {
		prevDominant[i] = -1
	}

	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && primeScores[stack[len(stack) - 1]] < primeScores[i] {
			topIndex := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			nextDominant[topIndex] = i
		}

		if len(stack) > 0 {
			prevDominant[i] = stack[len(stack) - 1]
		}
		stack = append(stack, i)
	}

	numOfSubarrays := make([]int64, n)
	for i, _ := range numOfSubarrays {
		numOfSubarrays[i] = int64(nextDominant[i] - i) * int64(i - prevDominant[i])
	}

	pq := &Heap{}
	heap.Init(pq)

	for i, num := range nums {
		heap.Push(pq, []int{num, i})
	}

	result := int64(1)

	for k > 0 {
		top := heap.Pop(pq).([]int)
		num := top[0]
		i := top[1]

		operations := min(k, int(numOfSubarrays[i]))
		result = (result * power(int64(num), int64(operations))) % int64(modulo)
		k -= operations
	}

	return int(result)
}

func power(base int64, exponent int64) int64 {
	result := int64(1)

	for exponent > 0 {
		if exponent % 2 == 1 {
			result = (result * base) % int64(modulo)
		}

		base = (base * base) % int64(modulo)
		exponent /= 2
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
	// result: 81
	// nums := []int{8,3,9,3,8}
	// k := int(2)

	// result: 4788
	nums := []int{19,12,14,6,10,18}
	k := int(3)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maximumScore(nums, k)
	fmt.Printf("result = %v\n", result)
}

