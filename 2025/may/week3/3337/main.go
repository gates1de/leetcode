package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

type Mat struct {
	Alphabets [26][26]int
}

func Constructor() *Mat {
	return &Mat{}
}

func NewMatCopy(from *Mat) *Mat {
	m := &Mat{}
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
          m.Alphabets[i][j] = from.Alphabets[i][j]
      }
  }
  return m
}

func (this *Mat) Mul(other *Mat) *Mat {
  result := Constructor()
  for i := 0; i < 26; i++ {
      for j := 0; j < 26; j++ {
          for k := 0; k < 26; k++ {
              result.Alphabets[i][j] = (result.Alphabets[i][j] + this.Alphabets[i][k] * other.Alphabets[k][j]) % modulo
          }
      }
  }
  return result
}

func InitMat() *Mat {
    m := Constructor()
    for i := 0; i < 26; i++ {
        m.Alphabets[i][i] = 1
    }
    return m
}

func quickMul(x *Mat, y int) *Mat {
  result := InitMat()
  copy := NewMatCopy(x)
  for y > 0 {
		if y & 1 == 1 {
			result = result.Mul(copy)
		}

		copy = copy.Mul(copy)
		y >>= 1
  }

  return result
}

func lengthAfterTransformations(s string, t int, nums []int) int {
	T := Constructor()
	for i := 0; i < 26; i++ {
		for j := 1; j <= nums[i]; j++ {
			T.Alphabets[(i + j) % 26][i] = 1
		}
	}

	mul := quickMul(T, t)
	m := make([]int, 26)
	for _, char := range s {
		m[char - 'a']++
	}

	result := int(0)
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			result = (result + mul.Alphabets[i][j] * m[j]) % modulo
		}
	}

	return result
}

func main() {
	// result: 7
	// s := "abcyy"
	// t := int(2)
	// nums := []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2}

	// result: 8
	s := "azbk"
	t := int(1)
	nums := []int{2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2}

	// result: 
	// s := ""
	// t := int(0)
	// nums := []int{}

	result := lengthAfterTransformations(s, t, nums)
	fmt.Printf("result = %v\n", result)
}

