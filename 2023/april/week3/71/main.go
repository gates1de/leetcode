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
	// result: "/home"
	// path := "/home/"

	// result: "/"
	// path := "/../"

	// result: "/home/foo"
	path := "/home//foo/"

	// result: 
	// path := ""

	result := simplifyPath(path)
	fmt.Printf("result = %v\n", result)
}

