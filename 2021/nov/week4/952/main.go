package main
import (
	"fmt"
	"sort"
)

func largestComponentSize(nums []int) int {
	sort.Ints(nums)
	primes := eratosthenes(nums[len(nums) - 1])
	// fmt.Printf("primes = %v\n", primes)

	m := map[int][]int{2: []int{}}
	for _, num := range nums {
		if num % 2 == 0 {
			m[2] = append(m[2], num)
		}

		// without primes[0] == 1
		for _, prime := range primes[1:] {
			if num < prime {
				break
			}
			if num % prime == 0 {
				if m[prime] == nil {
					m[prime] = []int{}
				}
				m[prime] = append(m[prime], num)
			}
		}
	}
	// fmt.Printf("m = %v\n", m)

	result := int(0)
	visited := map[int]bool{}
	for _, num := range nums {
		if visited[num] {
			continue
		}

		queue := []int{num}

		count := int(0)
		for len(queue) > 0 {
			target := queue[0]
			visited[target] = true
			count++
			if len(queue) >= 1 {
				queue = queue[1:]
			} else {
				queue = []int{}
			}

			for _, values := range m {
				if !contains(target, values) {
					continue
				}
				for _, val := range values {
					if target != val && !visited[val] {
						queue = append(queue, val)
						visited[val] = true
					}
				}
			}
		}

		if result < count {
			result = count
		}
		// fmt.Printf("num = %v, count = %v, visited = %v\n", num, count, visited)
	}

	return result
}

func contains(target int, nums []int) bool {
	for _, num := range nums {
		if target == num {
			return true
		}
	}
	return false
}

func eratosthenes(max int) []int {
	nums := make([]int, max + 1)

	for i := 2; i < len(nums); i++ {
		nums[i] = 1
	}

	count := int(1)
	for i := 2; i * i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == 0 {
				continue
			}

			if j % i == 0 {
				if nums[i] > 1 {
					count--
				}
				nums[j] = 0
			} else {
				if nums[i] == 1 {
					count++
				}
				nums[j] = j
			}
		}
	}

	result := make([]int, count)
	i := int(0)
	for _, num := range nums {
		if num == 0 {
			continue
		}
		result[i] = num
		i++
	}
	return result
}

func main() {
	// result: 4
	// nums := []int{4, 6, 15, 35}

	// result: 2 
	// nums := []int{20, 50, 9, 63}

	// result: 8
	nums := []int{2, 3, 6, 7, 4, 12, 21, 39}

	// result: 
	// nums := []int{}

	result := largestComponentSize(nums)
	fmt.Printf("result = %v\n", result)
}

