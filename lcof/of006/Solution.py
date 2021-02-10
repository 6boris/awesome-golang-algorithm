from typing import List


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reversePrint(self, head: ListNode) -> List[int]:
        return self.reversePrint(head.next) + [head.val] if head else []


class Solution2:
    def reversePrint(self, head: ListNode) -> List[int]:
        stack = []
        while head:
            stack.append(head.val)
            head = head.next
        stack.reverse()
        return stack



if __name__ == '__main__':
    node = head = ListNode(1)
    node.next = ListNode(3)
    node = node.next
    node.next = ListNode(2)
    print(Solution().reversePrint(head))
    print(Solution2().reversePrint(head))
