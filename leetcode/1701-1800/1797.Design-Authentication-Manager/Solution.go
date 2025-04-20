package Solution

type AuthenticationManager struct {
	live  int64
	token map[string]int64
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{live: int64(timeToLive), token: make(map[string]int64)}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.token[tokenId] = int64(currentTime) + this.live
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	expired, ok := this.token[tokenId]
	if !ok {
		return
	}
	ic := int64(currentTime)
	if expired > ic {
		this.token[tokenId] = ic + this.live
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	ans := 0
	ic := int64(currentTime)
	for _, c := range this.token {
		if c > ic {
			ans++
		}
	}
	return ans
}

type input struct {
	op    byte
	token string
	v     int
}

func Solution(n int, inputs []input) []int {
	c := Constructor(n)
	ans := make([]int, 0)
	for _, op := range inputs {
		if op.op == 'g' {
			c.Generate(op.token, op.v)
			continue
		}
		if op.op == 'r' {
			c.Renew(op.token, op.v)
			continue
		}
		ans = append(ans, c.CountUnexpiredTokens(op.v))
	}
	return ans
}
