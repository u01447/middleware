"""
Module Description: 82-删除排序链表中的重复元素 II
Solution：
Date: 2021-03-11 13:50:50
Author: Wang P
Problem:# 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
        #  示例 1:
        #  输入: 1->2->3->3->4->4->5
        # 输出: 1->2->5
        #  示例 2:
        #  输入: 1->1->1->2->3
        # 输出: 2->3
        #  Related Topics 链表
        #  👍 469 👎 0
"""
# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        hair, fast = ListNode(next=head), head
        slow = hair
        while fast:
            while fast.next and fast.val == fast.next.val:
                fast = fast.next
            if slow.next == fast:
                slow = slow.next
            else:
                slow.next = fast.next
            fast = fast.next

        return hair.next
# leetcode submit region end(Prohibit modification and deletion)
