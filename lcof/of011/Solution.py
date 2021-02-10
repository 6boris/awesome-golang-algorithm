from typing import List

# 4 5 6 2 3
# 4 5 1 2 3

class Solution:
    def minArray(self, numbers: List[int]) -> int:
        left, right = 0, len(numbers) - 1
        while left < right:
            mid = (left + right) // 2
            if numbers[mid] > numbers[right]:
                left = mid + 1
            elif numbers[mid] < numbers[right]:
                right = mid
            else:
                right -= 1

        return numbers[left]


if __name__ == '__main__':
    print(Solution().minArray([3, 4, 5, 1, 2]))
    print(Solution().minArray([2, 2, 2, 0, 1]))
