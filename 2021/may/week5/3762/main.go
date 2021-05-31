package main
import (
	"fmt"
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
    // sort strings O(nlogn)
    sort.Strings(products)

    ans := make([][]string, 0)

    for i := 0; i < len(searchWord); i++ {
        query := string(searchWord[:i+1])

        // binary search for minimum index i s.t. products[i] >= q
        // O(logn)
        target := sort.Search(len(products), func(i int) bool {
            return products[i] >= query
        })

        result := make([]string, 0)

        j, k := 0, target
        for j < 3 && k < len(products) && strings.Contains(products[k], query) {
            // ensure query is a substring of products[k]
            result = append(result, products[k])
            j += 1
            k += 1
        }

        ans = append(ans, result)
    }

    return ans
}

// A little slower
func mySolution(products []string, searchWord string) [][]string {
	result := make([][]string, len(searchWord))
	sort.Strings(products)
	for i, _ := range searchWord {
		newProducts := []string{}
		for _, product := range products {
			// fmt.Printf("searchWord[:i + 1] = %v, product[:i + 1] = %v\n", searchWord[:i + 1], product[:i + 1])
			if searchWord[:i + 1] == product[:min(i + 1, len(product))] {
				newProducts = append(newProducts, product)
			}
			if len(newProducts) == 3 {
				break
			}
		}
		fmt.Printf("newProducts = %v\n", newProducts)
		result[i] = newProducts
	}
	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: [
	//		     ["mobile","moneypot","monitor"],
	//		 	 ["mobile","moneypot","monitor"],
	//		 	 ["mouse","mousepad"],
	//		 	 ["mouse","mousepad"],
	//		 	 ["mouse","mousepad"]
	//		   ]
	// products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	// searchWord := "mouse"

	// result: [
	//		 	 ["havana"],
	//		 	 ["havana"],
	//		 	 ["havana"],
	//		 	 ["havana"],
	//		 	 ["havana"],
	//		 	 ["havana"],
	//		 	 ["havana"],
	//		   ]
	// products := []string{"havana"}
	// searchWord := "havana"

	// result: [
	//			 ["baggage","bags","banner"],
	//			 ["baggage","bags","banner"],
	//			 ["baggage","bags"],
	//			 ["bags"]
	//		   ]
	products := []string{"bags", "baggage", "banner", "box", "cloths", "bag"}
	searchWord := "bags"

	// result: [[[],[],[],[],[],[],[]]]
	// products := []string{"havana"}
	// searchWord := "tatiana"

	// result: 
	// products := []string{}
	// searchWord := ""

	result := suggestedProducts(products, searchWord)
	fmt.Printf("result = %v\n", result)
}

