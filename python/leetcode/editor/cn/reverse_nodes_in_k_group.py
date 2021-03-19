"""
Module Description: 25-K 个一组翻转链表
Solution：
Date: 2021-03-06 18:15:19
Author: Wang P
Problem:# 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。 
        #  k 是一个正整数，它的值小于或等于链表的长度。
        #  如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
        #  进阶：
        #  你可以设计一个只使用常数额外空间的算法来解决此问题吗？
        #  你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
        #
        #  示例 1：
        # 输入：head = [1,2,3,4,5], k = 2
        # 输出：[2,1,4,3,5]
        #
        #  示例 2：
        # 输入：head = [1,2,3,4,5], k = 3
        # 输出：[3,2,1,4,5]
        #
        #  示例 3：
        # 输入：head = [1,2,3,4,5], k = 1
        # 输出：[1,2,3,4,5]
        #
        #  示例 4：
        # 输入：head = [1], k = 1
        # 输出：[1]
        #
        #  提示：
        #  列表中节点的数量在范围 sz 内
        #  1 <= sz <= 5000
        #  0 <= Node.val <= 1000
        #  1 <= k <= sz
        #
        #  Related Topics 链表
        #  👍 977 👎 0
"""

# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        tail = pre = hair = ListNode(next=head)
        while head:
            # tail = pre
            for i in range(0, k):
                tail = tail.next
                if not tail:
                    return hair.next

            head, tail = self.reverse(head, tail)
            pre.next = head
            pre = tail
            head = pre.next

        return hair.next

    def reverse(self, head, tail: ListNode) -> (ListNode, ListNode):
        pre, cur, nex = tail.next, head, None
        while pre != tail:
            nex = cur.next
            cur.next = pre
            pre = cur
            cur = nex

        return tail, head

# leetcode submit region end(Prohibit modification and deletion)
