"""
Module Description: 基队列实现栈
Problem:
Solution：基于collections.deque实现栈
          基于 2 个队列实现
          基于 1 个队列实现
Date: 2020/2/15 
Author: Wang P
"""
from collections import deque
from queue import Queue


class Stack:
    def __init__(self):
        self.items = deque()

    def push(self, val):
        self.items.append(val)

    def top(self):
        return self.items[-1]

    def pop(self):
        return self.items.pop()

    def empty(self):
        return len(self.items) == 0


class MyStack:
    """
    使用 2 个队列实现
    """
    def __init__(self):
        """
        Initialize your data structure here.
        """
        # q1 作为进栈出栈，q2 作为中转站
        self.q1 = Queue()
        self.q2 = Queue()

    def push(self, x):
        """
        Push element x onto stack.
        ：type x： int
        ：rtype： void
        """
        self.q1.put(x)

    def pop(self):
        """
        Removes the element on top of the stack and returns that element.
        ：rtype： int
        """
        while self.q1.qsize() > 1:
            self.q2.put(self.q1.get())  # 将 q1 中除尾元素外的所有元素转到 q2 中
        if self.q1.qsize() == 1:
            res = self.q1.get()  # 弹出 q1 的最后一个元素
            self.q1, self.q2 = self.q2, self.q1  # 交换 q1,q2
            return res

    def top(self):
        """
        Get the top element.
        ：rtype： int
        """
        while self.q1.qsize() > 1:
            self.q2.put(self.q1.get())  # 将 q1 中除尾元素外的所有元素转到 q2 中
        if self.q1.qsize() == 1:
            res = self.q1.get()  # 弹出 q1 的最后一个元素
            self.q2.put(res)  # 与 pop 唯一不同的是需要将 q1 最后一个元素保存到 q2 中
            self.q1, self.q2 = self.q2, self.q1  # 交换 q1,q2
            return res

    def empty(self):
        """
        Returns whether the stack is empty.
        ：rtype： bool
        """
        return not bool(self.q1.qsize() + self.q2.qsize())  # 为空返回 True，不为空返回 False


class MyStack2(object):
    """
    使用 1 个队列实现
    """
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.sq1 = Queue()

    def push(self, x):
        """
        Push element x onto stack.
        ：type x： int
        ：rtype： void
        """
        self.sq1.put(x)

    def pop(self):
        """
        Removes the element on top of the stack and returns that element.
        ：rtype： int
        """
        count = self.sq1.qsize()
        if count == 0:
            return False
        while count > 1:
            x = self.sq1.get()
            self.sq1.put(x)
            count -= 1
        return self.sq1.get()

    def top(self):
        """
        Get the top element.
        ：rtype： int
        """
        count = self.sq1.qsize()
        if count == 0:
            return False
        while count:
            x = self.sq1.get()
            self.sq1.put(x)
            count -= 1
        return x

    def empty(self):
        """
        Returns whether the stack is empty.
        ：rtype： bool
        """
        return self.sq1.empty()


if __name__ == '__main__':
    obj = MyStack2()
    obj.push(1)
    obj.push(3)
    print(obj.top())
    print(obj.top())
    print(obj.top())

