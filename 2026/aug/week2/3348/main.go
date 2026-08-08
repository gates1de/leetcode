package main

import (
	"fmt"
)

const impossible = int(1e9)

type factorCount struct {
	two   int
	three int
	five  int
	seven int
}

var digitFactor = [10]factorCount{
	{},
	{},
	{two: 1},
	{three: 1},
	{two: 2},
	{five: 1},
	{two: 1, three: 1},
	{seven: 1},
	{two: 3},
	{three: 2},
}

var primeFactors = [4]int64{2, 3, 5, 7}

func remainingFactors(target factorCount, used factorCount) factorCount {
	result := factorCount{
		two:   max(target.two-used.two, 0),
		three: max(target.three-used.three, 0),
		five:  max(target.five-used.five, 0),
		seven: max(target.seven-used.seven, 0),
	}
	return result
}

func consumeFactors(required factorCount, digit int) factorCount {
	factor := digitFactor[digit]
	return remainingFactors(required, factor)
}

func minimumDigits(required factorCount, table [][]int) int {
	return table[required.two][required.three] + required.five + required.seven
}

func buildSmallest(required factorCount, length int, table [][]int) string {
	if minimumDigits(required, table) > length {
		return ""
	}

	resultBytes := make([]byte, length)
	for index := range resultBytes {
		remainingLength := length - index - 1
		for digit := range digitFactor[1:] {
			candidateDigit := digit + 1
			candidateRequired := consumeFactors(required, candidateDigit)
			if minimumDigits(candidateRequired, table) > remainingLength {
				continue
			}
			resultBytes[index] = byte('0' + candidateDigit)
			required = candidateRequired
			break
		}
	}
	return string(resultBytes)
}

func smallestNumber(num string, t int64) string {
	target := factorCount{}
	remaining := t
	for _, prime := range primeFactors {
		for remaining%prime == 0 {
			switch prime {
			case 2:
				target.two++
			case 3:
				target.three++
			case 5:
				target.five++
			case 7:
				target.seven++
			}
			remaining /= prime
		}
	}
	if remaining != 1 {
		return "-1"
	}

	table := make([][]int, target.two+1)
	for two := range table {
		table[two] = make([]int, target.three+1)
		for three := range table[two] {
			table[two][three] = impossible
		}
	}
	table[0][0] = 0
	for two := range table {
		for three := range table[two] {
			if two == 0 && three == 0 {
				continue
			}
			for digit, factor := range digitFactor {
				if digit < 2 || factor.two == 0 && factor.three == 0 {
					continue
				}
				previousTwo := max(two-factor.two, 0)
				previousThree := max(three-factor.three, 0)
				candidate := table[previousTwo][previousThree] + 1
				if candidate < table[two][three] {
					table[two][three] = candidate
				}
			}
		}
	}

	n := len(num)
	prefix := make([]factorCount, n+1)
	zeroPrefix := make([]int, n+1)
	for index, character := range num {
		prefix[index+1] = prefix[index]
		zeroPrefix[index+1] = zeroPrefix[index]
		if character == '0' {
			zeroPrefix[index+1]++
			continue
		}
		factor := digitFactor[character-'0']
		prefix[index+1].two += factor.two
		prefix[index+1].three += factor.three
		prefix[index+1].five += factor.five
		prefix[index+1].seven += factor.seven
	}

	if zeroPrefix[n] == 0 && remainingFactors(target, prefix[n]) == (factorCount{}) {
		return num
	}

	for reverse := range n {
		index := n - 1 - reverse
		if zeroPrefix[index] != 0 || num[index] == '9' {
			continue
		}

		for digit := range digitFactor[1:] {
			candidateDigit := digit + 1
			if candidateDigit <= int(num[index]-'0') {
				continue
			}
			required := consumeFactors(remainingFactors(target, prefix[index]), candidateDigit)
			suffixLength := n - index - 1
			if minimumDigits(required, table) > suffixLength {
				continue
			}

			resultBytes := make([]byte, n)
			copy(resultBytes, num[:index])
			resultBytes[index] = byte('0' + candidateDigit)
			suffix := buildSmallest(required, suffixLength, table)
			copy(resultBytes[index+1:], suffix)
			return string(resultBytes)
		}
	}

	length := max(n+1, minimumDigits(target, table))
	result := buildSmallest(target, length, table)
	if result == "" {
		return "-1"
	}
	return result
}

func main() {
	// result: "1488"
	// num := "1234"
	// t := int64(256)

	// result: "12355"
	// num := "12355"
	// t := int64(50)

	// result: "-1"
	// num := "11111"
	// t := int64(26)

	// result: "255555579"
	num := "12"
	t := int64(1968750)

	// result: ""
	// num := ""
	// t := int64(0)

	result := smallestNumber(num, t)
	fmt.Printf("result = %v\n", result)
}
