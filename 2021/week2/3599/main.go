package main
import (
	"fmt"
	"./bit"
)

// Accepted
func createSortedArray(instructions []int) int {
	bit := bit.New(getMax(instructions) + 2)
	cost := int(0)
	for i, val := range instructions {
		fmt.Printf("i = %v, val = %v\n", i, val)
		lessCost := bit.GetCost(val)
		greaterCost := i - bit.GetCost(val + 1)
		fmt.Printf("min(%v, %v) = %v\n", lessCost, greaterCost, min(lessCost, greaterCost))
		cost += min(lessCost, greaterCost)
		bit.Update(val + 1, 1)
	}
	return cost % (1e9 + 7)
}

// Timeout Limit Exceeded
func mySolution(instructions []int) int {
	sortedList := make([]int, len(instructions))
	cost := int(0)
	for i, val := range instructions {
		// fmt.Printf("val = %v, sortedList = %v\n", val, sortedList)

		if i == 0 {
			sortedList[0] = val
			continue
		}

		sortedListLength := len(sortedList)
		if val <= sortedList[0] {
			copy(sortedList[1:sortedListLength], sortedList[:sortedListLength])
			sortedList[0] = val
			// fmt.Printf("No cost insert: val is min %v\n", sortedList)
			continue
		}

		lessCost := int(0)
		greaterCost := int(0)
		for j, v := range sortedList {
			if v == 0 {
				continue
			}

			if v < val {
				lessCost++
			}
			reverseV := sortedList[i - j]
			if reverseV > val {
				greaterCost++
			}
		}
		copy(sortedList[lessCost + 1:sortedListLength], sortedList[lessCost:sortedListLength])
		sortedList[lessCost] = val
		cost += min(lessCost, greaterCost)
		// fmt.Printf("val %v: cost = min(%v, %v), total cost = %v\n", val, lessCost, greaterCost, cost)
	}

	return cost
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func getMax(s []int) int {
	result := int(0)
	for _, v := range s {
		if v > result {
			result = v
		}
	}
	return result
}

func count(list []int, target int) int {
	result := int(0)
	for _, v := range list {
		if target < v {
			break
		}
		if target == v {
			result++
		}
	}
	return result
}

func copySlice(target []int) []int {
	result := make([]int, len(target))
	copy(result, target)
	return result
}

func main() {
	instructions := []int{1, 5, 6, 2}
	// instructions := []int{1, 2, 3, 6, 5, 4}
	// instructions := []int{1, 3, 3, 3, 2, 4, 2, 1, 2}
	// instructions := []int{1,81615,2,81614,3,81613,4,81612,5,81611,6,81610,7,81609,8,81608,9,81607,10,81606,11,81605,12,81604,13,81603,14,81602,15,81601,16,81600,17,81599,18,81598,19,81597,20,81596,21,81595,22,81594,23,81593,24,81592,25,81591,26,81590,27,81589,28,81588,29,81587,30,81586,31,81585,32,81584,33,81583,34,81582,35,81581,36,81580,37,81579,38,81578,39,81577,40,81576,41,81575,42,81574,43,81573,44,81572,45,81571,46,81570,47,81569,48,81568,49,81567,50,81566,51,81565,52,81564,53,81563,54,81562,55,81561,56,81560,57,81559,58,81558,59,81557,60,81556,61,81555,62,81554,63,81553,64,81552,65,81551,66,81550,67,81549,68,81548,69,81547,70,81546,71,81545,72,81544,73,81543,74,81542}
	// instructions := []int{1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3}
	// instructions := []int{316,326,221,216,135,383,75,212,244,280,176,323,338,427,193,274,443,272,284,20,189,403,458,14,372,126,388,157,318,164,317,376,399,384,6,94,142,294,235,166,371,179,187,77,21,115,455,342,7,346,352,159,405,373,82,232,411,426,137,70,287,98,66,44,153,307,386,312,139,188,13,15,351,369,185,141,205,328,130,30,358,191,226,209,200,55,74,36,10,385,275,278,2,434,214,207,359,303,119,334,95,436,450,156,234,168,167,233,46,127,344,453,41,268,315,412,378,273,254,194,177,395,80,297,64,391,155,110,60,114,117,81,79,39,109,265,367,171,8,140,72,345,314,128,17,299,350,421,457,444,277,73,283,368,165,174,58,448,393,144,172,146,148,292,361,251,406,92,445,97,18,196,96,101,16,223,304,332,397,264,452,163,123,35,84,416,348,198,51,290,430,428,68,113,203,263,50,446,336,341,252,211,247,429,365,5,106,22,43,152,279,225,136,145,231,349,71,99,288,437,298,158,394,415,293,29,289,291,425,63,121,451,4,34,249,111,59,213,31,201,381,370,362,230,390,353,56,377,204,308,217,408,132,150,88,321,324,364,301,398,319,93,169,133,266,281,291,134,138,224,285,161,105,236,61,310,202,347,333,242,240,433,183,440,409,107,87,103,32,12,38,86,19,241,89,343,313,173,327,325,118,441,124,78,442,401,396,354,116,33,379,355,215,76,243,122,454,48,320,257,90,366,417,120,129,295,375,382,330,182,389,356,387,181,186,238,357,184,53,147,260,37,400,1,25,175,65,131,435,67,210,248,407,62,151,392,49,40,11,439,335,239,162,410,125,154,250,237,269,28,329,432,423,52,100,306,229,380,255,195,42,23,112,322,206,197,302,24,57,192,300,190,170,449,108,438,180,222,69,418,363,402,431,267,360,218,404,413,54,422,374,271,27,83,245,104,160,256,102,282,340,9,424,246,419,47,178,258,337,85,143,339,447,305,26,286,45,219,91,414,311,149,208,276,227,259,456,420,331,199,220,262,309,253,296,3,228,270,261}
	result := createSortedArray(instructions)
	fmt.Printf("result = %v\n", result)
}

