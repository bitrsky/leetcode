package lib_math

// 快速幂
func fastPow(x float64, n int) float64 {
	if n <= 0 {
		x = 1 / x
		n = -1 * n
	}
	ans := 1.0
	for n > 0 {
		if n&1 != 0 {
			ans = ans * x
		}
		x *= x
		n >>= 1
	}
	return ans
}
