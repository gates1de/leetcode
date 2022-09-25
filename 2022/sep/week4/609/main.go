package main
import (
	"fmt"
	"regexp"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	contentRegExp := regexp.MustCompile(`\(.*\)`)
	m := map[string][]string{}
	for _, path := range paths {
		components := strings.Split(path, " ")
		if len(components) == 0 {
			continue
		}
		dirPath := components[0]

		for _, component := range components {
			s := contentRegExp.FindStringSubmatch(component)
			if len(s) == 0 {
				continue
			}

			content := s[0]
			if m[content] == nil {
				m[content] = []string{}
			}
			file := strings.Replace(component, content, "", -1)
			m[content] = append(m[content], dirPath + "/" + file)
		}
	}

	result := [][]string{}
	for _, val := range m {
		if len(val) <= 1 {
			continue
		}
		result = append(result, val)
	}
	return result
}

func main() {
	// result: [["root/a/2.txt","root/c/d/4.txt","root/4.txt"],["root/a/1.txt","root/c/3.txt"]]
	// paths := []string{"root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)","root 4.txt(efgh)"}

	// result: [["root/a/2.txt","root/c/d/4.txt"],["root/a/1.txt","root/c/3.txt"]]
	// paths := []string{"root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)"}

	// result: []
	paths := []string{"root/a 1.txt(abcd) 2.txt(efsfgh)","root/c 3.txt(abdfcd)","root/c/d 4.txt(efggdfh)"}

	// result: 
	// paths := []string{}

	result := findDuplicate(paths)
	fmt.Printf("result = %v\n", result)
}

