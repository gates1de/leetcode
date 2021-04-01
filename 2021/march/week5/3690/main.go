package main
import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
    sort.Slice(envelopes, func (a, b int) bool {
        if envelopes[a][0] < envelopes[b][0] {
            return true
        } else if envelopes[a][0] > envelopes[b][0] {
            return false
        } else {
            return envelopes[a][1] > envelopes[b][1]
        }
    })

    dp := make([]int, len(envelopes))
    result := 0

    for i := 0; i < len(envelopes); i++  {
        pos := sort.SearchInts(dp[:result], envelopes[i][1])
		// fmt.Printf("pos = %v, result = %v, envelopes[%v][1] = %v\n", pos, result, i, envelopes[i][1])
        dp[pos] = envelopes[i][1]

        if pos == result {
            result++
        }
    }

    return result
}

// Wrong Answer
func ngSolution(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}

	// sort.Slice(envelopes, func (a, b int) bool {
	// 	if envelopes[a][0] > envelopes[b][0] {
	// 		return true
	// 	} else if envelopes[a][0] == envelopes[b][0] {
	// 		return envelopes[a][1] > envelopes[b][1]
	// 	}
	// 	return false
	// })
	sort.Slice(envelopes, func (a, b int) bool {
		return envelopes[a][0] + envelopes[a][1] > envelopes[b][0] + envelopes[b][1]
	})

	maxResult := int(0)

	for i, head := range envelopes {
		current := head
		result := int(1)
		for j, envelop := range envelopes {
			if i == j {
				continue
			}
			fmt.Printf("current = %v, envelop = %v\n", current, envelop)
			if current[0] > envelop[0] && current[1] > envelop[1] {
				current = envelop
				result++
			}
		}
		fmt.Printf("result = %v\n", result)
		if maxResult < result {
			maxResult = result
		}
	}

	return maxResult
}

func main() {
	// envelopes := [][]int{
	// 	{5, 4},
	// 	{6, 4},
	// 	{6, 7},
	// 	{2, 3},
	// }

	// envelopes := [][]int{
	// 	{1, 1},
	// 	{1, 1},
	// 	{1, 1},
	// }

	// envelopes := [][]int{
	// 	{46, 89},
	// 	{50, 53},
	// 	{52, 68},
	// 	{72, 45},
	// 	{77, 81},
	// }

	// envelopes := [][]int{
	// 	{2,100},{3,200},{4,300},{5,500},{5,400},{5,250},{6,370},{6,360},{7,380},
	// }

	// envelopes := [][]int{
	// 	{7,10},{14,7},{3,16},{6,3},{7,9},{5,2},{10,1},{6,11},{18,20},
	// }

	// envelopes := [][]int{
	// 	{7,8},{12,16},{12,5},{1,8},{4,19},{3,15},{4,10},{9,16},
	// }

	envelopes := [][]int{
		{1,19},{2,13},{10,10},{20,4},{11,5},{1,15},{5,17},{10,1},{15,9},{10,6},
	}

	result := maxEnvelopes(envelopes)
	fmt.Printf("result = %v\n", result)
}

