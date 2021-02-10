class Solution:
    def replaceSpace(self, s: str) -> str:
        ans = []
        for v in s:
            if v == ' ':
                ans.append('%20')
            else:
                ans.append(v)
        return ''.join(ans)


class Solution2:
    def replaceSpace(self, s: str) -> str:
        return s.replace(' ', '%20')

if __name__ == '__main__':
    print(Solution().replaceSpace("We are happy."))
    print(Solution2().replaceSpace("We are happy."))
