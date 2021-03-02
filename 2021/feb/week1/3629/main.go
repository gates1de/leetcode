package main
import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	pathQueue := []string{}
	for _, str := range strings.Split(path, "/") {
		if str == "" || str == "." {
			continue
		} else if str == ".." {
			if len(pathQueue) > 0 {
				pathQueue = pathQueue[:len(pathQueue) - 1]
			}
			continue
		}
		pathQueue = append(pathQueue, str)
		fmt.Printf("pathQueue = %v\n", pathQueue)
	}
	result := "/"
	slash := ""
	for _, str := range pathQueue {
		result += slash + str
		slash = "/"
	}
	return result
}

func main() {
	// path := "/home/"
	// path := "/../"
    // path := "/home//foo/"
	path := "/a/./b/../../c/"

	result := simplifyPath(path)
	fmt.Printf("result = %v\n", result)
}

