package main
import (
	"fmt"
	"sort"
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
	i := int(0)
	result := []string{}

	for i < len(strs) && !(m == 0 && n == 0) {
		sort.Slice(strs, func (a, b int) bool {
			if len(strs[a]) == 1 || len(strs[b]) == 1 {
				return len(strs[a]) < len(strs[b])
			}
			if len(strs[a]) > m + n || len(strs[b]) > m + n {
				return false
			}

			mCountA := strings.Count(strs[a], "0")
			nCountA := strings.Count(strs[a], "1")
			mCountB := strings.Count(strs[b], "0")
			nCountB := strings.Count(strs[b], "1")

			if mCountA <= mCountB && nCountA <= nCountB {
				return true
			}
			if m >= n {
				return nCountA <= nCountB
			}
			return mCountA < mCountB
		})

		fmt.Printf("strs = %v\n", strs)
		str := strs[i]

		if str == "0" && m > 0 {
			strs = append(strs[:i], strs[i + 1:]...)
			result = append(result, str)
			m--
			continue
		} else if str == "1" && n > 0 {
			strs = append(strs[:i], strs[i + 1:]...)
			result = append(result, str)
			n--
			continue
		}

		mCount := strings.Count(str, "0")
		nCount := strings.Count(str, "1")

		fmt.Printf("m = %v, n = %v, str = %v\n", m, n, str)
		if len(str) > m + n {
			break
		}

		if mCount > m || nCount > n {
			i++
			continue
		}
		strs = append(strs[:i], strs[i + 1:]...)
		result = append(result, str)
		m -= mCount
		n -= nCount
		i = 0
		// fmt.Printf("m = %v, n = %v, result = %v\n", m, n, result)
	}

	return len(result)
}

func main() {
	// result: 4
	// strs := []string{"10", "0001", "111001", "1", "0"}
	// m := int(5)
	// n := int(3)

	// result: 2
	// strs := []string{"10", "0", "1"}
	// m := int(1)
	// n := int(1)

	// result: 3
	// strs := []string{"10", "0001", "111001", "1", "0"}
	// m := int(4)
	// n := int(3)

	// result: 3
	// strs := []string{"111", "1000", "1000", "1000"}
	// m := int(9)
	// n := int(3)

	// result: 0
	// strs := []string{"00", "000"}
	// m := int(1)
	// n := int(10)

	// result: 4
	// strs := []string{"0", "0", "1", "1"}
	// m := int(2)
	// n := int(2)

	// result: 3
	// strs := []string{"10", "0001", "111001", "1", "0"}
	// m := int(3)
	// n := int(2)

	// result: 2
	strs := []string{"011111", "001", "001"}
	m := int(4)
	n := int(5)

	result := findMaxForm(strs, m, n)
	fmt.Printf("result = %v\n", result)
}

