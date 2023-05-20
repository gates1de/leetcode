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

    result := make([]float64, len(queries))
    TOP:
    for i, query := range queries {
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
	// values := []float64{2.0,3.0}
	// queries := [][]string{{"a","c"},{"b","a"},{"a","e"},{"a","a"},{"x","x"}}

	// result: [3.75000,0.40000,5.00000,0.20000]
	// equations := [][]string{{"a","b"},{"b","c"},{"bc","cd"}}
	// values := []float64{1.5,2.5,5.0}
	// queries := [][]string{{"a","c"},{"c","b"},{"bc","cd"},{"cd","bc"}}

	// result: [0.50000,2.00000,-1.00000,-1.00000]
	equations := [][]string{{"a","b"}}
	values := []float64{0.5}
	queries := [][]string{{"a","b"},{"b","a"},{"a","c"},{"x","y"}}

	// result: 
	// equations := [][]string{}
	// values := []float64{}
	// queries := [][]string{}

	result := calcEquation(equations, values, queries)
	fmt.Printf("result = %v\n", result)
}

