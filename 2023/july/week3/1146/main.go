package main
import (
	"fmt"
)

type SnapshotArray struct {
	Metadata [][][]int
	SnapshotId int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		Metadata: make([][][]int, length),
		SnapshotId: int(0),
	}
}

func (this *SnapshotArray) Set(index int, val int)  {
	if len(this.Metadata[index]) == 0 ||
		this.Metadata[index][len(this.Metadata[index]) - 1][0] != this.SnapshotId {
		this.Metadata[index] = append(this.Metadata[index], []int{this.SnapshotId, val})
		return
	}
	this.Metadata[index][len(this.Metadata[index]) - 1][1] = val
}

func (this *SnapshotArray) Snap() int {
	this.SnapshotId++
	return this.SnapshotId - 1
}

func (this *SnapshotArray) Get(index int, snapID int) int {
	return findLE(this.Metadata[index], snapID)
}

func findLE(data [][]int, snapID int) int {
	if len(data) == 0 {
		return 0
	}
	if data[0][0] > snapID {
		return 0
	}
	if data[len(data)-1][0] <= snapID {
		return data[len(data)-1][1]
	}

	l := int(0)
	r := len(data) - 1
	for l <= r {
		m := (l + r) >> 1
		if data[m][0] == snapID {
			return data[m][1]
		} else if data[m][0] < snapID {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return data[r][1]
}

func main() {
	// result: [null,null,0,null,5]
	length := int(3)
	operations := [][3]int{{0,0,5},{1,0,0},{0,0,6},{2,0,0}}

	// result: 
	// length := int(0)
	// operations := [][3]int{{0,0,0},{1,0,0},{2,0,0}}

	// result: 
	// length := int(0)
	// operations := [][3]int{{0,0,0},{1,0,0},{2,0,0}}

	obj := Constructor(length)
	for _, ope := range operations {
		if ope[0] == 0 {
			obj.Set(ope[1], ope[2])
		} else if ope[0] == 1 {
			fmt.Printf("obj.Snap() = %v\n", obj.Snap())
		} else if ope[0] == 2 {
			fmt.Printf("obj.Get(%v, %v) = %v\n", ope[1], ope[2], obj.Get(ope[1], ope[2]))
		}
	}
}

