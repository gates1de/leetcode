package main
import (
	"fmt"
)

type CombinationIterator struct {
	Combinations []string
	Index int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	combinations := []string{}
	helper(0, "", characters, combinationLength, &combinations)
	// fmt.Printf("combinations = %v\n", combinations)
	return CombinationIterator{combinations, 0}
}

func (this *CombinationIterator) Next() string {
	if !this.HasNext() {
		return ""
	}
	result := this.Combinations[this.Index]
	this.Index++
	return result
}

func (this *CombinationIterator) HasNext() bool {
	return this.Index < len(this.Combinations)
}

func helper(index int, currentChars string, chars string, length int, combinations *[]string) {
	if index == length {
		*combinations = append(*combinations, currentChars)
		return
	}

	for i := 0; i < len(chars); i++ {
		helper(index + 1, currentChars + string(chars[i]), chars[i + 1:], length, combinations)
	}
}

func main() {
	// result: [null, "ab", true, "ac", true, "bc", false]
	// characters := "abc"
	// combinationLength := int(2)

	// result: 
	// characters := ""
	// combinationLength := int(0)

	obj := Constructor(characters, combinationLength)
	fmt.Printf("obj.Next() = %v\n", obj.Next())
	fmt.Printf("obj.HasNext() = %v\n", obj.HasNext())
	fmt.Printf("obj.Next() = %v\n", obj.Next())
	fmt.Printf("obj.HasNext() = %v\n", obj.HasNext())
	fmt.Printf("obj.Next() = %v\n", obj.Next())
	fmt.Printf("obj.HasNext() = %v\n", obj.HasNext())
}

