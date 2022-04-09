package main
import (
	"fmt"
)


func maxArea(height []int) int {
    n := len(height)
    maxVolume := int(0)

    left := int(0)
    right := n - 1
    for left < right {
        minBar := min(height[left], height[right])
        capacity := minBar * (right - left)
        if maxVolume < capacity {
            maxVolume = capacity
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxVolume
}

// Slow
func mySolution(height []int) int {
	maxSum := int(0)
	maxHeight := int(0)
	for i, h1 := range height {
		sum := int(0)
		if h1 > maxHeight {
			maxHeight = h1
		} else {
			continue
		}

		for j := i + 1; j < len(height); j++ {
			h2 := height[j]
			sum = (j - i) * min(h1, h2)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 49
	// height := []int{1,8,6,2,5,4,8,3,7}

	// result: 1
	// height := []int{1, 1}

	// result: 16
	// height := []int{4, 3, 2, 1, 4}

	// result: 840311
	height := []int{6801,4040,7716,493,526,2755,957,1298,2477,6189,6442,8476,4745,8663,2812,8476,3802,7743,7746,2513,4434,3625,2470,4902,8135,5773,5457,4527,6798,8800,8824,9067,3494,915,68,4671,5868,7224,371,4667,38,5958,5616,945,6694,6584,1912,1184,45,9767,5674,7407,7451,6302,7674,8803,9528,1450,2176,7650,9998,3815,7815,4096,221,10000,5940,7124,7201,6648,6892,6805,9659,4132,1054,6037,7620,7515,2025,6628,1067,2431,175,2962,9361,3271,4362,7932,8800,2288,2295,3600,8283,3356,2265,498,9649,9290,8406,7133,6224,6545,39,8623,8549,1733,3448,5100,3112,2852,9945}

	// result: 
	// height := []int{}

	result := maxArea(height)
	fmt.Printf("result = %v\n", result)
}

