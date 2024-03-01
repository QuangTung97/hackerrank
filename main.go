package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}
	if n == 0 {
		return 1.0
	}
	v := myPow(x, n>>1)
	if n&0x1 == 0 {
		return v * v
	}
	return v * v * x
}
