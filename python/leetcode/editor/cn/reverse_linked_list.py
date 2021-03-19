"""
Module Description: 206-反转链表
Solution：
Date: 2021-03-11 17:45:14
Author: Wang P
Problem:# 反转一个单链表。 
        #
        #  示例:
        #
        #  输入: 1->2->3->4->5->NULL
        # 输出: 5->4->3->2->1->NULL
        #
        #  进阶:
        # 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
        #  Related Topics 链表
        #  👍 1568 👎 0

"""
# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        pre, cur, nex = None, head, None
        while cur:
            nex = cur.next
            cur.next = pre
            pre = cur
            cur = nex
        return pre
# leetcode submit region end(Prohibit modification and deletion)
