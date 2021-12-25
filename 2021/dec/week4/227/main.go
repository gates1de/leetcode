package main
import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"strconv"
)

func calculate(s string) (res int) {
    sidx, pidx := 0, 0
    ll := len(s)
    mode := 0
    pnum := 0
    op := byte('i')
    for pidx < ll {
        if s[pidx] == ' ' {
            pidx++
            continue
        }
        if mode == 1 {
            op = s[pidx]
            pidx++
            mode = 0
            continue
        }
        sidx = pidx
        for pidx < ll && s[pidx] >= '0' && s[pidx] <= '9' { pidx++ }
        cnum, _ := strconv.Atoi(string(s[sidx:pidx]))
        switch op {
            case '*':
            {
                pnum *= cnum
            }
            case '/':
            {
                pnum /= cnum
            }
            case '+':
            {
                res += pnum
                pnum = cnum
            }
            case '-':
            {
                res += pnum
                pnum = -cnum
            }
            case 'i':
            {
                pnum = cnum
            }
        }
        mode = 1
    }
    res += pnum
    return
}

// Slow & Use more memory
func mySolution(s string) int {
	s = strings.Replace(s, " ", "", -1)

	mulOrDivRe := regexp.MustCompile(`[0-9]+\*[0-9]+|[0-9]+/[0-9]+`)
	mulOrDivs := mulOrDivRe.FindAllString(s, 1)
	// fmt.Println(muls)
	for len(mulOrDivs) > 0 {
		// fmt.Printf("muls[0] = %v\n", muls[0])
		target := mulOrDivs[0]
		mulRe := regexp.MustCompile(`\*`)
		divRe := regexp.MustCompile(`/`)
		numStr := ""
		if mulRe.MatchString(target) {
			nums := mulRe.Split(target, -1)
			if len(nums) != 2 {
				return math.MinInt32
			}
			leftNum, _ := strconv.Atoi(nums[0])
			rightNum, _ := strconv.Atoi(nums[1])
			numStr = fmt.Sprintf("%v", leftNum * rightNum)
		} else {
			nums := divRe.Split(target, -1)
			if len(nums) != 2 {
				return math.MinInt32
			}
			leftNum, _ := strconv.Atoi(nums[0])
			rightNum, _ := strconv.Atoi(nums[1])
			numStr = fmt.Sprintf("%v", leftNum / rightNum)
		}

		s = strings.Replace(s, target, numStr, 1)
		mulOrDivs = mulOrDivRe.FindAllString(s, 1)
	}
	// fmt.Printf("calc mul or div s = %v\n", s)

	addOrSubRe := regexp.MustCompile(`\+|-`)
	addOrSubs := addOrSubRe.FindAllString(s, -1)
	nums := addOrSubRe.Split(s, -1)
	// fmt.Printf("addOrSubs = %v, nums = %v\n", addOrSubs, nums)
	for i, numStr := range nums[:len(nums) - 1] {
		if len(addOrSubs) < i {
			return math.MinInt32
		}

		nextNumStr := nums[i + 1]
		num, err1 := strconv.Atoi(numStr)
		nextNum, err2 := strconv.Atoi(nextNumStr)
		if err1 != nil || err2 != nil {
			return math.MinInt32
		}

		addOrSub := addOrSubs[i]
		calcNum := int(0)
		if addOrSub == "+" {
			calcNum = num + nextNum
		} else {
			calcNum = num - nextNum
		}
		calcNumStr := fmt.Sprintf("%v", calcNum)
		nums[i + 1] = calcNumStr
	}

	resultStr := nums[len(nums) - 1]
	result, err := strconv.Atoi(resultStr)
	if err != nil {
		return math.MinInt32
	}
	return result
}

func main() {
	// result: 7
	// s := "3+2*2"

	// result: 1
	// s := " 3/2 "

	// result: 5
	// s := " 3+5 / 2 "

	// result: 10
	// s := "50 +4*2 * 60 /6-120"

	// result: 1
	// s := "1-1+1"

	// result: 1
	// s := "1"

	// result: 8
	s := "14/3*2"

	// result: 
	// s := ""

	result := calculate(s)
	fmt.Printf("result = %v\n", result)
}

