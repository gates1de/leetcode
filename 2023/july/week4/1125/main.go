package main
import (
	"fmt"
	"math/bits"
)

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	n := len(people)
	m := len(req_skills)
	skillId := make(map[string]int)
	for i, skill := range req_skills {
		skillId[skill] = i
	}

	skillsMaskOfPerson := make([]int, n)
	for i, skills := range people {
		for _, skill := range skills {
			skillsMaskOfPerson[i] |= 1 << skillId[skill]
		}
	}

	dp := make([]int64, 1 << m)
	for i, _ := range dp {
		dp[i] = int64(1 << n) - 1
	}
	dp[0] = 0

	for skillsMask := 1; skillsMask < (1 << m); skillsMask++ {
		for i := 0; i < n; i++ {
			smallerSkillsMask := skillsMask & ^skillsMaskOfPerson[i]
			if smallerSkillsMask != skillsMask {
				peopleMask := dp[smallerSkillsMask] | (int64(1) << i)
				if bits.OnesCount64(uint64(peopleMask)) < bits.OnesCount64(uint64(dp[skillsMask])) {
					dp[skillsMask] = peopleMask
                }
            }
        }
    }

	answerMask := dp[(1 << m) - 1]
	result := make([]int, bits.OnesCount64(uint64(answerMask)))
	ptr := int(0)
	for i := 0; i < n; i++ {
		if ((answerMask >> i) & 1) == 1 {
			result[ptr] = i
			ptr++
		}
	}

	return result
}

func main() {
	// result: [0,2]
	// req_skills := []string{"java","nodejs","reactjs"}
	// people := [][]string{{"java"},{"nodejs"},{"nodejs","reactjs"}}

	// result: [1,2]
	req_skills := []string{"algorithms","math","java","reactjs","csharp","aws"}
	people := [][]string{{"algorithms","math","java"},{"algorithms","math","reactjs"},{"java","csharp","aws"},{"reactjs","csharp"},{"csharp","math"},{"aws","java"}}

	// result: 
	// req_skills := []string{}
	// people := [][]string{}

	result := smallestSufficientTeam(req_skills, people)
	fmt.Printf("result = %v\n", result)
}

