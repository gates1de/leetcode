package main
import (
	"fmt"
	"sort"
)

func equationsPossible(equations []string) bool {
	sort.Slice(equations, func(a, b int) bool {
		return equations[a][1] > equations[b][1]
	})

	uf := make([]int, 27)
	for i, _ := range uf {
		uf[i] = -1
	}

	for _, equation := range equations {
		var1 := int(equation[0]) - int('a')
		var2 := int(equation[3]) - int('a')
		operator := equation[1:3]
		if operator == "==" {
			uf[unionFind(uf, var1)] = unionFind(uf, var2)
		} else if operator == "!=" && unionFind(uf, var1) == unionFind(uf, var2) {
			return false
		}
	}
	return true
}

func unionFind(uf []int, i int) int {
	if uf[i] == -1 || uf[i] == i {
		return i
	}
	return unionFind(uf, uf[i])
}

func main() {
	// result: false
	// equations := []string{"a==b","b!=a"}

	// result: true
	// equations := []string{"b==a","a==b"}

	// result: true
	// equations := []string{"a==b","b==a","a==c"}

	// result: false
	// equations := []string{"a==b","b==a","a==c","c!=a"}

	// result: false
	// equations := []string{"a==b","b!=c","c==a"}

	// result: false
	// equations := []string{"a!=a"}

	// result: false
	// equations := []string{"a==b","e==c","b==c","a!=e"}

	// result: true
	equations := []string{"b==b","b==e","e==c","d!=e"}

	// result: 
	// equations := []string{}

	result := equationsPossible(equations)
	fmt.Printf("result = %v\n", result)
}

