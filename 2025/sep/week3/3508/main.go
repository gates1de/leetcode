package main

import (
	"fmt"
)

type Router struct {
	MemoryLimit int
	Histories [][3]int
	Packets map[[3]int]bool
	DestPackets map[int][]int
}

func Constructor(memoryLimit int) Router {
	return Router{
		MemoryLimit: memoryLimit,
		Histories: make([][3]int, 0),
		Packets: make(map[[3]int]bool),
		DestPackets: make(map[int][]int),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	packet := [3]int{source, destination, timestamp}
	if _, exists := this.Packets[packet]; exists {
		return false
	}

	if len(this.Packets) == this.MemoryLimit {
		_ = this.ForwardPacket()
	}

	this.Packets[packet] = true
	this.Histories = append(this.Histories, packet)
	this.DestPackets[destination] = append(this.DestPackets[destination], timestamp)
	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.Histories) == 0 {
		return nil
	}

	packet := this.Histories[0]
	this.Histories = this.Histories[1:]
	delete(this.Packets, packet)
	this.DestPackets[packet[1]] = this.DestPackets[packet[1]][1:]
	return []int{packet[0], packet[1], packet[2]}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	timestamps := this.DestPackets[destination]
	if len(timestamps) == 0 {
		return 0
	}

	left := int(0)
	right := len(timestamps) - 1
	for left < right {
		mid := (left + right) / 2
		if timestamps[mid] < startTime {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if timestamps[left] < startTime {
		return 0
	}

	left2 := left
	right2 := len(timestamps) - 1
	for left2 < right2 {
		mid := (left2 + right2 + 1) / 2
		if timestamps[mid] > endTime {
			right2 = mid - 1
		} else {
			left2 = mid
		}
	}

	if timestamps[left2] > endTime {
		return 0
	}

	return left2 - left + 1
}

func main() {
	// result: [null, true, true, false, true, true, [2, 5, 90], true, 1]
	// memoryLimit := int(3)
	// operations := [][]int{
	// 	{0, 1, 4, 90},
	// 	{0, 2, 5, 90},
	// 	{0, 1, 4, 90},
	// 	{0, 3, 5, 95},
	// 	{0, 4, 5, 105},
	// 	{1},
	// 	{0, 5, 2, 110},
	// 	{2, 5, 100, 110},
	// }

	// result: [null, true, [7, 4, 90], []]
	memoryLimit := int(2)
	operations := [][]int{
		{0, 7, 4, 90},
		{1},
		{1},
	}

	// result: []
	// memoryLimit := int(0)
	// operations := [][]int{}

	obj := Constructor(memoryLimit)
	for _, ope := range operations {
		if ope[0] == 0 {
			source := ope[1]
			destination := ope[2]
			timestamp := ope[3]

			fmt.Printf(
				"obj.AddPacket(%v, %v, %v) = %v\n",
				source,
				destination,
				timestamp,
				obj.AddPacket(source, destination, timestamp),
			)
		} else if ope[0] == 1 {
			fmt.Printf("obj.ForwardPacket() = %v\n", obj.ForwardPacket())
		} else if ope[0] == 2 {
			destination := ope[1]
			startTime := ope[2]
			endTime := ope[3]

			fmt.Printf(
				"obj.GetCount(%v, %v, %v) = %v\n",
				destination,
				startTime,
				endTime,
				obj.GetCount(destination, startTime, endTime),
			)
		}
	}
}

