package main
import (
	"fmt"
	"sort"
)

type TimeMap struct {
	ValuesMap map[string][]string
	TimestampIndexes map[string][]int
}

func Constructor() TimeMap {
	return TimeMap{ValuesMap: map[string][]string{}, TimestampIndexes: map[string][]int{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int)  {
	if _, isExists := this.ValuesMap[key]; !isExists {
		this.ValuesMap[key] = []string{value}
		this.TimestampIndexes[key] = []int{timestamp}
		return
	}

	values := this.ValuesMap[key]
	indexes := this.TimestampIndexes[key]
	insertIndex := sort.SearchInts(indexes, timestamp) 

	if insertIndex >= len(indexes) {
		values = append(values, value)
		indexes = append(indexes, timestamp)
		this.ValuesMap[key] = values
		this.TimestampIndexes[key] = indexes
		return
	}

	newIndexes := make([]int, len(indexes) + 1)
	copy(newIndexes[:insertIndex], indexes[:insertIndex])
	newIndexes[insertIndex] = timestamp
	copy(newIndexes[insertIndex + 1:], indexes[:insertIndex + 1])

	newValues := make([]string, len(values) + 1)
	copy(newValues[:insertIndex], values[:insertIndex])
	newValues[insertIndex] = value
	copy(newValues[insertIndex + 1:], values[:insertIndex + 1])

	this.ValuesMap[key] = newValues
	this.TimestampIndexes[key] = newIndexes
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, isExists := this.ValuesMap[key]; !isExists {
		return ""
	}

	values := this.ValuesMap[key]
	indexes := this.TimestampIndexes[key]
	targetIndex := sort.SearchInts(indexes, timestamp) 

	if targetIndex >= len(values) {
		return values[len(values) - 1]
	}
	if timestamp < indexes[targetIndex] {
		if targetIndex == 0 {
			return ""
		}
		return values[targetIndex - 1]
	}

	return values[targetIndex]
}

func main() {
	type KeyValue struct {
		SetOrGet int // 0 -> set, 1 -> get
		Key string
		Value string // in case of get -> ""
		Timestamp int
	}

	// result: [null,null,"bar","bar",null,"bar2","bar2"]
	// keyValues := []KeyValue{
	// 	KeyValue{SetOrGet: 0, Key: "foo", Value: "bar", Timestamp: 1},
	// 	KeyValue{SetOrGet: 1, Key: "foo", Value: "", Timestamp: 1},
	// 	KeyValue{SetOrGet: 1, Key: "foo", Value: "", Timestamp: 3},
	// 	KeyValue{SetOrGet: 0, Key: "foo", Value: "bar2", Timestamp: 4},
	// 	KeyValue{SetOrGet: 1, Key: "foo", Value: "", Timestamp: 4},
	// 	KeyValue{SetOrGet: 1, Key: "foo", Value: "", Timestamp: 5},
	// }

	// result: [null,null,"bar","bar",null,"bar2","bar2","bar"]
	keyValues := []KeyValue{
		KeyValue{SetOrGet: 0, Key: "foo", Value: "bar", Timestamp: 1},
		KeyValue{SetOrGet: 1, Key: "foo", Value: "", Timestamp: 1},
		KeyValue{SetOrGet: 1, Key: "foo", Value: "", Timestamp: 3},
		KeyValue{SetOrGet: 0, Key: "foo", Value: "bar2", Timestamp: 4},
		KeyValue{SetOrGet: 1, Key: "foo", Value: "", Timestamp: 4},
		KeyValue{SetOrGet: 1, Key: "foo", Value: "", Timestamp: 5},
		KeyValue{SetOrGet: 1, Key: "foo", Value: "", Timestamp: 1},
	}

	// result: 
	// keyValues := [][]KeyValue{
	// 	KeyValue{SetOrGet: 0, Key: "foo", Value: "bar", Timestamp: 1},
	// }

	obj := Constructor()
	for _, keyValue := range keyValues {
		if keyValue.SetOrGet == 0 {
			obj.Set(keyValue.Key, keyValue.Value, keyValue.Timestamp)
		} else if keyValue.SetOrGet == 1 {
			fmt.Printf("obj.Get(%v,%v) = %v\n", keyValue.Key, keyValue.Timestamp, obj.Get(keyValue.Key, keyValue.Timestamp))
		}
	}
}

