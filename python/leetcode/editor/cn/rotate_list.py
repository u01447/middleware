"""
Module Description: 61-旋转链表
Solution：
Date: 2021-03-18 20:23:59
Author: Wang P
Problem:  给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

          示例 1:
          输入: 1->2->3->4->5->NULL, k = 2
         输出: 4->5->1->2->3->NULL
         解释:
         向右旋转 1 步: 5->1->2->3->4->NULL
         向右旋转 2 步: 4->5->1->2->3->NULL

          示例 2:
          输入: 0->1->2->NULL, k = 4
         输出: 2->0->1->NULL
         解释:
         向右旋转 1 步: 2->0->1->NULL
         向右旋转 2 步: 1->2->0->NULL
         向右旋转 3 步: 0->1->2->NULL
         向右旋转 4 步: 2->0->1->NULL
          Related Topics 链表 双指针
          👍 444 👎 0
"""
# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:

    def rotateRight(self, head: ListNode, k: int) -> ListNode:
        if not head or not head.next:
            return head

        old_tail = head
        length = 1
        while old_tail.next:
            old_tail = old_tail.next
            length += 1

        old_tail.next = head

        new_tail = head
        for i in range(0, length - (k % length) - 1):
            new_tail = new_tail.next

        head = new_tail.next
        new_tail.next = None

        return head

# leetcode submit region end(Prohibit modification and deletion)
