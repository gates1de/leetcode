package main
import (
	"fmt"
)

const MAX = int(100000)

type DSU struct {
	Arr []int
	Size []int
}

func (this *DSU) Find(x int) int {
	if this.Arr[x] == x {
		return x
	}

	this.Arr[x] = this.Find(this.Arr[x])
	return this.Arr[x]
}

func (this *DSU) Merge(x int, y int) {
	fx := this.Find(x)
	fy := this.Find(y)

	if fx == fy {
		return
	}

	if this.Size[fx] > this.Size[fy] {
		fx, fy = fy, fx
	}
	this.Arr[fx] = fy
	this.Size[fy] += this.Size[fx]
}

func constructor(n int) DSU {
	arr := make([]int, n + 1)
	size := make([]int, n + 1)
	for i, _ := range arr {
		arr[i] = i
		size[i] = 1
	}
	return DSU{Arr: arr, Size: size}
}

func canTraverseAllPairs(nums []int) bool {
	n := len(nums)
	has := make([]bool, MAX + 1)
	for _, num := range nums {
		has[num] = true
	}

	if n == 1 {
		return true
	}

	if has[1] {
		return false
	}

	sieve := make([]int, MAX + 1)
	for d := 2; d <= MAX; d++ {
		if sieve[d] == 0 {
			for v := d; v <= MAX; v += d {
				sieve[v] = d
			}
		}
	}

	union := constructor(2 * MAX + 1)

	for _, num := range nums {
		val := num
		for val > 1 {
			prime := sieve[val]
			root := prime + MAX
			if union.Find(root) != union.Find(num) {
				union.Merge(root, num)
			}

			for val % prime == 0 {
				val /= prime
			}
		}
	}

	count := int(0)
	for i := 2; i <= MAX; i++ {
		if has[i] && union.Find(i) == i {
			count++
		}
	}

	return count == 1
}

func main() {
	// result: true
	// nums := []int{2,3,6}

	// result: false
	// nums := []int{3,9,5}

	// result: true
	nums := []int{4,3,12,8}

	// result: 
	// nums := []int{}

	result := canTraverseAllPairs(nums)
	fmt.Printf("result = %v\n", result)
}

