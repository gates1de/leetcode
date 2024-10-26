package main
import (
	"fmt"
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Slice(folder, func(a, b int) bool {
		return folder[a] < folder[b]
	})

	result := make([]string, 0)
	result = append(result, folder[0])

	for i := 1; i < len(folder); i++ {
		last := result[len(result) - 1]
		lastChars := []byte(last)
		lastChars = append(lastChars, '/')

		if !strings.HasPrefix(folder[i], string(lastChars)) {
			result = append(result, folder[i])
		}
	}

	return result
}

func main() {
	// result: ["/a","/c/d","/c/f"]
	// folder := []string{"/a","/a/b","/c/d","/c/d/e","/c/f"}

	// result: ["/a"]
	// folder := []string{"/a","/a/b/c","/a/b/d"}

	// result: ["/a/b/c","/a/b/ca","/a/b/d"]
	folder := []string{"/a/b/c","/a/b/ca","/a/b/d"}

	// result: []
	// folder := []string{}

	result := removeSubfolders(folder)
	fmt.Printf("result = %v\n", result)
}

