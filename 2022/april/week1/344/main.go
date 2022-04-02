package main
import (
	"fmt"
)

func reverseString(s []byte)  {
	for i := 0; i < len(s) / 2; i++ {
		s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
	}
}

func printBytes(s []byte) {
	for _, b := range s {
		fmt.Printf("%v", string(b))
	}
}

func main() {
	// result: ["o","l","l","e","h"]
	// s := []byte{'h', 'e', 'l', 'l', 'o'}

	// result: ["h","a","n","n","a","H"]
	// s := []byte{'H', 'a', 'n', 'n', 'a', 'h'}

	// result: ["a"]
	s := []byte{'a'}

	// result: 
	// s := []byte{}

	reverseString(s)
	printBytes(s)
}

