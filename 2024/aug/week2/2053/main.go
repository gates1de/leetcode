package main
import (
	"fmt"
)

func kthDistinct(arr []string, k int) string {
	counter := make(map[string]int)

	for _, str := range arr {
		counter[str]++
	}

	for _, str := range arr {
		if counter[str] == 1 {
			k--
		}

		if k == 0 {
			return str
		}
	}

	return ""
}

func main() {
	// result: "a"
	// arr := []string{"d","b","c","b","c","a"}
	// k := int(2)

	// result: "aaa"
	// arr := []string{"aaa","aa","a"}
	// k := int(1)

	// result: ""
	arr := []string{"a","b","a"}
	k := int(3)

	// result: 
	// arr := []string{}
	// k := int(0)

	result := kthDistinct(arr, k)
	fmt.Printf("result = %v\n", result)
}

