"""
Module Description: 24-两两交换链表中的节点
Solution：
Date: 2021-03-18 20:25:15
Author: Wang P
Problem:  给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
          你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

          示例 1：
         输入：head = [1,2,3,4]
         输出：[2,1,4,3]

          示例 2：
         输入：head = []
         输出：[]

          示例 3：
         输入：head = [1]
         输出：[1]

          提示：
          链表中节点的数目在范围 [0, 100] 内
          0 <= Node.val <= 100
          进阶：你能在不修改链表节点值的情况下解决这个问题吗?（也就是说，仅修改节点本身。）
          Related Topics 递归 链表
          👍 859 👎 0
"""
# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:

    def swapPairs(self, head: ListNode) -> ListNode:
        tail = pre = hair = ListNode(next=head)
        while head:
            # tail = pre
            for i in range(0, 2):
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
