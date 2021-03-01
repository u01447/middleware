"""
Module Description: 基于collections.deque实现队列
                    基于栈实现队列
Problem:
Solution：
Date: 2020/2/15 
Author: Wang P
"""
from collections import deque
from algorithm.data_structure.stack_cases import Stack


class Queue:
    def __init__(self):
        self.items = deque()

    def push(self, val):
        self.items.append(val)

    def pop(self):
        return self.items.popleft()

    def empty(self):
        return len(self.items) == 0


class QueueByStack:
    def __init__(self):
        self.stack_one = Stack()
        self.stack_two = Stack()

    def push(self, val):
        self.stack_one.push(val)

    def pop(self):
        if self.stack_two.empty():
            if self.stack_one.empty():
                return 'nil'
            while not self.stack_one.empty():
                self.stack_two.push(self.stack_one.pop())
        return self.stack_two.pop()

    def peek(self):
        if self.stack_two.empty():
            if self.stack_one.empty():
                return 'nil'
            while not self.stack_one.empty():
                self.stack_two.push(self.stack_one.pop())
        return self.stack_two.top()

    def empty(self):
        return self.stack_two


class QueueByList:
    """Queue represented by a Python list"""
    def __init__(self):
        self.entries = []
        self.length = 0
        self.front = 0

    def __str__(self):
        printed = "<" + str(self.entries)[1:-1] + ">"
        return printed

    def put(self, item):
        """
        Enqueues {@code item}
            @param item
                item to enqueue
        """
        self.entries.append(item)
        self.length = self.length + 1

    def get(self):
        """
        Dequeues {@code item}
            @requirement: |self.length| > 0
            @return dequeued
                item that was dequeued
        """
        self.length = self.length - 1
        dequeued = self.entries[self.front]
        self.entries = self.entries[1:]
        return dequeued

    def rotate(self, rotation):
        """
        Rotates the queue {@code rotation} times
            @param rotation
                number of times to rotate queue
        """
        for i in range(rotation):
            self.put(self.get())

    def get_front(self):
        """
        Enqueues {@code item}
            @return item at front of self.entries
        """
        return self.entries[0]

    def size(self):
        """
        Returns the length of this.entries
        """
        return self.length


if __name__ == "__main__":
    qu = QueueByStack()
    qu.push(1)
    qu.push(2)
    qu.push(3)
    print(qu.pop())
    print(qu.peek())
    print(qu.pop())
    print(qu.pop())
    print(qu.pop())
