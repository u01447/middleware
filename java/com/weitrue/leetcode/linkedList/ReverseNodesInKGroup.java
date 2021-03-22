/*
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 17:42:34
 * Description: 25-Reverse Nodes in k-Group
 **/
 
package com.weitrue.leetcode.editor.cn;

public class ReverseNodesInKGroup{
    public static void main(String[] args){
        Solution s = new ReverseNodesInKGroup().new Solution();
    }

    //Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
    //k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
    // Follow up:
    // Could you solve the problem in O(1) extra memory space?
    // You may not alter the values in the list's nodes, only nodes itself may be changed.

    // Example 1:
    //Input: head = [1,2,3,4,5], k = 2
    //Output: [2,1,4,3,5]

    // Example 2:
    //Input: head = [1,2,3,4,5], k = 3
    //Output: [3,2,1,4,5]

    // Example 3:
    //Input: head = [1,2,3,4,5], k = 1
    //Output: [1,2,3,4,5]

    // Example 4:
    //Input: head = [1], k = 1
    //Output: [1]

    // Constraints:
    // The number of nodes in the list is in the range sz.
    // 1 <= sz <= 5000
    // 0 <= Node.val <= 1000
    // 1 <= k <= sz
    //
    // Related Topics 链表
    // 👍 977 👎 0
 
    //leetcode submit region begin(Prohibit modification and deletion)
    /**
     * Definition for singly-linked list.
     */
    public class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }

    class Solution {
        public ListNode reverseKGroup(ListNode head, int k) {
            ListNode hair = new ListNode(0, head), pre = hair, tail = pre;
            while (head != null) {
                //tail = pre;
                for (int i=0; i<k; i++) {
                    tail = tail.next;
                    if (tail == null) {
                        return hair.next;
                    }
                }
                // 反转长度为k的链表，并返回链表的新头尾节点
                ListNode[] ln = reverse(head, tail);
                head = ln[0];
                tail = ln[1];
                pre.next = head;
                pre = tail;
                head = pre.next;
            }

            return hair.next;
        }

        public ListNode[] reverse(ListNode head, ListNode tail){
            ListNode pre = tail.next, cur = head, next = null;

            while (tail != pre) {
                next = cur.next;
                cur.next = pre;
                pre = cur;
                cur = next;
            }

            return new ListNode[]{tail, head};
        }
    }
//leetcode submit region end(Prohibit modification and deletion)

}
