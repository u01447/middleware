/*
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 20:20:15
 * Description: 61-Rotate List
 **/
 
package com.weitrue.leetcode.linear.linkedList;

public class RotateList{
    public static void main(String[] args){

        Solution s = new RotateList().new Solution();
        ListNode n5 = new RotateList().new ListNode(5, null);
        ListNode n4 = new RotateList().new ListNode(4, n5);
        ListNode n3 = new RotateList().new ListNode(3, n4);
        ListNode n2 = new RotateList().new ListNode(2, n3);
        ListNode n1 = new RotateList().new ListNode(1, n2);
        ListNode n = s.rotateRight(n1, 2);

        while (n != null) {
            System.out.print(n.val);
            System.out.print("->");
            n = n.next;
        }
    }
    
    //Given the head of a linked list, rotate the list to the right by k places. 
    //
    // Example 1:
    //Input: head = [1,2,3,4,5], k = 2
    //Output: [4,5,1,2,3]
    //
    // Example 2:
    //Input: head = [0,1,2], k = 4
    //Output: [2,0,1]
    //
    // Constraints:
    // The number of nodes in the list is in the range [0, 500].
    // -100 <= Node.val <= 100
    // 0 <= k <= 2 * 109
    //
    // Related Topics 链表 双指针
    // 👍 444 👎 0

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
        public ListNode rotateRight(ListNode head, int k) {
            if (head == null || head.next == null) {
                return head;
            }
            int length = 1;
            ListNode oldTail = head;
            while (oldTail.next != null) {
                oldTail = oldTail.next;
                length++;
            }
            oldTail.next = head;
            ListNode newTail = head;
            for (int i = 0; i < length - (k % length) - 1; i++){
                newTail = newTail.next;
            }
            ListNode newHead = newTail.next;
            newTail.next = null;
            return newHead;
        }

    }
    //leetcode submit region end(Prohibit modification and deletion)

}
