package main

import (
	"fmt"
)

const modulo = int64(1e9 + 7)

func zigZagArrays(n int, l int, r int) int {
	m := r - l + 1

	buildMatrixA := func() []int64 {
		mat := make([]int64, m * m)
		for i := range m {
			for j := range m {
				if i < j {
					mat[i * m + j] = int64(i)
				} else {
					mat[i * m + j] = int64(j)
				}
			}
		}
		return mat
	}

	buildMatrixB := func() []int64 {
		mat := make([]int64, m * m)
		for i := range m {
			for j := range m {
				if i > j {
					mat[i * m + j] = int64(m - 1 - i)
				} else {
					mat[i * m + j] = int64(m - 1 - j)
				}
			}
		}
		return mat
	}

	multiplyMatrix := func(a []int64, b []int64) []int64 {
		result := make([]int64, m * m)
		for i := range m {
			row := i * m
			for k := range m {
				aik := a[row + k]
				if aik == 0 {
					continue
				}

				col := k * m
				for j := range m {
					result[row + j] = (result[row + j] + aik * b[col + j]) % modulo
				}
			}
		}
		return result
	}

	applyMatrix := func(mat []int64, vec []int64) []int64 {
		result := make([]int64, m)
		for i := range m {
			sum := int64(0)
			row := i * m
			for j := range m {
				sum = (sum + mat[row + j] * vec[j]) % modulo
			}
			result[i] = sum
		}
		return result
	}

	powApply := func(mat []int64, vec []int64, exp int) []int64 {
		base := mat
		for exp > 0 {
			if exp&1 == 1 {
				vec = applyMatrix(base, vec)
			}
			exp >>= 1
			if exp > 0 {
				base = multiplyMatrix(base, base)
			}
		}
		return vec
	}

	if n == 1 {
		return m
	}

	matrixA := buildMatrixA()
	matrixB := buildMatrixB()

	result := int64(0)
	if n%2 == 1 {
		vec := make([]int64, m)
		for i := range m {
			vec[i] = 1
		}
		k := (n - 1) / 2
		up := powApply(matrixA, append([]int64(nil), vec...), k)
		down := powApply(matrixB, vec, k)
		for i := range m {
			result = (result + up[i] + down[i]) % modulo
		}
		return int(result)
	}

	upVec := make([]int64, m)
	downVec := make([]int64, m)
	for i := range m {
		upVec[i] = int64(i)
		downVec[i] = int64(m - 1 - i)
	}

	k := (n - 2) / 2
	up := powApply(matrixA, upVec, k)
	down := powApply(matrixB, downVec, k)
	for i := range m {
		result = (result + up[i] + down[i]) % modulo
	}

	return int(result)
}

func main() {
	// result: 2
	// n := int(3)
	// l := int(4)
	// r := int(5)

	// result: 10
	n := int(3)
	l := int(1)
	r := int(3)

	// result: 
	// n := int(0)
	// l := int(0)
	// r := int(0)

	result := zigZagArrays(n, l, r)
	fmt.Printf("result = %v\n", result)
}
