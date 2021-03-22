"""
Module Description: 19-删除链表的倒数第 N 个结点
Solution：
Date: 2021-03-11 13:39:13
Author: Wang P
Problem:# 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。 
        #
        #  进阶：你能尝试使用一趟扫描实现吗？

        #  示例 1：
        # 输入：head = [1,2,3,4,5], n = 2
        # 输出：[1,2,3,5]
        #
        #  示例 2：
        # 输入：head = [1], n = 1
        # 输出：[]
        #
        #  示例 3：
        # 输入：head = [1,2], n = 1
        # 输出：[1]
        #
        #  提示：
        #  链表中结点的数目为 sz
        #  1 <= sz <= 30
        #  0 <= Node.val <= 100
        #  1 <= n <= sz
        #
        #  Related Topics 链表 双指针
        #  👍 1264 👎 0
"""
# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        hair = ListNode(next=head)
        slow, fast = hair, head;
        while n > 0:
            fast = fast.next
            n -= 1

        while fast:
            fast = fast.next
            slow = slow.next

        slow.next = slow.next.next
        return hair.next
# leetcode submit region end(Prohibit modification and deletion)
