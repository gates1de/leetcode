package main
import (
	"fmt"
	"strconv"
	"strings"
)

func hasAllCodes(s string, k int) bool {
    tmp := pow(k)
    for i := tmp; i >= tmp / 2; i-- {
        bin := strconv.FormatInt(int64(i), 2)
        bin = bin[len(bin) - k:]
        if !strings.Contains(s, bin) {
            return false
        }
    }
    return true
}

func pow(k int) int {
    result := int(1)
    for i := 0; i < k; i++ {
        result *= 2
    }
    return result
}

func main() {
	// result: true
	// s := "00110110"
	// k := int(2)

	// result: true
	// s := "0110"
	// k := int(1)

	// result: false
	s := "0110"
	k := int(2)

	// result: 
	// s := ""
	// k := int(0)

	result := hasAllCodes(s, k)
	fmt.Printf("result = %v\n", result)
}

