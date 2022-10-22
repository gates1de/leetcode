package main
import (
	"fmt"
	"sort"
)

type Frequency struct {
	Word string
	Count int
}

func topKFrequent(words []string, k int) []string {
	counts := map[string]int{}
	for _, word := range words {
		counts[word]++
	}

	frequencies := make([]Frequency, 0, len(counts))
	for word, count := range counts {
		frequencies = append(frequencies, Frequency{Word: word, Count: count})
	}

	sort.Slice(frequencies, func(a, b int) bool {
		freqA := frequencies[a]
		freqB := frequencies[b]
		if freqA.Count > freqB.Count {
			return true
		} else if freqA.Count < freqB.Count {
			return false
		}

		short := freqA.Word
		long := freqB.Word
		isShortA := true
		if len(short) > len(long) {
			short, long = long, short
			isShortA = false
		}

		for i, char1 := range short {
			char2 := rune(long[i])
			if char1 < char2 {
				return isShortA
			} else if char1 > char2 {
				return !isShortA
			}
		}
		return len(freqA.Word) < len(freqB.Word)
	})

	result := make([]string, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, frequencies[i].Word)
	}
	return result
}

func main() {
	// result: ["i","love"]
	// words := []string{"i","love","leetcode","i","love","coding"}
	// k := int(2)

	// result: ["the","is","sunny","day"]
	// words := []string{"the","day","is","sunny","the","the","the","sunny","is","is"}
	// k := int(4)

	// result: ["dot"]
	// words := []string{"hot","dot","dog","dot","hot"}
	// k := int(1)

	// result: ["bad","bag"]
	// words := []string{"bat","bag","bat","bag","bad","bad"}
	// k := int(2)

	// result: ["fvvdtpnzx"]
	words := []string{"plpaboutit","jnoqzdute","sfvkdqf","mjc","nkpllqzjzp","foqqenbey","ssnanizsav","nkpllqzjzp","sfvkdqf","isnjmy","pnqsz","hhqpvvt","fvvdtpnzx","jkqonvenhx","cyxwlef","hhqpvvt","fvvdtpnzx","plpaboutit","sfvkdqf","mjc","fvvdtpnzx","bwumsj","foqqenbey","isnjmy","nkpllqzjzp","hhqpvvt","foqqenbey","fvvdtpnzx","bwumsj","hhqpvvt","fvvdtpnzx","jkqonvenhx","jnoqzdute","foqqenbey","jnoqzdute","foqqenbey","hhqpvvt","ssnanizsav","mjc","foqqenbey","bwumsj","ssnanizsav","fvvdtpnzx","nkpllqzjzp","jkqonvenhx","hhqpvvt","mjc","isnjmy","bwumsj","pnqsz","hhqpvvt","nkpllqzjzp","jnoqzdute","pnqsz","nkpllqzjzp","jnoqzdute","foqqenbey","nkpllqzjzp","hhqpvvt","fvvdtpnzx","plpaboutit","jnoqzdute","sfvkdqf","fvvdtpnzx","jkqonvenhx","jnoqzdute","nkpllqzjzp","jnoqzdute","fvvdtpnzx","jkqonvenhx","hhqpvvt","isnjmy","jkqonvenhx","ssnanizsav","jnoqzdute","jkqonvenhx","fvvdtpnzx","hhqpvvt","bwumsj","nkpllqzjzp","bwumsj","jkqonvenhx","jnoqzdute","pnqsz","foqqenbey","sfvkdqf","sfvkdqf"}
	k := int(1)

	// result: 
	// words := []string{}
	// k := int(0)

	result := topKFrequent(words, k)
	fmt.Printf("result = %v\n", result)
}

