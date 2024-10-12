package main
import (
	"fmt"
	"strings"
)

func minLength(s string) int {
	shouldContinue := false

	for true {
		shouldContinue = false

		if strings.Contains(s, "AB") {
			s = strings.Replace(s, "AB", "", -1)
			shouldContinue = true
		}

		if strings.Contains(s, "CD") {
			s = strings.Replace(s, "CD", "", -1)
			shouldContinue = true
		}

		if !shouldContinue {
			break
		}
	}

	return len(s)
}

func main() {
	// result: 2
	// s := "ABFCACDB"

	// result: 5
	s := "ACBBD"

	// result: 
	// s := ""

	result := minLength(s)
	fmt.Printf("result = %v\n", result)
}

