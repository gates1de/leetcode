package main
import (
	"fmt"
)

type ParkingSystem struct {
	Big int
	Medium int
	Small int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{Big: big, Medium: medium, Small: small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 && this.Big > 0 {
		this.Big--
		return true
	}
	if carType == 2 && this.Medium > 0 {
		this.Medium--
		return true
	}
	if carType == 3 && this.Small > 0 {
		this.Small--
		return true
	}

	return false
}

func main() {
	// result: [null, true, true, false, false]
	// spaces := []int{1,1,0}
	// carTypes := []int{1,2,3,1}

	// result: [null, true, true, false, true]
	spaces := []int{2,1,0}
	carTypes := []int{1,2,3,1}

	// result: 
	// spaces := []int{}
	// carTypes := []int{}

	obj := Constructor(spaces[0], spaces[1], spaces[2])
	for _, carType := range carTypes {
		fmt.Printf("obj.AddCar(%v) = %v\n", carType, obj.AddCar(carType))
	}
}

