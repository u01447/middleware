"""
Module Description: 字符串相关demo
Problem:
Solution：
Date: 2020/2/18 
Author: Wang P
"""


class ReverseString:
    """
    字符数组反转
    """
    def reverse_string(self, x, reverse=False):
        """
        :param x: 字符数组
        :param reverse: 是否之用内置库
        :return:
        """
        if reverse:
            x.reverse()
            return x
        else:
            if not x:
                return
            beg, end = 0, len(x)-1
            while beg < end:
                x[beg], x[end] = x[end], x[beg]
                beg += 1
                end -= 1
            return x


class Palindrome:
    """
    判断字符串是否是回文
    """
    def is_palindrome(self, x):
        if isinstance(x, int) and x < 0:
            return False
        xx = str(x)
        beg, end = 0, len(xx)-1
        while beg < end:
            if xx[beg] == xx[end]:
                beg += 1
                end -= 1
            else:
                return False
        return True


if __name__ == "__main__":
    r = ReverseString()
    print(r.reverse_string(['d', 'h', 's', 'j', 'f'], False))
