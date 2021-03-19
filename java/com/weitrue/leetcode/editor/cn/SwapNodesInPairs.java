/*
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 20:22:33
 * Description: 24-Swap Nodes in Pairs
 **/
 
package com.weitrue.leetcode.editor.cn;

public class SwapNodesInPairs{
    public static void main(String[] args){
        Solution s = new SwapNodesInPairs().new Solution();

    }
    
    //Given a linked list, swap every two adjacent nodes and return its head. 
    //
    // Example 1:
    //Input: head = [1,2,3,4]
    //Output: [2,1,4,3]
    //
    // Example 2:
    //Input: head = []
    //Output: []
    //
    // Example 3:
    //Input: head = [1]
    //Output: [1]
    //
    // Constraints:
    // The number of nodes in the list is in the range [0, 100].
    // 0 <= Node.val <= 100
    //Follow up: Can you solve the problem without modifying the values in the list'
    //s nodes? (i.e., Only nodes themselves may be changed.) Related Topics 递归 链表
    // 👍 859 👎 0

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
        public ListNode swapPairs(ListNode head) {
            ListNode hair = new ListNode(0, head), pre = hair, tail = pre;
            while (head != null) {
                for (int i=0; i<2; i++) {
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
