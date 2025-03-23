package Solution

type Bank struct {
	b  []int64
	ok int
}

func Constructor(balance []int64) Bank {
	// 0, 1, 2, 3, 4, 5
	b := make([]int64, len(balance)+1)
	for i := 1; i <= len(balance); i++ {
		b[i] = balance[i-1]
	}
	return Bank{b: b, ok: len(balance)}
}

func (this *Bank) validateAccount(a int) bool {
	return a >= 1 && a <= this.ok
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.validateAccount(account1) || !this.validateAccount(account2) {
		return false
	}
	have := this.b[account1]
	if have >= money {
		this.b[account1] -= money
		this.b[account2] += money
		return true
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.validateAccount(account) {
		return false
	}
	this.b[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	// 0 ,1, 2, 3
	if !this.validateAccount(account) {
		return false
	}
	have := this.b[account]
	if have >= money {
		this.b[account] -= money
		return true
	}
	return false
}

type action struct {
	a        string
	from, to int
	money    int64
}

func Solution(balance []int64, actions []action) []bool {
	c := Constructor(balance)
	ans := make([]bool, len(actions))
	for i, ac := range actions {
		if ac.a == "t" {
			ans[i] = c.Transfer(ac.from, ac.to, ac.money)
			continue
		}
		if ac.a == "d" {
			ans[i] = c.Deposit(ac.from, ac.money)
			continue
		}
		ans[i] = c.Withdraw(ac.from, ac.money)
	}
	return ans
}
