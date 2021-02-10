class Solution:
    def fib(self, n: int) -> int:
        if n < 2: return n
        return self.fib(n - 1) + self.fib(n - 2) % 1000000007


class Solution2:
    def fib(self, n: int) -> int:
        records = [-1 for i in range(n + 1)]  # 记录计算的值
        if n == 0: return 0
        if n == 1: return 1
        if records[n] == -1:  # 表明这个值没有算过
            records[n] = self.fib(n - 1) + self.fib(n - 2)
        return records[n]


class Solution3:
    def fib(self, n: int) -> int:
        dp = {}
        dp[0] = 0
        dp[1] = 1
        if n >= 2:
            for i in range(2, n + 1):
                dp[i] = dp[i - 1] + dp[i - 2]
        return dp[n] % 1000000007


if __name__ == '__main__':
    print(Solution().fib(5))
    print(Solution2().fib(5))
    print(Solution3().fib(5))
