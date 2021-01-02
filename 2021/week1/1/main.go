package main
import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	arr := []int{15, 85}
	pieces := [][]int{{85}, {15}}

	result := canFormArray(arr, pieces)
	fmt.Printf("result = %v\n", result)
}

func canFormArray(arr []int, pieces [][]int) bool {
	sortedArr := arr
	sort.Ints(sortedArr)

	fPieces := flatten(pieces, []int{})
	sortedFP := fPieces
	sort.Ints(sortedFP)

	if !reflect.DeepEqual(sortedArr, sortedFP) {
		return false
	}
	if len(pieces) == 1 {
		return reflect.DeepEqual(arr, fPieces)
	}

	// TODO: main process

	return false
}

func flatten(target [][]int, result []int) []int {
    for _, x := range target {
		result = append(result, x...)
    }
    return result
}

