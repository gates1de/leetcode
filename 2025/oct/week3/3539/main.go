package main

import (
	"fmt"
	"math/bits"
)

const modulo = int64(1e9 + 7)

func magicalSum(m int, k int, nums []int) int {
	fac := make([]int64, m + 1)
	fac[0] = 1
	for i := 1; i <= m; i++ {
		fac[i] = fac[i - 1] * int64(i) % modulo
	}

	ifac := make([]int64, m + 1)
	ifac[0] = 1
	ifac[1] = 1
	for i := 2; i <= m; i++ {
		ifac[i] = multiple(int64(i), modulo - 2, modulo)
	}
	for i := 2; i <= m; i++ {
		ifac[i] = ifac[i - 1] * ifac[i] % modulo
	}

	powers := make([][]int64, len(nums))
	for i := range nums {
		powers[i] = make([]int64, m + 1)
		powers[i][0] = 1

		for j := 1; j <= m; j++ {
			powers[i][j] = powers[i][j - 1] * int64(nums[i]) % modulo
		}
	}

	f := make([][][][]int64, len(nums))
	for i := range nums {
			f[i] = make([][][]int64, m+1)
			for j := 0; j <= m; j++ {
					f[i][j] = make([][]int64, m * 2 + 1)
					for p := 0; p <= m * 2; p++ {
							f[i][j][p] = make([]int64, k + 1)
					}
			}
	}

	for j := 0; j <= m; j++ {
		f[0][j][j][0] = powers[0][j] * ifac[j] % modulo
	}
	for i := 0; i + 1 < len(nums); i++ {
		for j := 0; j <= m; j++ {
			for p := 0; p <= m * 2; p++ {
				for q := 0; q <= k; q++ {
					q2 := p % 2 + q
					if q2 > k {
						break
					}

					for r := 0; r + j <= m; r++ {
						p2 := p / 2 + r
						f[i + 1][j + r][p2][q2] += f[i][j][p][q] * powers[i + 1][r] % modulo * ifac[r] % modulo
						f[i + 1][j + r][p2][q2] %= modulo
					}
				}
			}
		}
	}

	result := int64(0)
	for p := 0; p <= m * 2; p++ {
		for q := 0; q <= k; q++ {
			if bits.OnesCount(uint(p)) + q == k {
				result = (result + f[len(nums) - 1][m][p][q] * fac[m] % modulo) % modulo
			}
		}
	}

	return int(result)
}

func multiple(x, y, mod int64) int64 {
	result := int64(1)
	tmp := x % mod

	for y > 0 {
		if y & 1 == 1 {
			result = result * tmp % mod
		}

		y >>= 1
		tmp = tmp * tmp % mod
	}

	return result
}

func main() {
	// result: 991600007
	// m := int(5)
	// k := int(5)
	// nums := []int{1,10,100,10000,1000000}

	// result: 170
	// m := int(2)
	// k := int(2)
	// nums := []int{5,4,3,2,1}

	// result: 28
	m := int(1)
	k := int(1)
	nums := []int{28}

	// result: 
	// m := int(0)
	// k := int(0)
	// nums := []int{}

	result := magicalSum(m, k, nums)
	fmt.Printf("result = %v\n", result)
}

