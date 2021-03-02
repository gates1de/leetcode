package main
import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	arr := []int{49, 18, 16}
	pieces := [][]int{{16, 18, 49}}

	result := canFormArray(arr, pieces)
	fmt.Printf("result = %v\n", result)
}

func canFormArray(arr []int, pieces [][]int) bool {
	sortedArr := copySlice(arr)
	sort.Ints(sortedArr)

	fPieces := flatten(pieces, []int{})
	sortedFP := copySlice(fPieces)
	sort.Ints(sortedFP)

	if !reflect.DeepEqual(sortedArr, sortedFP) {
		return false
	}
	if len(pieces) == 1 {
		return reflect.DeepEqual(arr, fPieces)
	}

	// TODO: main process
	for _, arrV := range arr {
		for _, piece := range pieces {
			pieceLen := len(piece)
			pieceIndex := indexOf(arrV, piece)
			if pieceIndex != 0 {
				continue
			}
			slicedArr := arr[:pieceLen]
			if !reflect.DeepEqual(slicedArr, piece) {
				return false
			}
			arr = arr[pieceLen:]
		}
	}

	if len(arr) == 0 {
		return true
	}

	return false
}

func copySlice(target []int) []int {
	result := make([]int, len(target))
	copy(result, target)
	return result
}

func flatten(target [][]int, result []int) []int {
    for _, x := range target {
		result = append(result, x...)
    }
    return result
}

func indexOf(target int, arr []int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

