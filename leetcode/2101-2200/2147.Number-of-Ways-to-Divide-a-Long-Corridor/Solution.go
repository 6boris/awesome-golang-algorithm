package Solution

const mod2147 = 1000000007

func Solution(corridor string) int {
	totalS := 0 // 记录s的总数
	tmpS := 0   //记录s是否到2个

	bs := []byte(corridor)

	//记录p的个数
	p := 0
	ans := 1
	for _, b := range bs {
		// sppss p spss
		if b == 'P' {

			if tmpS == 2 {
				// 他的前面已经有两个
				p++
			}
			continue
		}

		if b == 'S' {
			totalS++
			tmpS++
			if tmpS == 3 {
				ans = (ans * (1 + p)) % mod2147
				tmpS = 1
			}
			p = 0
		}
	}
	// 奇数没法分
	if totalS&1 == 1 || totalS == 0 {
		return 0
	}
	return ans
}
