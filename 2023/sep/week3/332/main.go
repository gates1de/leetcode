package main
import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string)

	for _, ticket := range tickets {
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
	}

	for key := range graph {
		sort.Sort(sort.Reverse(sort.StringSlice(graph[key])))
	}

	result := make([]string, 0, 300)
	dfs("JFK", &result, graph)

	for i := 0; i < len(result) / 2; i++ {
		result[i], result[len(result) - 1 - i] = result[len(result) - 1 - i], result[i]
	}

	return result
}

func dfs(airport string, itinerary *[]string, graph map[string][]string) {
	for len(graph[airport]) > 0 {
		next := graph[airport][len(graph[airport]) - 1]
		graph[airport] = graph[airport][:len(graph[airport]) - 1]
		dfs(next, itinerary, graph)
	}
	*itinerary = append(*itinerary, airport)
}

func main() {
	// result: ["JFK","MUC","LHR","SFO","SJC"]
	// tickets := [][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}

	// result: ["JFK","ATL","JFK","SFO","ATL","SFO"]
	tickets := [][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}}

	// result: 
	// tickets := [][]string{}

	result := findItinerary(tickets)
	fmt.Printf("result = %v\n", result)
}

