"""
Module Description: 获取大量元素中 top-n 大个元素，固定内存
                    合并多个有序链表/list
Problem:
Solution：
Date: 2020/2/16 
Author: Wang P
"""
import heapq


class TopN:
    """
    先放入元素前n个建立小顶堆
    迭代剩余元素：
        如果当前元素小于堆顶元素，跳过该元素
        否则退换堆顶元素为当前元素，并重新调整堆
    """
    def __init__(self, iterable, n):
        self.min_heap = []
        self.capacity = n
        self.iterable = iterable

    def push(self, val):
        if len(self.min_heap) >= self.capacity:
            min_val = self.min_heap[0]
            if val < min_val:  # 可以省略
                pass
            else:
                heapq.heapreplace(self.min_heap, val)  # 返回并且pop堆顶最小值，推入心得val并调整堆
        else:
            heapq.heapreplace(self.min_heap, val)  # 前n个元素直接放入堆中

    def get_top_n(self):
        for val in self.iterable:
            self.push(val)
        return self.min_heap

