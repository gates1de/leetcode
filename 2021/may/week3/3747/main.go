package main
import (
	"fmt"
	"regexp"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	m := map[string][]string{}
	contentReg, _ := regexp.Compile("\\((.*)\\)")

	for _, path := range paths {
		// fmt.Printf("path = %v\n", path)
		pathAndFiles := strings.Split(path, " ")
		// fmt.Printf("pathAndFiles = %v\n", pathAndFiles)
		if len(pathAndFiles) <= 1 {
			break
		}
		path := pathAndFiles[0]
		for _, component := range pathAndFiles {
			match := contentReg.FindString(component)
			// fmt.Printf("%v: match = %v\n", component, match)
			if match != "" {
				filename := strings.ReplaceAll(component, match, "")
				m[match] = append(m[match], path + "/" + filename)
			}
		}
	}
	// fmt.Printf("m = %v\n", m)

	result := [][]string{}
	for _, files := range m {
		if len(files) <= 1 {
			continue
		}
		result = append(result, files)
	}
	return result
}

func main() {
	// result: [["root/a/2.txt","root/c/d/4.txt","root/4.txt"],["root/a/1.txt","root/c/3.txt"]]
	// paths := []string{"root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)","root 4.txt(efgh)"}

	// result: ["root/a/2.txt","root/c/d/4.txt"],["root/a/1.txt","root/c/3.txt"}
	paths := []string{"root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)"}

	// result: 
	// paths := []string{}

	result := findDuplicate(paths)
	fmt.Printf("result = %v\n", result)
}

