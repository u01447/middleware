"""
Module Description: 快速排序算法-分治法
Problem:
Solution：Partition:选择基准分割数组为两个字数组，小于基准和大于基准
         对两个字数组分别快排
         合并结果
Date: 2020/2/15 
Author: Wang P
"""
import random
import pysnooper


class QuickSort:

    @pysnooper.snoop()
    def quick_sort(self, arr):
        if len(arr) < 2:
            return arr
        pivot_index = 0  # 第一个数作为pivot
        pivot = arr[pivot_index]
        less_part = [i for i in arr[pivot_index+1:] if i < pivot]
        great_part = [i for i in arr[pivot_index+1:] if i >= pivot]
        return self.quick_sort(less_part) + [pivot] + self.quick_sort(great_part)


if __name__ == "__main__":
    qs = QuickSort()
    ll = list(range(10))
    random.shuffle(ll)
    print(ll)
    ll = qs.quick_sort(ll)
    print(ll)
