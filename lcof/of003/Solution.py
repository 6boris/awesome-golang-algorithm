from typing import List


class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        dic = set()
        for v in nums:
            if v in dic: return v
            dic.add(v)
        return -1


class Solution2:
    def findRepeatNumber(self, nums: List[int]) -> int:
        nums.sort()
        for idx in range(len(nums) - 1):
            if nums[idx] == nums[idx + 1]:
                return nums[idx]


class Solution3:
    def findRepeatNumber(self, nums: List[int]) -> int:
        for idx, v in enumerate(nums):
            if idx != v and nums[v] == v: return v
            nums[v], nums[idx] = nums[idx], nums[v]


if __name__ == '__main__':
    print(Solution().findRepeatNumber([2, 3, 1, 0, 2, 5, 3]))
    print(Solution2().findRepeatNumber([2, 3, 1, 0, 2, 5, 3]))
    print(Solution3().findRepeatNumber([2, 3, 1, 0, 2, 5, 3]))
