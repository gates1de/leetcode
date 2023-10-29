package main
import (
	"fmt"
)

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
    if buckets <= 1 ||  minutesToDie == 0 || minutesToTest == 0 {
        return 0
    }

    t := minutesToTest / minutesToDie

    result := int(1)
    for true {
        if buckets <= pow(t + 1, result) {
            break
        }
        result++
    }

    return result
}

func pow(x int, y int) int {
    result := int(1)
    for i := 0; i < y; i++ {
        result *= x
    }
    return result
}

func main() {
	// result: 2
	// buckets := int(4)
	// minutesToDie := int(15)
	// minutesToTest := int(15)

	// result: 2
	buckets := int(4)
	minutesToDie := int(15)
	minutesToTest := int(30)

	// result: 
	// buckets := int(0)
	// minutesToDie := int(0)
	// minutesToTest := int(0)

	result := poorPigs(buckets, minutesToDie, minutesToTest)
	fmt.Printf("result = %v\n", result)
}

