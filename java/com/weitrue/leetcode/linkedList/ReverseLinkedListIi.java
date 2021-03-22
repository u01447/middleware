/*
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 19:05:32
 * Description: 92-Reverse Linked List II
 **/
 
package com.weitrue.leetcode.editor.cn;

public class ReverseLinkedListIi{
    public static void main(String[] args){

        Solution s = new ReverseLinkedListIi().new Solution();
        ListNode n5 = new ReverseLinkedListIi().new ListNode(5, null);
        ListNode n4 = new ReverseLinkedListIi().new ListNode(4, n5);
        ListNode n3 = new ReverseLinkedListIi().new ListNode(3, n4);
        ListNode n2 = new ReverseLinkedListIi().new ListNode(2, n3);
        ListNode n1 = new ReverseLinkedListIi().new ListNode(1, n2);
        ListNode n = s.reverseBetween(n1, 2, 4);
        while (n != null && n.val > 0) {
            System.out.print(n.val);
            System.out.print("->");
            n = n.next;
        }
    }
    
    //Given the head of a singly linked list and two integers left and right where left <= right,
    // reverse the nodes of the list from position left to position right, and return the reversed list.
    //
    // Example 1:
    //Input: head = [1,2,3,4,5], left = 2, right = 4
    //Output: [1,4,3,2,5]
    //
    // Example 2:
    //Input: head = [5], left = 1, right = 1
    //Output: [5]
    //
    // Constraints:
    // The number of nodes in the list is n.
    // 1 <= n <= 500
    // -500 <= Node.val <= 500
    // 1 <= left <= right <= n
    //
    //Follow up: Could you do it in one pass? Related Topics 链表
    // 👍 709 👎 0
 
    //leetcode submit region begin(Prohibit modification and deletion)
    /**
     * Definition for singly-linked list.
     */
    class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }

    class Solution {
        public ListNode reverseBetween(ListNode head, int left, int right) {
            // 需要虚拟节点
            ListNode hair = new ListNode(0, head), cur = hair ;
            int n = right - left + 1;
            while (left > 1) {
                cur = cur.next;
                left--;
            }
            cur.next = reverseN(cur.next, n);
            return hair.next;
        }

        public ListNode reverse(ListNode head, int x) {
            if (head == null) {
                return null;
            }
            System.out.println(head.val + "  "+ x);
            // 从头节点开始，反转n个节点
            ListNode pre = new ListNode(0), cur = head, next = null;
            while (x > 0) {
                next = cur.next;
                cur.next = pre.next;
                pre.next = cur;
                cur = next;
                x--;
            }
            head.next = cur;
            return pre.next;
        }

        public ListNode reverseN(ListNode head, int n) {
            if (n == 1) {
                return head;
            }
            ListNode tail = head.next, p = reverseN(head.next, n-1);
            head.next = tail.next;
            tail.next = head;
            return p;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

}
