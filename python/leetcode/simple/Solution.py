# -*- coding: utf-8 -*- 
"""
Module Description:
Time : 2020/8/6 15:57 
Author : Wang P
File : Solution.py
version: 
"""
from typing import List


class Solution:

    def twoSum(self, nums: List[int], target: int, flag=None) -> List[int]:
        """
        给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
        你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
        示例:
            给定 nums = [2, 7, 11, 15], target = 9
            因为 nums[0] + nums[1] = 2 + 7 = 9
            所以返回 [0, 1]
        Related Topics 数组 哈希表
        👍 8826 👎 0
        :param nums:
        :param target:
        :param flag:
        :return:
        """
        if not nums or len(nums) <= 1:
            raise Exception("Illegal nums")

        if not flag:
            num_map = {nums[0]: 0}
            temp = nums[0]
            for i in range(1, len(nums)):
                temp = target - nums[i]
                if temp in num_map:
                    return [num_map[temp], i]
                else:
                    num_map.update({int(nums[i]): i})
        else:
            for i in range(0, len(nums)-1):
                for j in range(i+1, len(nums)):
                    if nums[i] + nums[j] == target:
                        return [nums[i], nums[j]]
        raise Exception("No two sum solution")

    def reverse(self, x: int) -> int:
        """
        给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

         示例 1:

         输入: 123
         输出: 321

         示例 2:

         输入: -123
         输出: -321
         示例 3:

         输入: 120
        输出: 21
         注意:

        假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31, 2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
        Related Topics 数学python 负数
        :param x:
        :return:
        """
        from copy import deepcopy
        res = 0
        y = deepcopy(x)
        x = abs(x)
        while x != 0:
            res = res * 10 + x % 10
            x = x // 10
        res = res if y >= 0 else -res
        return 0 if -(1 << 31) > res or ((1 << 31) - 1) < res else res

    def isPalindrome(self, x: int) -> bool:
        """
        判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
         示例 1:
         输入: 121
         输出: true
         示例 2:
         输入: -121
         输出: false
         解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
         示例 3:
         输入: 10
         输出: false
         解释: 从右向左读, 为 01 。因此它不是一个回文数。

         进阶:
         你能不将整数转为字符串来解决这个问题吗？
         Related Topics 数学
         👍 1174 👎 0
        :param x:
        :return:
        """
        if x == 0:
            return True
        if x < 0 or x % 10 == 0:
            return False

        r = 0
        while x > r:
            r = r * 10 + x % 10
            x = x // 10
        return x == r or r // 10 == x

