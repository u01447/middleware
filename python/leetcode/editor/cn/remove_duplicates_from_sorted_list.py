"""
Module Description: 83-删除排序链表中的重复元素
Solution：
Date: 2021-03-11 13:50:26
Author: Wang P
Problem:# 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。 
        #
        #  示例 1:
        #
        #  输入: 1->1->2
        # 输出: 1->2
        #
        #
        #  示例 2:
        #
        #  输入: 1->1->2->3->3
        # 输出: 1->2->3
        #  Related Topics 链表
        #  👍 488 👎 0
"""
# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        if not head:
            return None
        fast, slow = head.next, head

        while fast:
            if fast.val == slow.val:
                slow.next = fast.next
                fast = slow.next
            else:
                fast = fast.next
                slow = slow.next

        return head
# leetcode submit region end(Prohibit modification and deletion)
