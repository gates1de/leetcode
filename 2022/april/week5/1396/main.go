package main
import (
	"fmt"
)

type UndergroundSystem struct {
	Customers map[int]string
	Stations map[string]map[int]int
	Trips map[string][]int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		Customers: map[int]string{},
		Stations: map[string]map[int]int{},
		Trips: map[string][]int{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
	this.Customers[id] = stationName
	stationTimes := this.Stations[stationName]
	if stationTimes == nil {
		stationTimes = map[int]int{}
	}
	stationTimes[id] = t
	this.Stations[stationName] = stationTimes
	// fmt.Println("----------- check in -----------")
	// fmt.Println(this.Customers)
	// fmt.Println(this.Stations)
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
	startStation := this.Customers[id]
	startStationTimes := this.Stations[startStation]
	startStationTime := startStationTimes[id]
	tripKey := fmt.Sprintf("%v:%v", startStation, stationName)
	trips := this.Trips[tripKey]
	if trips == nil {
		trips = []int{}
	}
	trips = append(trips, t - startStationTime)
	this.Trips[tripKey] = trips
	// fmt.Println("----------- check out -----------")
	// fmt.Println(this.Trips)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	result := float64(0)
	tripKey := fmt.Sprintf("%v:%v", startStation, endStation)
	trips := this.Trips[tripKey]
	if trips == nil || len(trips) == 0 {
		return result
	}
	for _, t := range trips {
		result += float64(t)
	}
	result /= float64(len(trips))
	return result
}

func main() {
	obj := Constructor()

	// result: [null,null,null,null,null,null,null,14.00000,11.00000,null,11.00000,null,12.00000]
	// obj.CheckIn(45, "Leyton", 3)
	// obj.CheckIn(32, "Paradise", 8)
	// obj.CheckIn(27, "Leyton", 10)
	// obj.CheckOut(45, "Waterloo", 15)
	// obj.CheckOut(27, "Waterloo", 20)
	// obj.CheckOut(32, "Cambridge", 22)
	// fmt.Printf("obj.GetAverageTime(%v, %v) = %v\n", "Paradise", "Cambridge", obj.GetAverageTime("Paradise", "Cambridge"))
	// fmt.Printf("obj.GetAverageTime(%v, %v) = %v\n", "Leyton", "Waterloo", obj.GetAverageTime("Leyton", "Waterloo"))
	// obj.CheckIn(10, "Leyton", 24)
	// fmt.Printf("obj.GetAverageTime(%v, %v) = %v\n", "Leyton", "Waterloo", obj.GetAverageTime("Leyton", "Waterloo"))
	// obj.CheckOut(10, "Waterloo", 38)
	// fmt.Printf("obj.GetAverageTime(%v, %v) = %v\n", "Leyton", "Waterloo", obj.GetAverageTime("Leyton", "Waterloo"))

	// result: [null,null,null,5.00000,null,null,5.50000,null,null,6.66667]
	obj.CheckIn(10, "Leyton", 3)
	obj.CheckOut(10, "Paradise", 8)
	fmt.Printf("obj.GetAverageTime(%v, %v) = %v\n", "Leyton", "Paradise", obj.GetAverageTime("Leyton", "Paradise"))
	obj.CheckIn(5, "Leyton", 10)
	obj.CheckOut(5, "Paradise", 16)
	fmt.Printf("obj.GetAverageTime(%v, %v) = %v\n", "Leyton", "Paradise", obj.GetAverageTime("Leyton", "Paradise"))
	obj.CheckIn(2, "Leyton", 21)
	obj.CheckOut(2, "Paradise", 30)
	fmt.Printf("obj.GetAverageTime(%v, %v) = %v\n", "Leyton", "Paradise", obj.GetAverageTime("Leyton", "Paradise"))

	// obj.CheckIn(0, "", 0)
	// obj.CheckOut(0, "", 0)
	// fmt.Printf("obj.GetAverageTime(%v, %v) = %v\n", "", "", obj.GetAverageTime("", ""))
}

