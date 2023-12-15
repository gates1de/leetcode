package main
import (
	"fmt"
)

func destCity(paths [][]string) string {
	hasOutgoing := make(map[string]bool)
	for _, path := range paths {
		cityA := path[0]
		hasOutgoing[cityA] = true
	}

	for _, path := range paths {
		cityB := path[1]
		if !hasOutgoing[cityB] {
			return cityB
		}
	}

	return ""
}

func main() {
	// result: "Sao Paulo" 
	// paths := [][]string{{"London","New York"},{"New York","Lima"},{"Lima","Sao Paulo"}}

	// result: "A"
	// paths := [][]string{{"B","C"},{"D","B"},{"C","A"}}

	// result: "Z"
	paths := [][]string{{"A","Z"}}

	// result: ""
	// paths := [][]string{}

	result := destCity(paths)
	fmt.Printf("result = %v\n", result)
}

