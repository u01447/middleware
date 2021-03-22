/*
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 13:50:14
 * Description: 82-Remove Duplicates from Sorted List II
 **/
 
package com.weitrue.leetcode.editor.cn;

public class RemoveDuplicatesFromSortedListIi{
    public static void main(String[] args){
        Solution s = new RemoveDuplicatesFromSortedListIi().new Solution();
    }
    
    //Given the head of a sorted linked list, delete all nodes that have duplicate n
    //umbers, leaving only distinct numbers from the original list. Return the linked
    //list sorted as well.
    //
    // Example 1:
    //Input: head = [1,2,3,3,4,4,5]
    //Output: [1,2,5]
    //
    // Example 2:
    //Input: head = [1,1,1,2,3]
    //Output: [2,3]
    //
    // Constraints:
    // The number of nodes in the list is in the range [0, 300].
    // -100 <= Node.val <= 100
    // The list is guaranteed to be sorted in ascending order.
    //
    // Related Topics 链表
    // 👍 469 👎 0
 
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
            ListNode hair = new ListNode(0, head), fast = head, slow = hair;
            while (fast != null) {
                while (fast.next != null && fast.val == fast.next.val) {
                    fast = fast.next;
                }
                if (slow.next == fast) {
                    slow = slow.next;
                }else {
                    slow.next = fast.next;
                }
                fast = fast.next;
            }
            return hair.next;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)

}
