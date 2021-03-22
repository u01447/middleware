/*
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 17:30:32
 * Description: 206-Reverse Linked List
 **/
 
package com.weitrue.leetcode.linear.linkedList;

public class ReverseLinkedList{
    public static void main(String[] args){
        Solution s = new ReverseLinkedList().new Solution();
    }
    
    //Given the head of a singly linked list, reverse the list, and return the reversed list.
    //
    // Example 1:
    //Input: head = [1,2,3,4,5]
    //Output: [5,4,3,2,1]
    //
    // Example 2:
    //Input: head = [1,2]
    //Output: [2,1]
    //
    // Example 3:
    //Input: head = []
    //Output: []
    //
    // Constraints:
    // The number of nodes in the list is the range [0, 5000].
    // -5000 <= Node.val <= 5000
    //
    // Follow up: A linked list can be reversed either iteratively or recursively. C
    //ould you implement both?
    // Related Topics 链表
    // 👍 1568 👎 0
 
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
        public ListNode reverseList(ListNode head) {
            ListNode pre = null, cur = head, next = null;
            while (cur != null) {
                next = cur.next;
                cur.next = pre;
                pre = cur;
                cur = next;
            }
            return pre;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)

}
