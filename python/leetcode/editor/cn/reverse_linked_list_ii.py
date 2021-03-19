"""
Module Description: 92-反转链表 II
Solution：
Date: 2021-03-15 19:06:58
Author: Wang P
Problem:# 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。 
        #
        #  说明:
        # 1 ≤ m ≤ n ≤ 链表长度。
        #
        #  示例:
        #
        #  输入: 1->2->3->4->5->NULL, m = 2, n = 4
        # 输出: 1->4->3->2->5->NULL
        #  Related Topics 链表
        #  👍 709 👎 0
"""
# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.

# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    def reverseBetween(self, head: ListNode, left: int, right: int) -> ListNode:
        cur = hair = ListNode(next=head)
        n = right - left + 1
        while left > 1:
            cur = cur.next
            left -= 1

        cur.next = self.reverse(cur.next, n)
        return hair.next

    def reverse(self, head: ListNode, n: int) -> ListNode:
        if not head:
            return None

        pre, cur, nex = ListNode(), head, None
        while n > 0:
            nex = cur.next
            cur.next = pre.next
            pre.next = cur
            cur = nex
            n -= 1
        head.next = cur
        return pre.next
# leetcode submit region end(Prohibit modification and deletion)