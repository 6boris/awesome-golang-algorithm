# 深度搜索
class Solution:
    def movingCount(self, m: int, n: int, k: int) -> int:
        def bitSum(x):
            s = 0
            while x != 0:
                s += x % 10
                x //= 10
            return s

        def dfs(i, j, si, sj):
            if i >= m or j >= n or k < si + sj or (i, j) in visited: return 0
            visited.add((i, j))
            return 1 + dfs(i + 1, j, bitSum(i + 1), sj) + dfs(i, j + 1, si, bitSum(j + 1))

        visited = set()
        return dfs(0, 0, 0, 0)


class Solution2:
    def movingCount(self, m: int, n: int, k: int) -> int:
        def bitSum(x):
            s = 0
            while x != 0:
                s += x % 10
                x //= 10
            return s

        queue, visited = [(0, 0, 0, 0)], set()

        while queue:
            i, j, si, sj = queue.pop()
            if i >= m or j >= n or k < si + sj or (i, j) in visited: continue
            visited.add((i, j))
            queue.append((i + 1, j, bitSum(i + 1), sj))
            queue.append((i, j + 1, si, bitSum(j + 1)))

        return len(visited)


if __name__ == '__main__':
    print(Solution().movingCount(2, 3, 1))
    print(Solution().movingCount(38, 15, 9))

    print(Solution2().movingCount(2, 3, 1))
    print(Solution2().movingCount(38, 15, 9))
