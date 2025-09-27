package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")
	n1 := len(s1)
	n2 := len(s2)
	v1 := make([]int, n1)
	v2 := make([]int, n2)

	for i, char := range s1 {
		v1[i], _ = strconv.Atoi(char)
	}
	for i, char := range s2 {
		v2[i], _ = strconv.Atoi(char)
	}

	if n1 > n2 {
		for i := 0; i < n1 - n2; i++ {
			v2 = append(v2, 0)
		}
	} else if n2 > n1 {
		for i := 0; i < n2 - n1; i++ {
			v1 = append(v1, 0)
		}
	}

	for i := 0; i < len(v1); i++ {
		if v1[i] > v2[i] {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		}
	}

	return 0
}

func main() {
		// result: 0
	// version1 := "1.01"
	// version2 := "1.001"

	// result: 0
	// version1 := "1.0"
	// version2 := "1.0.0"

	// result: -1
	version1 := "0.1"
	version2 := "1.1"

	// result: 
	// version1 := ""
	// version2 := ""

	result := compareVersion(version1, version2)
	fmt.Printf("result = %v\n", result)
}

