/*
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 13:48:45
 * Description: 83-Remove Duplicates from Sorted List
 **/
 
package com.weitrue.leetcode.linear.linkedList;

public class RemoveDuplicatesFromSortedList{
    public static void main(String[] args){
        Solution s = new RemoveDuplicatesFromSortedList().new Solution();
    }
    
    //Given the head of a sorted linked list, delete all duplicates such that each e
    //lement appears only once. Return the linked list sorted as well.
    //
    //
    // Example 1:
    //
    //
    //Input: head = [1,1,2]
    //Output: [1,2]
    //
    //
    // Example 2:
    //
    //
    //Input: head = [1,1,2,3,3]
    //Output: [1,2,3]
    //
    //
    //
    // Constraints:
    //
    //
    // The number of nodes in the list is in the range [0, 300].
    // -100 <= Node.val <= 100
    // The list is guaranteed to be sorted in ascending order.
    //
    // Related Topics 链表
    // 👍 488 👎 0
 
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
        public ListNode deleteDuplicates(ListNode head) {
            if(head == null) {
                return null;
            }
            ListNode fast = head.next, slow = head;
            while (fast != null) {
                if (fast.val == slow.val) {
                    slow.next = fast.next;
                    fast = slow.next;
                }else {
                    fast = fast.next;
                    slow = slow.next;
                }
            }
            return head;
        }
}
//leetcode submit region end(Prohibit modification and deletion)

}
