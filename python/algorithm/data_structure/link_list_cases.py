"""
Module Description: 链表 
Problem:
Solution：
Date: 2020/2/16 
Author: Wang P
"""


class ListNode:
    def __init__(self, val):
        self.val = val
        self.next = None


class DeleteNode:
    """
    链表中删除一个节点（已知要删除的节点，不知链表的head）
    思路：将要删除的链表节点node的值修改为node.next的值
         node.next指向node.next.next
         这样相当于删除了node节点
    """
    def __init__(self, node_list=None, node=None):
        if node_list and isinstance(node_list, list):
            self._head = ListNode(node_list[0])
            cur_index = self._head
            for i in range(1, len(node_list)):
                if cur_index.val == node:
                    self._node = cur_index
                cur_index.next = ListNode(node_list[i])
                cur_index = cur_index.next
        else:
            self._head = None
            self._node = None

    @staticmethod
    def delete_node_in_a_linked_list(node):
        next_node = node.next
        next_next_node = node.next.next
        node.val = next_node.val
        node.next = next_next_node

    def traversal(self):
        cur = self._head
        link_list = []
        while cur:
            link_list.append(cur.val)
            cur = cur.next
        print(link_list)


class MergeLinkedList:
    """
    合并两个有序链表
    思路：新建一个链表root，将list_one, list_two遍历，按大小放入root
    """
    def __init__(self, list_one=None, list_two=None):
        if list_one and isinstance(list_one, list) and list_two and isinstance(list_two, list):
            self.list_one = ListNode(list_one[0])
            cur_index = self.list_one
            for i in range(1, len(list_one)):
                cur_index.next = ListNode(list_one[i])
                cur_index = cur_index.next
            self.list_two = ListNode(list_two[0])
            cur_index = self.list_two
            for i in range(1, len(list_two)):
                cur_index.next = ListNode(list_two[i])
                cur_index = cur_index.next
        else:
            self.list_one = None
            self.list_two = None
        self._head = None

    def merge_linked_list(self):
        one_cur_val = self.list_one
        two_cur_val = self.list_two
        self._head = ListNode(None)
        cur = self._head
        while one_cur_val and two_cur_val:
            if one_cur_val.val < two_cur_val.val:
                node = ListNode(one_cur_val.val)
                one_cur_val = one_cur_val.next
            else:
                node = ListNode(two_cur_val.val)
                two_cur_val = two_cur_val.next
            cur.next = node
            cur = node
        cur.next = one_cur_val or two_cur_val

    def traversal(self):
        cur = self._head.next
        link_list = []
        while cur:
            link_list.append(cur.val)
            cur = cur.next
        print(link_list)


class ReverseLinkList:
    """
    单链表反转
    """
    def __init__(self, head=None):
        """链表的头部"""
        self._head = head

    def add(self, val: int):
        """
        给链表添加元素
        ：param val： 传过来的数字
        ：return：
        """
        # 创建一个节点
        node = ListNode(val)
        if self._head is None:
            self._head = node
        else:
            cur = self._head
            while cur.next is not None:
                cur = cur.next  # 移动游标
            cur.next = node  # 如果 next 后面没了证明以及到最后一个节点了

    def traversal(self):
        if not self._head:
            return
        cur = self._head
        link_list = []
        while cur:
            link_list.append(cur.val)
            cur = cur.next
        print(link_list)

    def size(self):
        """
        获取链表的大小
        ：return：
        """
        count = 0
        if self._head is None:
            return count
        else:
            cur = self._head
            while cur is not None:
                count += 1
                cur = cur.next
            return count

    def reverse_link(self):
        """
        单链表反转
        思路：
        让 cur.next 先断开即指向 none，指向设定 pre 游标指向断开的元素，然后
        cur.next 指向断开的元素，再把开始 self._head 再最后一个元素的时候.
        ：return：
        """
        if self._head is None or self.size() == 1:
            return
        else:
            pre = None
            cur = self._head
            while cur is not None:
                post = cur.next
                cur.next = pre
                pre = cur
                cur = post
            self._head = pre  # 逆向后的头节点


