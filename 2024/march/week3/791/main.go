package main
import (
	"fmt"
)

func customSortString(order string, str string) string {
    orderMap := map[byte]int{}
    for i, r := range order {
        orderMap[byte(r)] = i
    }
    strBytes := make([]string, len(str))
    for i, r := range str {
        strBytes[i] = string(r)
    }
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
	// result: "cbad"
	// order := "cba"
	// str := "abcd"

	// result: "bcad"
	order := "bcafg"
	str := "abcd"

	// result: ""
	// order := ""
	// str := ""

	result := customSortString(order, str)
	fmt.Printf("result = %v\n", result)
}

