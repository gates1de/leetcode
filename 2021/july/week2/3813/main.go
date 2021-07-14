package main
import (
	"fmt"
)

// More fast solution: https://leetcode.com/problems/custom-sort-string/discuss/309820/Simple-Go-Solution-all-beats-100.

func customSortString(order string, str string) string {
	orderMap := map[byte]int{}
	for i, r := range order {
		orderMap[byte(r)] = i
	}
	// fmt.Printf("orderMap = %v\n", orderMap)
	strBytes := make([]string, len(str))
	for i, r := range str {
		strBytes[i] = string(r)
	}
	// fmt.Printf("strBytes = %v\n", strBytes)
	result := mergeSort([]byte(str), orderMap)
	return string(result)
}

func mergeSort(slice []byte, orderMap map[byte]int) []byte {
    if len(slice) <= 1 {
        return slice
    }
    mid := len(slice) / 2
    left := mergeSort(slice[:mid], orderMap)
    right := mergeSort(slice[mid:], orderMap)
    return merge(left, right, orderMap)
}

func merge(left []byte, right []byte, orderMap map[byte]int) []byte {
    size := len(left) + len(right)
	k := int(0)
	j := int(0)
    tmp := make([]byte, size)

    for i := 0; i < size; i++ {
        if len(left) <= j {
            tmp[i] = right[k]
            k++
        } else if len(right) <= k {
            tmp[i] = left[j]
            j++
        } else if orderMap[right[k]] < orderMap[left[j]] {
            tmp[i] = right[k]
            k++
        } else {
            tmp[i] = left[j]
            j++
        }
    }
    return tmp
}

func main() {
	// result: "cbad" or "dcba" or "cdba" or "cbda"
	order := "cba"
	str := "abcd"

	// result: 
	// order := ""
	// str := ""

	result := customSortString(order, str)
	fmt.Printf("result = %v\n", result)
}

