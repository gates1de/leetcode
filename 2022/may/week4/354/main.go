package main
import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
    sort.Slice(envelopes, func (a, b int) bool {
		e1 := envelopes[a]
		e2 := envelopes[b]
        if e1[0] < e2[0] {
            return true
        } else if e1[0] > e2[0] {
            return false
        } else {
            return e1[1] > e2[1]
        }
    })

    dp := make([]int, len(envelopes))
    result := int(0)

    for i := 0; i < len(envelopes); i++  {
        pos := sort.SearchInts(dp[:result], envelopes[i][1])
        dp[pos] = envelopes[i][1]

        if pos == result {
            result++
        }
    }

    return result
}

func main() {
	// result: 3
	// envelopes := [][]int{{5,4},{6,4},{6,7},{2,3}}

	// result: 1
	// envelopes := [][]int{{1,1},{1,1},{1,1}}

	// result: 3
	// envelopes := [][]int{{46,89},{50,53},{52,68},{72,45},{77,81}}

	// result: 5
	// envelopes := [][]int{{2,100},{3,200},{4,300},{5,500},{5,400},{5,250},{6,370},{6,360},{7,380}}

	// result: 3
	// envelopes := [][]int{{7,8},{12,16},{12,5},{1,8},{4,19},{3,15},{4,10},{9,16}}

	// result: 4
	// envelopes := [][]int{{4,5},{4,6},{6,7},{2,3},{1,1}}

	// result: 3
	// envelopes := [][]int{{1,19},{2,13},{10,10},{20,4},{11,5},{1,15},{5,17},{10,1},{15,9},{10,6}}

	// result: 3
	// envelopes := [][]int{{1,15},{7,18},{7,6},{7,100},{2,200},{17,30},{17,45},{3,5},{7,8},{3,6},{3,10},{7,20},{17,3},{17,45}}

	// result: 5
	envelopes := [][]int{{8,18},{4,14},{16,1},{9,11},{14,15},{12,19},{2,15},{4,4},{18,3},{20,8},{19,18},{18,2},{1,10},{12,1},{10,16},{1,1},{3,19}}

	// result: 
	// envelopes := [][]int{}

	result := maxEnvelopes(envelopes)
	fmt.Printf("result = %v\n", result)
}

