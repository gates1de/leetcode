package main
import (
	"fmt"
)

const L = byte(76)
const R = byte(82)

func pushDominoes(dominoes string) string {
	dominoesInByteSlice := []byte(dominoes)
	lastIsRight := false
	lastIndex := int(0)
	for i := 0; i < len(dominoesInByteSlice); i++ {
		switch dominoesInByteSlice[i] {
		case 'L':
			if lastIsRight {
				count := (i - lastIndex + 1) / 2 - 1
				for m := lastIndex + 1; m < lastIndex + 1 + count; m++ {
					dominoesInByteSlice[m] = 'R'
				}
				for m := i - 1; m > i - 1 - count; m-- {
					dominoesInByteSlice[m] = 'L'
				}
			} else {
				for m := lastIndex; m < i; m++ {
					dominoesInByteSlice[m] = 'L'
				}
			}
			lastIsRight = false
			lastIndex = i + 1
		case 'R':
			if lastIsRight {
				for m := lastIndex; m < i; m++ {
					dominoesInByteSlice[m] = 'R'
				}
			}
			lastIsRight = true
			lastIndex = i
		}
	}

	if lastIsRight {
		for m := lastIndex; m < len(dominoesInByteSlice); m++ {
			dominoesInByteSlice[m] = 'R'
		}
	}

	return string(dominoesInByteSlice)
}

// Slow
func mySolution(dominoes string) string {
	if len(dominoes) == 1 {
		return dominoes
	}

	result := ""
	count := int(0)
	memo := byte(0)
	for _, r := range dominoes {
		count++
		if byte(r) == L {
			if memo == R {
				// ex) "R...L" -> "RR.LL"
				for i := 1; i <= count; i++ {
					if count % 2 == 0 && i == count / 2 {
						result += "."
					} else if i <= count / 2 {
						result += "R"
					} else {
						result += "L"
					}
				}
				count = 0
				memo = byte(0)
			} else {
				// ex) "...L" -> "LLLL"
				for i := 1; i <= count; i++ {
					result += "L"
				}
				count = 0
				memo = byte(0)
			}
		} else if byte(r) == R {
			if memo == R {
				// ex) "R.R" -> "RRR"
				for i := 1; i < count; i++ {
					result += "R"
				}
			} else {
				// ex) "...R" -> "...R"
				for i := 1; i < count; i++ {
					result += "."
				}
			}
			result += "R"
			count = 0
			memo = R
		}
	}

	if count > 0 {
		if memo == R {
			for i := 1; i <= count; i++ {
				result += "R"
			}
		} else {
			for i := 1; i <= count; i++ {
				result += "."
			}
		}
	}
	return result
}

func main() {
	// result: "RR.L"
	// dominoes := "RR.L"

	// result: "LL.RR.LLRRLL.."
	dominoes := ".L.R...LR..L.."

	// result:
	// dominoes := ""

	result := pushDominoes(dominoes)
	fmt.Printf("result = %v\n", result)
}

