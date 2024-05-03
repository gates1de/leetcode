package main
import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

func compareVersion(version1 string, version2 string) int {
	if len(version1) == 1 && len(version2) == 1 {
		if version1 == version2 {
			return 0
		} else if version1 > version2 {
			return 1
		}
		return -1
	}

	rep := regexp.MustCompile(`^0*`)
	version1 = rep.ReplaceAllString(version1, "")
	version2 = rep.ReplaceAllString(version2, "")

	rep = regexp.MustCompile(`\.0*\.`)
	version1 = rep.ReplaceAllString(version1, ".*.")
	version2 = rep.ReplaceAllString(version2, ".*.")

	rep = regexp.MustCompile(`\.0*`)
	version1 = rep.ReplaceAllString(version1, ".")
	version2 = rep.ReplaceAllString(version2, ".")

	rep = regexp.MustCompile(`\.{2}`)
	version1 = rep.ReplaceAllString(version1, "")
	version2 = rep.ReplaceAllString(version2, "")

	rep = regexp.MustCompile(`\*`)
	version1 = rep.ReplaceAllString(version1, "0")
	version2 = rep.ReplaceAllString(version2, "0")

	// fmt.Println(version1, version2)

	v1Split := strings.Split(version1, ".")
	v2Split := strings.Split(version2, ".")

	// fmt.Println(v1Split, v2Split)

	if len(v1Split) == len(v2Split) {
		for i := 0; i < len(v1Split); i++ {
			v1, _ := strconv.Atoi(v1Split[i])
			v2, _ := strconv.Atoi(v2Split[i])
			if v1 == v2 {
				continue
			} else if v1 < v2 {
				return -1
			}
			return 1
		}

		return 0
	} else if len(v1Split) < len(v2Split) {
		for i, v1Str := range v1Split {
			v1, _ := strconv.Atoi(v1Str)
			v2, _ := strconv.Atoi(v2Split[i])
			if v1 == v2 {
				continue
			} else if v1 < v2 {
				return -1
			}
			return 1
		}

		for _, v2Str := range v2Split[len(v1Split):] {
			v2, _ := strconv.Atoi(v2Str)
			if v2 > 0 {
				return -1
			}
		}
	} else {
		for i, v2Str := range v2Split {
			v1, _ := strconv.Atoi(v1Split[i])
			v2, _ := strconv.Atoi(v2Str)
			if v1 == v2 {
				continue
			} else if v1 < v2 {
				return -1
			}
			return 1
		}

		for _, v1Str := range v1Split[len(v2Split):] {
			v1, _ := strconv.Atoi(v1Str)
			if v1 > 0 {
				return 1
			}
		}
	}

	return 0
}

func main() {
	// result: 0
	// version1 := "1.01"
	// version2 := "1.001"

	// result: 0
	// version1 := "1.0"
	// version2 := "1.0.0"

	// result: -1
	version1 := "0.1"
	version2 := "1.1"

	// result: 
	// version1 := ""
	// version2 := ""

	result := compareVersion(version1, version2)
	fmt.Printf("result = %v\n", result)
}


