package Solution

type Cashier struct {
	n, cur int

	discount float64

	pp map[int]int
}

func Constructor1357(n int, discount int, products []int, prices []int) Cashier {
	pp := map[int]int{}
	for i := range products {
		pp[products[i]] = prices[i]
	}
	return Cashier{
		n: n, cur: 0, discount: float64(100-discount) / 100.0,
		pp: pp,
	}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	cost := 0
	for i := range product {
		cost += this.pp[product[i]] * amount[i]
	}
	fcost := float64(cost)
	if this.cur == this.n-1 {
		fcost *= this.discount
	}
	this.cur = (this.cur + 1) % this.n
	return fcost
}

type input struct {
	product, amount []int
}

func Solution(n, discount int, products, prices []int, inputs []input) []float64 {
	c := Constructor1357(n, discount, products, prices)
	ans := make([]float64, len(inputs))
	for i := range inputs {
		ans[i] = c.GetBill(inputs[i].product, inputs[i].amount)
	}
	return ans
}
