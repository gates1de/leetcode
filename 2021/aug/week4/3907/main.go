package main
import (
	"fmt"
	"sort"
)

const modulo = int(1e9 + 7)

func rectangleArea(rectangles [][]int) int {
	xs := getXs(rectangles)

	idxs := make(map[int]int, 2 * len(xs))
	for idx, x := range xs {
		idxs[x] = idx
	}

	labels := getLabels(rectangles)

	count := make([]int, len(xs))

	curY, curXSum, area := 0, 0, 0

	for _, l := range labels {
		y, x1, x2, sig := l[0], l[1], l[2], l[3]
		area += (y - curY) * curXSum
		curY = y
		curXSum = 0
		for i := idxs[x1]; i < idxs[x2]; i++ {
			count[i] += sig
		}
		for i := 0; i+1 < len(count); i++ {
			if count[i] > 0 {
				curXSum += xs[i+1] - xs[i]
			}
		}
	}

	return area % modulo
}

func getXs(rects [][]int) []int {
	size := len(rects)
	xs := make([]int, 0, size*2)
	xMap := make(map[int]bool, size*2)
	for _, r := range rects {
		xMap[r[0]] = true
		xMap[r[2]] = true
	}
	for k := range xMap {
		xs = append(xs, k)
	}
	sort.Ints(xs)
	return xs
}

func getLabels(rects [][]int) [][]int {
	labels := make([][]int, 0, 2*len(rects))
	for _, r := range rects {
		x1, y1, x2, y2 := r[0], r[1], r[2], r[3]
		labels = append(labels,
			[]int{y1, x1, x2, 1},
			[]int{y2, x1, x2, -1},
		)
	}

	sort.Slice(labels, func(i int, j int) bool {
		return labels[i][0] < labels[j][0]
	})

	return labels
}

// Other solution
// func rectangleArea(rectangles [][]int) int {
// 
// 	xAxis, yAxis := make([]int, len(rectangles)*2), make([]int, len(rectangles)*2)
// 	for i := range rectangles {
// 		xAxis[i*2], xAxis[i*2+1] = rectangles[i][0], rectangles[i][2]
// 		yAxis[i*2], yAxis[i*2+1] = rectangles[i][1], rectangles[i][3]
// 	}
// 
// 	sort.Ints(xAxis)
// 	sort.Ints(yAxis)
// 
// 	removeDuplicated := func(l []int) []int {
// 
// 		l1 := l[:1]
// 		for _, i := range l[1:] {
// 			if i != l1[len(l1)-1] {
// 				l1 = append(l1, i)
// 			}
// 		}
// 		return l1
// 	}
// 
// 	xAxis = removeDuplicated(xAxis)
// 	yAxis = removeDuplicated(yAxis)
// 	m := make([][]bool, len(xAxis))
// 	for i := range m {
// 		m[i] = make([]bool, len(yAxis))
// 	}
// 
// 	for i := range rectangles {
// 
// 		x0, x1 := sort.SearchInts(xAxis, rectangles[i][0]), sort.SearchInts(xAxis, rectangles[i][2])
// 		y0, y1 := sort.SearchInts(yAxis, rectangles[i][1]), sort.SearchInts(yAxis, rectangles[i][3])
// 
// 		for x := x0; x < x1; x++ {
// 			for y := y0; y < y1; y++ {
// 				m[x][y] = true
// 			}
// 		}
// 	}
// 
// 	ans := 0
// 	const mod = 1000000007
// 	for i := 0; i < len(xAxis); i++ {
// 		for j := 0; j < len(yAxis); j++ {
// 
// 			if !m[i][j] {
// 				continue
// 			}
// 
// 			ans = (ans + (xAxis[i+1]-xAxis[i])*(yAxis[j+1]-yAxis[j])) % mod
// 		}
// 	}
// 
// 	return ans
// }

// Out of memory
func ngSolution(rectangles [][]int) int {
	if len(rectangles) == 1 {
		rect := rectangles[0]
		return (rect[2] - rect[0]) * (rect[3] - rect[1]) % modulo
	}

	m := map[[2]int]bool{}
	for _, rect := range rectangles {
		for i := rect[0]; i < rect[2]; i++ {
			for j:= rect[1]; j < rect[3]; j++ {
				m[[2]int{i, j}] = true
			}
		}
	}
	result := int(0)
	for range m {
		result++
	}
	return result % modulo
}

func main() {
	// result: 6
	rectangles := [][]int{
		{0, 0, 2, 2},
		{1, 0, 2, 3},
		{1, 0, 3, 1},
	}

	// result: 49
	// rectangles := [][]int{
	// 	{0, 0, 1000000000, 1000000000},
	// }

	// result: 
	// rectangles := [][]int{
	// 	{, , , },
	// 	{, , , },
	// 	{, , , },
	// }

	result := rectangleArea(rectangles)
	fmt.Printf("result = %v\n", result)
}

