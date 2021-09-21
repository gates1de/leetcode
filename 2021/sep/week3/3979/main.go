package main
import (
	"fmt"
	"strconv"
)

func addOperators(num string, target int) []string {
    var ret = []string{}
    var buf = make([]byte, len(num)*2) // shared bytes buffer

    search(num, target, &ret, 0, 0, buf, 0)

    return ret
}

// `l` is the length of valid content in `buf`
func search(num string, target int, ret *[]string, sum, diff int, buf []byte, l int) {
    if len(num) == 0 {
        if sum == target {
            *ret = append(*ret, string(buf[0:l]))
        }
        return
    }

    var n = 0
	// try all valid digit combinations
    for i := 0; i < len(num); i++ {
        n = n*10 + int(num[i] - '0')
        var digits = num[0:i+1]

        if l == 0 {
            copy(buf[l:], digits)
            search(num[i+1:], target, ret, sum + n, n, buf, l+i+1)
        } else {
            copy(buf[l+1:], digits) // one copy for the three operators

            buf[l] = '+'
            search(num[i+1:], target, ret, sum + n, n, buf, l+1+i+1) // add `1 operator + (i+1) digits` to the length of buffer

            buf[l] = '-'
            search(num[i+1:], target, ret, sum - n, -n, buf, l+1+i+1)

            buf[l] = '*'
            search(num[i+1:], target, ret, sum - diff + diff * n, diff * n, buf, l+1+i+1)
        }

        if i == 0 && num[i] == '0' { // skip any number starts with 0
            break
        }
    }
}

// Slow & Use more memory
func mySolution(num string, target int) []string {
	result := []string{}

	str := string(rune(num[0]))
	helper(str, 1, num, target, &result)

	return result
}

func helper(currentStr string, index int, num string, target int, result *[]string) {
	if index >= len(num) {
		n := calc(currentStr, target)
		// fmt.Printf("%v = %v\n", currentStr, n)
		if n == target {
			*result = append(*result, currentStr)
		}
		return
	}

	str := string(rune(num[index]))
	helper(currentStr + "+" + str, index + 1, num, target, result)
	helper(currentStr + "-" + str, index + 1, num, target, result)
	helper(currentStr + "*" + str, index + 1, num, target, result)
	helper(currentStr + str, index + 1, num, target, result)
}

func calc(num string, target int) int {
	result := int(0)
	arr := []string{}
	str := ""
	for _, r := range num {
		if r == rune('+') {
			arr = append(arr, str)
			arr = append(arr, "+")
			str = ""
		} else if r == rune('-') {
			arr = append(arr, str)
			arr = append(arr, "-")
			str = ""
		} else if r == rune('*') {
			arr = append(arr, str)
			arr = append(arr, "*")
			str = ""
		} else {
			str += string(r)
			if len(str) >= 2 && rune(str[0]) == rune('0') {
				return target - 1
			}
		}
	}
	arr = append(arr, str)

	for i, str := range arr {
		if str == "*" {
			pre := arr[i - 1]
			next := arr[i + 1]
			preNum, _ := strconv.Atoi(pre)
			nextNum, _ := strconv.Atoi(next)
			arr[i - 1] = ""
			arr[i] = ""
			arr[i + 1] = strconv.Itoa(preNum * nextNum)
		}
	}

	// fmt.Printf("arr = %v\n", arr)
	isPlus := true
	for _, str := range arr {
		if str == "+" {
			isPlus = true
		} else if str == "-" {
			isPlus = false
		} else if str == "" {
			continue
		} else {
			n, _ := strconv.Atoi(str)
			if isPlus {
				result += n
			} else {
				result -= n
			}
		}
	}

	return result
}

func main() {
	// result: ["1*2*3","1+2+3"]
	// num := "123"
	// target:= int(6)

	// result: ["2*3+2","2+3*2"]
	// num := "232"
	// target:= int(8)

	// result: ["1*0+5","10-5"]
	// num := "105"
	// target:= int(5)

	// result: ["0*0","0+0","0-0"]
	// num := "00"
	// target:= int(0)

	// result: []
	num := "3456237490"
	target:= int(9191)

	// result: ["0"]
	// num := "0"
	// target:= int(0)

	result := addOperators(num, target)
	fmt.Printf("result = %v\n", result)
}

