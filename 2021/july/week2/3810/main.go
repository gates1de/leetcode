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
	return MedianFinder{}
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
	fmt.Printf("init = %v\n", obj)
	// obj.AddNum(1)
	// obj.AddNum(2)
	// fmt.Printf("obj.FindMedian() = %v\n", obj.FindMedian())
	// obj.AddNum(3)
	// fmt.Printf("obj.FindMedian() = %v\n", obj.FindMedian())

	obj.AddNum(1004)
	obj.AddNum(20002)
	fmt.Printf("obj.FindMedian() = %v\n", obj.FindMedian())
	obj.AddNum(345)
	fmt.Printf("obj.FindMedian() = %v\n", obj.FindMedian())
	obj.AddNum(1)
	fmt.Printf("obj.FindMedian() = %v\n", obj.FindMedian())
}

