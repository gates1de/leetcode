package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)
const ast = byte(42)
const zero = byte(48)
const one = byte(49)
const two = byte(50)
const six = byte(54)

func numDecodings(s string) int {
	if s[0] == zero {
		return 0
	}
	if len(s) == 1 {
		if s == "*" {
			return 9
		}
		return 1
	}

	if len(s) >= 3 && s[len(s) - 1] == zero {
		thirdTailChar := s[len(s) - 3]
		secondTailChar := s[len(s) - 2]
		tail := string([]byte{thirdTailChar, secondTailChar})
		if (thirdTailChar != ast && secondTailChar != ast) && (tail < "11" || "22" < tail) {
			return 0
		}
	}

	dp := map[string]int{}
	dp[""] = 1
	if s[0] == ast {
		dp[string(s[0])] = 9
	} else if s[0] == zero {
		dp[""] = 0
	} else {
		dp[string(s[0])] = 1
	}

	for i := 1; i < len(s); i++ {
		preS1 := s[:i - 1]
		preS2 := s[:i]
		currentS := s[:i + 1]

		preChar := s[i - 1]
		char := s[i]

		if char == zero {
			if preChar == ast {
				dp[preS2] = dp[preS1] * 2
				dp[currentS] = dp[preS1] * 2
			} else if preChar == zero {
				return 0
			} else {
				if preS1 == "" && preChar > two {
					dp[preS2] = 0
					dp[currentS] = 0
				} else {
					dp[preS2] = dp[preS1]
					dp[currentS] = dp[preS2]
				}
			}
		} else if char == ast {
			if preChar == ast {
				countS1 := dp[preS1] * 15
				countS2 := dp[preS2] * 9
				dp[currentS] = countS1 + countS2
			} else if preChar == zero {
				dp[currentS] = dp[preS1] * 9
			} else if preChar == one {
				dp[currentS] = dp[preS1] * 9 + dp[preS2] * 9
			} else if preChar == two {
				dp[currentS] = dp[preS1] * 6 + dp[preS2] * 9
			} else {
				dp[currentS] = dp[preS2] * 9
			}
		} else {
			twoSubstring := string([]byte{preChar, char})
			if preChar == ast {
				countS1 := dp[preS1]
				if char <= six {
					countS1 = dp[preS1] * 2
				}
				countS2 := dp[preS2]
				dp[currentS] = countS1 + countS2
			} else if preChar == zero {
				dp[currentS] = dp[preS2]
			} else if "11" <= twoSubstring && twoSubstring <= "26" {
				dp[currentS] = dp[preS2] + dp[preS1]
			} else {
				dp[currentS] = dp[preS2]
			}
		}
		// fmt.Printf("currentS = %v, dp = %v\n", currentS, dp)
	}
	return dp[s] % modulo
}

func main() {
	// result: 9
	// s := "*"

	// result: 18
	// s := "1*"

	// result: 15
	// s := "2*"

	// result: 2
	// s := "11106"

	// result: 17
	// s := "2*3"

	// result: 0
	// s := "2*3112826*004"

	// result: 408
	// s := "2*3112826*014"

	// result: 404
	// s := "*1*1*0"

	// result: 0
	// s := "904"

	// result: 0
	// s := "1040"

	// result: 1
	// s := "10"

	// result: 36
	// s := "*0**0"

	// result: 285
	// s := "1*72*"

	// result: 0
	s := "96603"

	// result: 
	// s := ""

	result := numDecodings(s)
	fmt.Printf("result = %v\n", result)
}

