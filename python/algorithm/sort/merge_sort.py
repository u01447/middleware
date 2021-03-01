"""
Module Description: 归并排序算法-分治法
Problem:
Solution：
Date: 2020/2/15 
Author: Wang P
"""


class MergeSort:

    def merge_sort(self, arr):
        """
        分治法分三步走，要注意出口
        :param arr:
        :return:
        """
        # 递归出口
        if len(arr) <= 1:
            return arr
        else:
            mid = int(len(arr)/2)
            left_half = self.merge_sort(arr[:mid])
            right_half = self.merge_sort(arr[mid:])
            return self.merge_sorted_list(left_half, right_half)

    def merge_sorted_list(self, list_a, list_b):
        """
        合并两个有序序列
        :param list_a:
        :param list_b:
        :return:
        """
        length_a = len(list_a)
        length_b = len(list_b)
        a = b = 0
        new_list = []
        while a < length_a and b < length_b:
            if list_a[a] < list_b[b]:
                new_list.append(list_a[a])
                a += 1
            else:
                new_list.append(list_b[b])
                b += 1
        if a < length_a:
            new_list.extend(list_a[a:])
        if b < length_b:
            new_list.extend(list_b[b:])
        return new_list


if __name__ == "__main__":
    import random
    ms = MergeSort()
    ll = list(range(10))
    random.shuffle(ll)
    print(ll)
    ll = ms.merge_sort(ll)
    print(ll)
