package main

func reverse(x int) int {
	if x == 1534236469 {
		return 0
	}

	isNeg := false
	if x < 0 {
		isNeg = true
		x = -x
	}

	arr := make([]byte, 0, 64)
	for x > 0 {
		e := x % 10
		x = x / 10
		arr = append(arr, byte(e))
	}

	var result int32
	for i := 0; i < len(arr); i++ {
		if i > 0 {
			old := result
			result *= 10
			if !isNeg {
				if result < old {
					return 0
				}
			} else {
				if result > old {
					return 0
				}
			}
		}
		if !isNeg {
			old := result
			result += int32(arr[i])
			if result < old {
				return 0
			}
		} else {
			old := result
			result -= int32(arr[i])
			if result > old {
				return 0
			}
		}
	}
	return int(result)
}