class LinkListIntersectionNode:

    def get_intersection_node(self, headA, headB):
        """
        cur1、cur2，2 个指针的初始位置是链表 headA、headB 头结点，cur1、cur2 两个指针一直往后遍历。
        直到 cur1 指针走到链表的末尾，然后 cur1 指向 headB；
        直到 cur2 指针走到链表的末尾，然后 cur2 指向 headA；
        然后再继续遍历；
        每次 cur1、cur2 指向 None，则将 cur1、cur2 分别指向 headB、headA。
        循环的次数越多，cur1、cur2 的距离越接近，直到 cur1 等于 cur2。则是两个链表的相交点。
        ：tye head1, head1： ListNode
        ：rtye： ListNode
        """
        if headA is not None and headB is not None:
            cur1, cur2 = headA, headB

            while cur1 != cur2:
                cur1 = cur1.next if cur1 is not None else headA
                cur2 = cur2.next if cur2 is not None else headB

            return cur1


class Node(object):

    def __init__(self, prev=None, next=None, key=None, value=None):
        self.prev, self.next, self.key, self.value = prev, next, key, value


class CircularDoubleLinkedList(object):

    def __init__(self):
        node = Node()
        node.prev, node.next = node, node
        self.rootnode = node

    def headnode(self):
        return self.rootnode.next

    def tailnode(self):
        return self.rootnode.prev

    def remove(self, node):
        if node is self.rootnode:
            return
        else:
            node.prev.next = node.next
            node.next.prev = node.prev

    def append(self, node):
        tailnode = self.tailnode()
        tailnode.next = node
        node.next = self.rootnode
        self.rootnode.prev = node


class LRUCache(object):

    def __init__(self, maxsize=16):
        self.maxsize = maxsize
        self.cache = {}
        self.access = CircularDoubleLinkedList()
        self.isfull = len(self.cache) >= self.maxsize

    def __call__(self, func):
        def wrapper(n):
            cachenode = self.cache.get(n)
            if cachenode is not None:  # hit
                self.access.remove(cachenode)
                self.access.append(cachenode)
                return cachenode.value
            else:  # miss
                value = func(n)
                if not self.isfull:
                    tailnode = self.access.tailnode()
                    newnode = Node(tailnode, self.access.rootnode, n, value)
                    self.access.append(newnode)
                    self.cache[n] = newnode
                    self.isfull = len(self.cache) >= self.maxsize
                    return value
                else:  # full
                    lru_node = self.access.headnode()
                    del self.cache[lru_node.key]
                    self.access.remove(lru_node)
                    tailnode = self.access.tailnode()
                    newnode = Node(tailnode, self.access.rootnode, n, value)
                    self.access.append(newnode)
                    self.cache[n] = newnode
                return value

        return wrapper


@LRUCache()
def fib(n):
    if n <= 2:
        return 1
    else:
        return fib(n - 1) + fib(n - 2)


if __name__ == "__main__":
    # dn = DeleteNode([5, 3, 1, 9], 3)
    # dn.traversal()
    # dn.delete_node_in_a_linked_list(dn._node)
    # dn.traversal()
    #
    # mn = MergeLinkedList([1, 3, 5], [1, 2, 4, 7])
    # mn.merge_linked_list()
    # mn.traversal()
    #
    # r_link = ReverseLinkList()
    # r_link.add(3)
    # r_link.add(5)
    # r_link.add(6)
    # r_link.add(7)
    # r_link.add(8)
    # print("对链表进行遍历")
    # r_link.traversal()
    # print(f"size：{r_link.size()}")
    # print("对链表进行逆向操作之后")
    # r_link.reverse_link()
    # r_link.traversal()

    for i in range(1, 10):
        print(fib(i))

