package main
import (
	"fmt"
	"sort"
)

type SeatManager struct {
	ReservableSeats []int
	ReservedSeats map[int]bool
}

func Constructor(n int) SeatManager {
	reservableSeats := make([]int, n)
	for i, _ := range reservableSeats {
		reservableSeats[i] = i + 1
	}

	reservedSeats := make(map[int]bool, n)
	return SeatManager{ReservableSeats: reservableSeats, ReservedSeats: reservedSeats}
}

func (this *SeatManager) Reserve() int {
	if len(this.ReservableSeats) == 0 {
		return -1
	}

	result := this.ReservableSeats[0]
	this.ReservableSeats = this.ReservableSeats[1:]
	this.ReservedSeats[result] = true
	return result
}

func (this *SeatManager) Unreserve(seatNumber int)  {
	if !this.ReservedSeats[seatNumber] {
		return
	}

	this.ReservedSeats[seatNumber] = false
	i := sort.Search(len(this.ReservableSeats), func(i int) bool {
		return this.ReservableSeats[i] >= seatNumber
	})

	if i < len(this.ReservableSeats) && this.ReservableSeats[i] == seatNumber {
		return
	}

    this.ReservableSeats = append(this.ReservableSeats, 0)
    copy(this.ReservableSeats[i + 1:], this.ReservableSeats[i:])
    this.ReservableSeats[i] = seatNumber
}

func main() {
	// result: [null, 1, 2, null, 2, 3, 4, 5, null]
	n := int(5)
	operations := [][]int{{0,0},{0,0},{1,2},{0,0},{0,0},{0,0},{0,0},{1,5}}

	// result: 
	// n := int(0)
	// operations := [][]int{{0,0},{1,0}}

	obj := Constructor(n)
	for _, ope := range operations {
		if ope[0] == 0 {
			fmt.Printf("obj.Reserve() = %v\n", obj.Reserve())
		} else if ope[0] == 1 {
			obj.Unreserve(ope[1])
		}
	}
}

