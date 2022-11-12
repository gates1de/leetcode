package main
import (
	"fmt"
	"math"
	"sort"
)

const sigDigits = float64(10000)

type MedianFinder struct {
    Nums []int
}

func Constructor() MedianFinder {
	return MedianFinder{ Nums: make([]int, 0, 50000) }
}

func (this *MedianFinder) AddNum(num int)  {
    i := sort.Search(len(this.Nums), func(i int) bool { return this.Nums[i] >= num })
    this.Nums = append(this.Nums, 0)
    copy(this.Nums[i + 1:], this.Nums[i:])
    this.Nums[i] = num
}

func (this *MedianFinder) FindMedian() float64 {
    n := len(this.Nums)
    if n % 2 == 0 {
        result := (float64(this.Nums[n / 2 - 1]) + float64(this.Nums[n / 2])) / 2
        return math.Round(result * sigDigits) / sigDigits
    }
    return float64(this.Nums[n / 2])
}

func main() {
	obj := Constructor()

	// result: [null, null, null, 1.5, null, 2.0]
	// controls := [][]int{{0, 1}, {0, 2}, {1, 0}, {0, 3}, {1, 0}}

	// result: [null,null,1768.00000,null,2474.50000,null,2366.00000]
	controls := [][]int{{0, 1768}, {1, 0}, {0, 3181}, {1, 0}, {0, 2366}, {1, 0}}

	// result: 
	// controls := [][]int{}

	for _, control := range controls {
		if control[0] == 0 {
			obj.AddNum(control[1])
		} else if control[0] == 1 {
			fmt.Printf("obj.FindMedian() = %v\n", obj.FindMedian())
		}
	}
}

