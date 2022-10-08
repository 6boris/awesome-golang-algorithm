package Solution

import "container/heap"

/*
[天数][股票状态]
股票状态: 奇数表示第 k 次交易持有/买入, 偶数表示第 k 次交易不持有/卖出, 0 表示没有操作
*/

func maxProfit_1(k int, prices []int) int {
	n, dp := len(prices), [][]int{}
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, k*2+1))
	}
	for i := 1; i < k*2; i += 2 {
		dp[0][i] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < k*2; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[n-1][k*2]
}

func maxProfit_2(k int, prices []int) int {
	n := len(prices)
	dp := make([]int, k*2+1)
	for i := 1; i < k*2; i += 2 {
		dp[i] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < k*2; j += 2 {
			dp[j+1] = max(dp[j+1], dp[j]-prices[i])
			dp[j+2] = max(dp[j+2], dp[j+1]+prices[i])
		}
	}
	return dp[k+2]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// heap + stack
type stock []int

func (s *stock) Len() int {
	return len(*s)
}

func (s *stock) Less(i, j int) bool {
	return (*s)[i] > (*s)[j]
}

func (s *stock) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *stock) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *stock) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

type stockPair struct {
	start, end int
}

func maxProfit_3(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}

	// 开始忘记合并1 ,3, 2, 5这种情况，如何合并参考了 https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/discuss/54145/O(n)-time-8ms-Accepted-Solution-with-Detailed-Explanation-(C%2B%2B)
	//diff := 0
	//differences := &stock{}
	//
	//// 3, 2, 3, 6, 5, 0, 3
	//
	//min := prices[0]
	//for idx := 1; idx < len(prices); idx++ {
	//	if prices[idx] < prices[idx-1] {
	//		if diff != 0 && prices[idx] < min {
	//			heap.Push(differences, diff)
	//			diff = 0
	//		}
	//		fmt.Println()
	//		continue
	//	}
	//
	//	diff += prices[idx] - prices[idx-1]
	//	fmt.Printf("diff += prices[%d]-prices[%d]=%d\n", idx, idx-1, diff)
	//}
	//if diff != 0 {
	//	heap.Push(differences, diff)
	//}
	//loop := 0
	//ans := 0
	//for differences.Len() > 0 && loop < k {
	//	pop := heap.Pop(differences).(int)
	//	ans += pop
	//	loop++
	//}

	s, e := 0, -1
	pairs := make([]stockPair, 0)
	profit := &stock{}
	for {
		for s = e + 1; s+1 < len(prices) && prices[s] >= prices[s+1]; s++ {
		}
		for e = s; e+1 < len(prices) && prices[e] <= prices[e+1]; e++ {
		}

		if s == e {
			break
		}

		/*
			情况1：价格[v1]<=价格[v2]，价格[p1]<=价格[p2]。
			在这种情况下，如果我们能进行一次交易，我们将使用（v1，p2）。
			如果我们可以进行两次交易，我们将使用（v1，p1）和（v2，p2）。
			等价地，我们可以把（v1，p2）视为一个交易机会，把（v2，p1）视为另一个交易机会。
			关键的想法是，这两个原始谷峰对提供了两个交易机会：（v1，p2）和（v2，p1）。

			情况2：价格[v1]>=价格[v2]或价格[p1]>=价格[p2]。
			在这种情况下，如果我们可以进行一次交易，我们将使用（v1，p1）或（v2，p2）。
			如果我们可以进行两次交易，我们将同时使用（v1，p1）和（v2，p2）。
			也就是说，这两个谷峰对提供了两个交易机会：（v1，p1）和（v2，p2）。
		*/

		for l := len(pairs); l > 0 && prices[s] <= prices[pairs[l-1].start]; l-- {
			heap.Push(profit, prices[pairs[l-1].end]-prices[pairs[l-1].start])
			pairs = pairs[:l-1]
		}

		for l := len(pairs); l > 0 && prices[e] >= prices[pairs[l-1].end]; l-- {
			heap.Push(profit, prices[pairs[l-1].end]-prices[s])
			s = pairs[l-1].start
			pairs = pairs[:l-1]
		}

		pairs = append(pairs, stockPair{s, e})
	}

	for l := len(pairs); l > 0; l-- {
		heap.Push(profit, prices[pairs[l-1].end]-prices[pairs[l-1].start])
		pairs = pairs[:l-1]
	}

	ans := 0
	loop := 0
	for profit.Len() > 0 && loop < k {
		top := heap.Pop(profit).(int)
		ans += top
		loop++
	}
	return ans
}
