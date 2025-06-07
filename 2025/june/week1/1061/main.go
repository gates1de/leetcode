package main
import (
	"fmt"
)

type UnionFind struct {
	Parents []int
	Count  int
}

func Constructor(num int) *UnionFind {
	parents := make([]int, num)
	for i := 0; i < num; i++ {
		parents[i] = i
	}

	return &UnionFind{
		Parents: parents,
		Count: num,
	}
}

func (u *UnionFind) Union(a, b int) bool {
	parentA := u.Find(a)
	parentB := u.Find(b)

	if parentA != parentB {
		if parentA < parentB {
			u.Parents[parentB] = parentA
		} else {
			u.Parents[parentA] = parentB
		}

		u.Count--
		return true
	}

	return false
}

func (u *UnionFind) Find(a int) int {
	if u.Parents[a] == a {
		return a
	}

	u.Parents[a] = u.Find(u.Parents[a])
	return u.Parents[a]
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	uf := Constructor(26)
	for i := 0; i < len(s1); i++ {
		uf.Union(int(s1[i] - 'a'), int(s2[i] - 'a'))
	}

	result := ""
	for i := 0; i < len(baseStr); i++ {
		result += string('a' + uf.Find(int(baseStr[i] - 'a')))
	}

	return result
}

func main() {
	// result: "makkek"
	// s1 := "parker"
	// s2 := "morris"
	// baseStr := "parser"

	// result: "hdld"
	// s1 := "hello"
	// s2 := "world"
	// baseStr := "hold"

	// result: "aauaaaaada"
	s1 := "leetcode"
	s2 := "programs"
	baseStr := "sourcecode"

	// result: 
	// s1 := ""
	// s2 := ""
	// baseStr := ""

	result := smallestEquivalentString(s1, s2, baseStr)
	fmt.Printf("result = %v\n", result)
}

