package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Spreadsheet struct {
	Table [][]int
}

func Constructor(rows int) Spreadsheet {
	table := make([][]int, 26)
	for i := range table {
		table[i] = make([]int, rows)
	}
	return Spreadsheet{Table: table}
}

func (this *Spreadsheet) SetCell(cell string, value int)  {
	c, r := this.Parse(cell)
	this.Table[c][r] = value
}

func (this *Spreadsheet) ResetCell(cell string)  {
	c, r := this.Parse(cell)
	this.Table[c][r] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
	chars := []byte(formula)
	chars = chars[1:]

	strs := strings.Split(string(chars), "+")
	left := strs[0]
	right := strs[1]

	leftVal, err := strconv.Atoi(left)
	if err != nil {
			c, r := this.Parse(left)
			leftVal = this.Table[c][r]
	}

	rightVal, err := strconv.Atoi(right)
	if err != nil {
			c, r := this.Parse(right)
			rightVal = this.Table[c][r]
	}

	return leftVal + rightVal
}

func (this *Spreadsheet) Parse(cell string) (column int, row int) {
	chars := []byte(cell)
	column = int(chars[0]) - 65
	r, _ := strconv.Atoi(string(chars[1:]))
	if r == len(this.Table[0]) {
		r -= 1
	}
	row = r

	return
}

func main() {
	// result: [null, 12, null, 16, null, 25, null, 15]
	rows := int(3)
	operations := [][]string{
		{"2", "=5+7"},
		{"0", "A1", "10"},
		{"2", "=A1+6"},
		{"0", "B2", "15"},
		{"2", "=A1+B2"},
		{"1", "A1"},
		{"2", "=A1+B2"},
	}

	// result: []
	// rows := int(0)
	// operations := [][]int{}

	obj := Constructor(rows)
	for _, ope := range operations {
		if ope[0] == "0" {
			val, _ := strconv.Atoi(ope[2])
			obj.SetCell(ope[1], val)
		} else if ope[0] == "1" {
			obj.ResetCell(ope[1])
		} else if ope[0] == "2" {
			formula := ope[1]
			fmt.Printf("obj.GetValue(%v) = %v\n", formula, obj.GetValue(formula))
		}
	}
}

