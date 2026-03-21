package main

import (
	"fmt"
)

const mod = int(1e9 + 7)

type Fancy struct {
   V []int
   A int
   B int
}

func Constructor() Fancy {
	return Fancy{
		V: make([]int, 0),
		A: 1,
		B: 0,
	}
}

func (this *Fancy) Append(val int)  {
	adjustedVal := ((val - this.B + mod) % mod) * this.Inverse(this.A) % mod
	this.V = append(this.V, adjustedVal)
}

func (this *Fancy) AddAll(inc int)  {
	this.B = (this.B + inc) % mod
}

func (this *Fancy) MultAll(m int)  {
	this.A = (this.A * m) % mod
	this.B = (this.B * m) % mod
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= len(this.V) {
		return -1
	}

	result := (this.A * this.V[idx] % mod + this.B) % mod
	return result
}

func (this *Fancy) QuickMul(x, y int) int {
	result := 1
	multiple := x
	for y > 0 {
		if y & 1 != 0 {
			result = (result * multiple) % mod
		}

		multiple = (multiple * multiple) % mod
		y >>= 1
	}

	return result
}

func (this *Fancy) Inverse(x int) int {
	return this.QuickMul(x, mod - 2)
}

func main() {
	// result: [null, null, null, null, null, 10, null, null, null, 26, 34, 20]
	operations := [][]int{{0,2},{1,3},{0,7},{2,2},{3,0},{1,3},{0,10},{2,2},{3,0},{3,1},{3,2}}

	obj := Constructor()
	for _, ope := range operations {
		if ope[0] == 0 {
			obj.Append(ope[1])
		} else if ope[0] == 1 {
			obj.AddAll(ope[1])
		} else if ope[0] == 2 {
			obj.MultAll(ope[1])
		} else if ope[0] == 3 {
			fmt.Printf("obj.GetIndex(%d) = %v\n", ope[1], obj.GetIndex(ope[1]))
		}
	}
}

