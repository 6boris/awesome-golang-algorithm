class CQueue:

    def __init__(self):

        self.in_stack, self.out_stack = [], []

    def appendTail(self, value: int) -> None:
        # 栈为顶队列尾部
        self.in_stack.append(value)

    def deleteHead(self) -> int:
        # 出栈有元素直接弹出
        if self.out_stack: return self.out_stack.pop()
        # 2个栈都没有元素
        if not self.in_stack: return -1
        # 入栈有元素，需要将入栈的元素反过来再删除头部
        while self.in_stack: self.out_stack.append(self.in_stack.pop())
        return self.out_stack.pop()


if __name__ == '__main__':
    obj = CQueue()
    obj.appendTail(1)
    param_2 = obj.deleteHead()
