package main
import (
	"fmt"
)

func backspaceCompare(s string, t string) bool {
    s = convert(s)
    t = convert(t)
    return s == t
}

func convert(s string) string {
    for i := 0; i < len(s); i++ {
        if s[i] != byte('#') {
            continue
        }

        if i - 1 < 0 {
            s = s[i + 1:]
            i--
            continue
        }

        s = s[:i - 1] + s[i + 1:]
        i -= 2
    }
    return s
}

func main() {
	// result: true
	// s := "ab#c"
	// t := "ad#c"

	// result: true
	// s := "ab##"
	// t := "c#d#"

	// result: false
	s := "a#c"
	t := "b"

	// result: 
	// s := ""
	// t := ""

	result := backspaceCompare(s, t)
	fmt.Printf("result = %v\n", result)
}

