package bit

// import "fmt"

// REF:
// https://dev.to/ruarfff/day-10-create-sorted-array-through-instructions-2bo4
// https://algo-logic.info/binary-indexed-tree/
// https://qiita.com/DaikiSuyama/items/7295f5160a51684554a7 

type BinaryIndexedTree struct {
	length int
	tree []int
}

func New(length int) *BinaryIndexedTree {
	bit := new(BinaryIndexedTree)
	bit.length = length
	bit.tree = make([]int, length)
	// fmt.Printf("bit: {length: %v, tree: %v}\n", bit.length, bit.tree)
	return bit
}

func (bit *BinaryIndexedTree) GetCost(targetIndex int) int {
	result := int(0)
	for targetIndex >= 1 {
		// fmt.Printf("target index = %v\n", targetIndex)
		result += bit.tree[targetIndex]
		targetIndex -= targetIndex & -targetIndex
		// fmt.Printf("current cost = %v\n", result)
	}
	return result
}

func (bit *BinaryIndexedTree) Update(targetIndex int, val int) {
	for targetIndex < bit.length {
		bit.tree[targetIndex] += val
		// fmt.Printf("bit.tree[%v] = %v\n", targetIndex, bit.tree[targetIndex])
		targetIndex += targetIndex & -targetIndex
	}
}

