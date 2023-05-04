package main
import (
	"fmt"
)

func predictPartyVictory(senate string) string {
    rCount := int(0)
    dCount := int(0)
    rFloatingBan := int(0)
    dFloatingBan := int(0)

    queue := make([]byte, len(senate))
    for i := 0; i < len(senate); i++ {
        queue[i] = senate[i]
        if senate[i] == 'R' {
            rCount++
        } else {
            dCount++
        }
    }

    for rCount > 0 && dCount > 0 {
        current := queue[0]
        queue = queue[1:]

        if current == 'D' {
            if dFloatingBan > 0 {
                dFloatingBan--
                dCount--
            } else {
                rFloatingBan++
                queue = append(queue, 'D')
            }
        } else {
            if rFloatingBan > 0 {
                rFloatingBan--
                rCount--
            } else {
                dFloatingBan++
                queue = append(queue, 'R')
            }
        }
    }

    if rCount > 0 {
        return "Radiant"
    } else {
        return "Dire"
    }
}

func main() {
	// result: "Radiant"
	// senate := "RD"

	// result: "Dire"
	senate := "RDD"

	// result: 
	// senate := ""

	result := predictPartyVictory(senate)
	fmt.Printf("result = %v\n", result)
}

