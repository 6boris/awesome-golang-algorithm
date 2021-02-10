from typing import List


class Solution:
    def printNumbers(self, n: int) -> List[int]:
        ans = []
        for i in range(1, 10 ** n):
            ans.append(i)
        return ans


class Solution2:
    def printNumbers(self, n: int) -> List[int]:
        return list(range(1, 10 ** n))


if __name__ == '__main__':
    print(Solution().printNumbers(3))
    print(Solution2().printNumbers(3))
    print(10 ** 2)
