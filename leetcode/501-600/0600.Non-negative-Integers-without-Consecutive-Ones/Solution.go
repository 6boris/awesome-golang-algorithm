package Solution

func nBit(n int) ([]int, int) {
	a := make([]int, 32)
	i := 31
	for n > 0 {
		a[i] = n & 1
		n >>= 1
		i--
	}
	return a, i + 1
}

func dfsBit(preBit, index int, less bool, array []int) int {
	if index == 32 {
		return 1
	}
	if preBit == 1 {
		if array[index] == 1 {
			return dfsBit(0, index+1, true, array)
		}
		return dfsBit(0, index+1, less, array)
	}
	if array[index] == 1 {
		return dfsBit(1, index+1, less, array) + dfsBit(0, index+1, true, array)
	}
	a := dfsBit(0, index+1, less, array)
	if less {
		a += dfsBit(1, index+1, less, array)
	}
	return a
}

func Solution(n int) int {
	if n <= 2 {
		return n + 1
	}
	bits, startIndex := nBit(n)
	cnt := 0
	zero, one := 1, 1
	for i := 30; i > startIndex; i-- {
		tmpZero := zero + one
		tmpOne := zero
		zero, one = tmpZero, tmpOne
	}
	cnt += zero + one

	return cnt + dfsBit(1, startIndex+1, false, bits)
}
