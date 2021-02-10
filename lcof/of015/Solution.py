class Solution:
    def hammingWeight(self, n: int) -> int:
        ans = 0
        while n:
            ans += n & 1
            n >>= 1
        return ans


class Solution2:
    def hammingWeight(self, n: int) -> int:
        ans = 0
        while n:
            ans += 1
            n &= n - 1
            print(ans,n)
        return ans


if __name__ == '__main__':
    print(Solution().hammingWeight(0b000000000000000000000000000010101))
    print(Solution2().hammingWeight(0b000000000000000000000000000010101))

    print(bin(21))
    print(bin(20))
    print(21 & 20)
    print((21 & 20) >> 1)


    print(bin(20))
    print(bin(19))
    print(20 & 19)
    print((20 & 19) >> 1)