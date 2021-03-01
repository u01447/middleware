"""
Module Description: 冒泡排序
Problem:
Solution：
Date: 2020/3/23 
Author: Wang P
"""
import random


class BubbleSort:

    def bubble_sort(self, arr):
        """
        冒泡排序
        :param arr:
        :return:
        """
        n = len(arr)
        for i in range(n):
            for j in range(n-i-1):
                if arr[j] > arr[j+1]:
                    arr[j], arr[j+1] = arr[j+1], arr[j]


if __name__ == "__main__":
    list_a = [random.randint(0, 8) for _ in range(10)]
    print(list_a)
    bs = BubbleSort()
    bs.bubble_sort(list_a)
    print(list_a)
