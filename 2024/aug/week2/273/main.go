package main
import (
	"fmt"
)

type NumberWord struct {
	Val int
	Word string
}

var numberWords = []NumberWord{
	NumberWord{1000000000, "Billion"}, NumberWord{1000000, "Million"},
    NumberWord{1000, "Thousand"}, NumberWord{100, "Hundred"},
    NumberWord{90, "Ninety"}, NumberWord{80, "Eighty"},
    NumberWord{70, "Seventy"}, NumberWord{60, "Sixty"},
    NumberWord{50, "Fifty"}, NumberWord{40, "Forty"},
    NumberWord{30, "Thirty"}, NumberWord{20, "Twenty"},
    NumberWord{19, "Nineteen"}, NumberWord{18, "Eighteen"},
    NumberWord{17, "Seventeen"}, NumberWord{16, "Sixteen"},
    NumberWord{15, "Fifteen"}, NumberWord{14, "Fourteen"},
    NumberWord{13, "Thirteen"}, NumberWord{12, "Twelve"},
    NumberWord{11, "Eleven"}, NumberWord{10, "Ten"},
    NumberWord{9, "Nine"}, NumberWord{8, "Eight"},
    NumberWord{7, "Seven"}, NumberWord{6, "Six"},
    NumberWord{5, "Five"}, NumberWord{4, "Four"},
    NumberWord{3, "Three"}, NumberWord{2, "Two"},
    NumberWord{1, "One"},
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	for _, numberWord := range numberWords {
		if num >= numberWord.Val {
			prefix := ""
			if num >= 100 {
				prefix = numberToWords(num / numberWord.Val) + " "
			}

			unit := numberWord.Word
			suffix := ""
			if num % numberWord.Val != 0 {
				suffix = " " + numberToWords(num % numberWord.Val)
			}

			return prefix + unit + suffix
		}
	}

	return ""
}

func main() {
	// result: "One Hundred Twenty Three"
	// num := int(123)

	// result: "Twelve Thousand Three Hundred Forty Five"
	// num := int(12345)

	// result: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
	num := int(1234567)

	// result: ""
	// num := int(0)

	result := numberToWords(num)
	fmt.Printf("result = %v\n", result)
}

