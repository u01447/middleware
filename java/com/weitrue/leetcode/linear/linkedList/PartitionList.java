/*
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 20:29:46
 * Description: 86-Partition List
 **/
 
package com.weitrue.leetcode.linear.linkedList;

public class PartitionList{
    public static void main(String[] args){
        Solution s = new PartitionList().new Solution();
    }
    
    // Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
    // You should preserve the original relative order of the nodes in each of the two partitions.
    //
    // Example 1:
    //Input: head = [1,4,3,2,5,2], x = 3
    //Output: [1,2,2,4,3,5]
    //
    // Example 2:
    //Input: head = [2,1], x = 2
    //Output: [1,2]
    //
    // Constraints:
    // The number of nodes in the list is in the range [0, 200].
    // -100 <= Node.val <= 100
    // -200 <= x <= 200
    //
    // Related Topics 链表 双指针
    // 👍 378 👎 0

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
        public ListNode partition(ListNode head, int x) {
            ListNode less = new ListNode(0, head), greater = new ListNode(0, head);
            ListNode curLess = less, curGreater = greater;
            while (head != null) {
                if (head.val < x) {
                    curLess.next = head;
                    curLess = curLess.next;
                } else {
                    curGreater.next = head;
                    curGreater = curGreater.next;
                }
                head = head.next;
            }
            curGreater.next = null;
            curLess.next = greater.next;
            return less.next;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

}
