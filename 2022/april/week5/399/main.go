package main
import (
	"fmt"
)

type Right struct {
	Num float64
	Var string
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	rightsMap := map[string][]Right{}
	for i, e := range equations {
		val := values[i]
		if rightsMap[e[0]] == nil {
			rightsMap[e[0]] = []Right{Right{Num: 1, Var: e[0]}}
		}
		queue := []Right{Right{Num: val, Var: e[1]}}
		searched := map[string]bool{}
		for len(queue) > 0 {
			target := queue[0]
			queue = queue[1:]
			rightsMap[e[0]] = append(rightsMap[e[0]], target)
			for j, eq := range equations {
				if searched[eq[0]] && searched[eq[1]] {
					continue
				}

				if target.Var == eq[0] {
					queue = append(queue, Right{Num: target.Num * values[j], Var: eq[1]})
					searched[eq[0]] = true
				} else if target.Var == eq[1] {
					queue = append(queue, Right{Num: target.Num / values[j], Var: eq[0]})
					searched[eq[1]] = true
				}
			}
		}

		if rightsMap[e[1]] == nil {
			rightsMap[e[1]] = []Right{Right{Num: 1, Var: e[1]}}
		}
		queue = []Right{Right{Num: 1 / val, Var: e[0]}}
		searched = map[string]bool{}
		for len(queue) > 0 {
			target := queue[0]
			queue = queue[1:]
			rightsMap[e[1]] = append(rightsMap[e[1]], target)
			for j, eq := range equations {
				if searched[eq[0]] && searched[eq[1]] {
					continue
				}

				if target.Var == eq[0] {
					queue = append(queue, Right{Num: target.Num * values[j], Var: eq[1]})
					searched[eq[0]] = true
				} else if target.Var == eq[1] {
					queue = append(queue, Right{Num: target.Num / values[j], Var: eq[0]})
					searched[eq[1]] = true
				}
			}
		}
	}

	// fmt.Println(rightsMap)

	result := make([]float64, len(queries))
	TOP:
	for i, query := range queries {
		// fmt.Printf("------------ %v -----------\n", query)
		if query[0] == query[1] {
			if rightsMap[query[0]] != nil {
				result[i] = 1
			} else {
				result[i] = -1
			}
			continue
		}

		for _, otherRight1 := range rightsMap[query[0]] {
			for _, otherRight2 := range rightsMap[query[1]] {
				if otherRight1.Var == otherRight2.Var {
					result[i] = otherRight1.Num / otherRight2.Num
					continue TOP
				}
			}
		}

		result[i] = -1
	}

	return result
}

func main() {
	// result: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
	// equations := [][]string{{"a","b"},{"b","c"}}
	// values := []float64{2, 3}
	// queries := [][]string{{"a","c"},{"b","a"},{"a","e"},{"a","a"},{"x","x"}}

	// result: [3.75000,0.40000,5.00000,0.20000]
	// equations := [][]string{{"a","b"},{"b","c"},{"bc","cd"}}
	// values := []float64{1.5, 2.5, 5}
	// queries := [][]string{{"a","c"},{"c","b"},{"bc","cd"},{"cd","bc"}}

	// result: [0.50000,2.00000,-1.00000,-1.00000]
	// equations := [][]string{{"a","b"}}
	// values := []float64{0.5}
	// queries := [][]string{{"a","b"},{"b","a"},{"a","c"},{"x","y"}}

	// result: [360.00000,0.00833,20.00000,1.00000,-1.00000,-1.00000]
	// equations := [][]string{{"x1","x2"},{"x2","x3"},{"x3","x4"},{"x4","x5"}}
	// values := []float64{3, 4, 5, 6}
	// queries := [][]string{{"x1","x5"},{"x5","x2"},{"x2","x4"},{"x2","x2"},{"x2","x9"},{"x9","x9"}}

	// result: [0.29412,10.94800,1.00000,1.00000,-1.00000,-1.00000,0.71429]
	equations := [][]string{{"a","b"},{"e","f"},{"b","e"}}
	values := []float64{3.4,1.4,2.3}
	queries := [][]string{{"b","a"},{"a","f"},{"f","f"},{"e","e"},{"c","c"},{"a","c"},{"f","e"}}

	// result: 
	// equations := [][]string{}
	// values := []float64{}
	// queries := [][]string{}

	result := calcEquation(equations, values, queries)
	fmt.Printf("result = %v\n", result)
}

