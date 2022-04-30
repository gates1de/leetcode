package main
import (
	"fmt"
	"sort"
)


func smallestStringWithSwaps(s string, pairs [][]int) string {
    result := []byte(s)
    ds := make([]int, len(s))
    for i := 0; i < len(ds); i++ {
        ds[i] = i
    }

    m := make([][]int, len(s))
    for _, p := range pairs {
        merge(ds, p[0], p[1])
    }

    for i, _ := range s {
        u := find(ds, i)
        m[u] = append(m[u], i)
    }

    for _, idxs := range m {
        bs := []byte{}
        for _, idx := range idxs {
            bs = append(bs, s[idx])
        }

        sort.Slice(bs, func(i, j int) bool {
            return bs[i] < bs[j]
        })

        for i, idx := range idxs {
            result[idx] = bs[i]
        }
    }

    return string(result)
}

func find(ds []int, u int) int {
    if ds[u] == u {
        return u
    }
    ds[u] = find(ds, ds[u])
    return ds[u]
}

func merge(ds []int, u, v int) {
    u = find(ds, u)
    v = find(ds, v)
    if u != v {
        ds[u] = v
    }
}

type IndexedRune struct {
	Rune rune
	Index int
}

// Time Limit Exceeded
func ngSolution(s string, pairs [][]int) string {
	components := [][]int{}
	added := map[int]bool{}
	for i, _ := range s {
		if added[i] {
			continue
		}

		queue := []int{i}
		nums := []int{}
		for len(queue) > 0 {
			num := queue[0]
			queue = queue[1:]
			nums = append(nums, num)
			added[num] = true

			for _, pair := range pairs {
				if num == pair[0] && !added[pair[1]] {
					added[pair[1]] = true
					queue = append(queue, pair[1])
				} else if num == pair[1] && !added[pair[0]] {
					added[pair[0]] = true
					queue = append(queue, pair[0])
				}
			}
		}

		components = append(components, nums)
	}

	result := make([]rune, len(s))
	for _, c := range components {
		sort.Ints(c)
		str := make([]IndexedRune, len(c))
		for i, sIndex := range c {
			str[i] = IndexedRune{Rune: rune(s[sIndex]), Index: sIndex}
		}
		sort.Slice(str, func(a, b int) bool { return str[a].Rune < str[b].Rune })
		for j, index := range c {
			ir := str[j]
			result[index] = rune(s[ir.Index])
		}
	}

	return string(result)
}

func main() {
	// result: "bacd"
	// s := "dcab"
	// pairs := [][]int{{0, 3}, {1, 2}}

	// result: "abcd"
	// s := "dcab"
	// pairs := [][]int{{0, 3}, {1, 2}, {0, 2}}

	// result: "abc"
	s := "cba"
	pairs := [][]int{{0, 1}, {1, 2}}

	// result: 
	// s := ""
	// pairs := [][]int{}

	result := smallestStringWithSwaps(s, pairs)
	fmt.Printf("result = %v\n", result)
}

