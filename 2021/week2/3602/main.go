package main
import (
	"fmt"
	"sort"
)

// Accepted
func numRescueBoats(people []int, limit int) int {
	sort.Slice(people, func(i, j int) bool { return people[i] < people[j] })
	// fmt.Printf("sortedP = %v\n", people)
	i := int(0)
	j := len(people) - 1
	result := int(0)
    for i <= j {
        result += 1
        if people[i] + people[j] <= limit {
            i += 1
		}
        j -= 1
	}
    return result
}

// Timeout Limit Exceeded
func mySolution(people []int, limit int) int {
	if len(people) == 0 {
		return 0
	}
	sort.Slice(people, func(i, j int) bool { return people[i] > people[j] })
	fmt.Printf("sortedP = %v\n", people)
	result := int(0)
	currentLimit := limit

	for i, num := range people {
		if num == 0 {
			continue
		}
		if num == limit {
			result += 1
			continue
		}

		if currentLimit < limit {
			remainPeople := people[i:]
			index := searchMore(remainPeople, currentLimit)
			if index >= 0 {
				lessNum := remainPeople[index]
				people[i + index] = 0
				currentLimit -= lessNum
			}
			currentLimit = limit
			result += 1
		}
		if people[i] == 0 {
			continue
		}

		if num < limit {
			currentLimit -= num % limit
			people[i] = 0
		} else {
			currentLimit = limit
			result += 1
		}
		// fmt.Printf("boat %v remaining: %v people can get on the boat -> boats %v\n", currentLimit, num, result)
	}

	if currentLimit < limit {
		result += 1
	}
	return result
}

func searchMore(s []int, target int) int {
	result := int(-1)
	for i, _ := range s {
		val := s[len(s) - i - 1]
		if val <= 0 {
			continue
		}
		if val == target {
			return i
		}
		if result < val && val < target {
			result = i
		}
	}
	return result
}

func main() {
	// people := []int{1, 2}
	// limit := int(3)

	// people := []int{3, 2, 2, 1}
	// limit := int(3)

	// people := []int{3, 5, 3, 4}
	// limit := int(5)

	// people := []int{21, 40, 16, 24, 30}
	// limit := int(50)

	// people := []int{3,8,4,9,2,2,7,1,6,10,6,7,1,7,7,6,4,4,10,1}
	// limit := int(10)

	people := []int{3, 2, 3, 2, 2}
	limit := int(6)

	result := numRescueBoats(people, limit)
	fmt.Printf("result = %v\n", result)
}

