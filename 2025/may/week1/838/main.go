package main
import (
	"fmt"
)

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	resultBytes := make([]byte, n)

  for i := 0; i < len(resultBytes); i++ {
    resultBytes[i] = '.'
  }

  i := int(0)
  for i < n {
    for i < n && dominoes[i] != '.' {
      resultBytes[i] = dominoes[i]
      i++
    }

    j := i
    for j < n && dominoes[j] == '.' {
      j++
    }

    if i == 0 && j == n {
      return dominoes
    }

    if i == 0 {
      if dominoes[j] == 'L' {
          for k := 0; k < j; k++ {
              resultBytes[k] = dominoes[j]
          }
      }
    } else if j == n {
      if dominoes[i - 1] =='R' {
        for k := i; k < j; k++ {
          resultBytes[k] = dominoes[i - 1]
        }
      }
    } else if dominoes[i - 1] == dominoes[j] {
      for k := i; k < j; k++ {
        resultBytes[k] = dominoes[j]
      }
    } else if dominoes[i - 1] == 'R' && dominoes[j] == 'L' {
      ii, jj := i, j - 1
      for ii < jj {
        resultBytes[ii] = dominoes[i - 1]
        resultBytes[jj] = dominoes[j]
        ii++
        jj--
      }
    }

    i = j
  }

  return string(resultBytes)
}

func main() {
	// result: "RR.L"
	// dominoes := "RR.L"

	// result: "LL.RR.LLRRLL.."
	// dominoes := ".L.R...LR..L.."

	// result: "RRRL"
	// dominoes := "RRRL"

	// result: "LLRRLL"
	// dominoes := ".LR..L"

	// result: "L"
	// dominoes := "L"

	// result: "L.RR.LLLRRLRLRRRRRRRRRLLLLRRLLR"
	dominoes := "L.R...LLRRLRLR....R..RLLLLR..LR"

	// result: ""
	// dominoes := ""

	result := pushDominoes(dominoes)
	fmt.Printf("result = %v\n", result)
}
