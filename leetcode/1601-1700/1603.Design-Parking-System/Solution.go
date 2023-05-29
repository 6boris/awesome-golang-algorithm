package Solution

type ParkingSystem struct {
	solt [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{solt: [3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.solt[carType-1] > 0 {
		this.solt[carType-1]--
		return true
	}
	return false
}

func Solution(big, medium, small int, ops []int) []bool {
	c := Constructor(big, medium, small)
	ans := make([]bool, len(ops))
	for i, op := range ops {
		ans[i] = c.AddCar(op)
	}
	return ans
}
