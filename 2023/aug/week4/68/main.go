package main
import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)
	i := int(0)

	for i < len(words) {
		currentLine := getWords(i, words, maxWidth)
		i += len(currentLine)
		result = append(result, createLine(currentLine, i, words, maxWidth))
	}

	return result
}

func getWords(i int, words []string, maxWidth int) []string {
	result := make([]string, 0)
	currentLength := int(0)

	for i < len(words) && currentLength + len(words[i]) <= maxWidth {
		result = append(result, words[i])
		currentLength += len(words[i]) + 1
		i++
	}

	return result
}

func createLine(line []string, i int, words []string, maxWidth int) string {
	baseLength := int(-1)
	for _, word := range line {
		baseLength += len(word) + 1
	}

	extraSpaces := maxWidth - baseLength

	if len(line) == 1 || i == len(words) {
		return strings.Join(line, " ") + strings.Repeat(" ", extraSpaces)
	}

	wordCount := len(line) - 1
	spacesPerWord := extraSpaces / wordCount
	needsExtraSpace := extraSpaces % wordCount

	for j := 0; j < needsExtraSpace; j++ {
		line[j] += " "
	}

	for j := 0; j < wordCount; j++ {
		line[j] += strings.Repeat(" ", spacesPerWord)
	}

	return strings.Join(line, " ")
}

func main() {
	// result: 
	// [
	//    "This    is    an",
	//    "example  of text",
	//    "justification.  "
	// ]
	// words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	// maxWidth := int(16)

	// result: 
	// [
	//   "What   must   be",
	//   "acknowledgment  ",
	//   "shall be        "
	// ]
	// words := []string{"What","must","be","acknowledgment","shall","be"}
	// maxWidth := int(16)

	// result: 
	// [
	//   "Science  is  what we",
	//   "understand      well",
	//   "enough to explain to",
	//   "a  computer.  Art is",
	//   "everything  else  we",
	//   "do                  "
	// ]
	 words := []string{"Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"}
	 maxWidth := int(20)

	// result: 
	// words := []string{}
	// maxWidth := int(0)

	result := fullJustify(words, maxWidth)
	fmt.Printf("result = %v\n", result)
}

