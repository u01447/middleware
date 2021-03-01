"""
Module Description: 二叉树的遍历：前序、中序、后序、层次遍历
                    二叉树的反转
Problem:
Solution：
Date: 2020/2/16 
Author: Wang P
"""


class TreeNode:
    def __init__(self, val, left, right):
        self.val, self.left, self.right = val, left, right


class BinTreeTravel:
    def __init__(self, node_list=None):
        if node_list and isinstance(node_list, list):
            tree_node_list = []
            for i in range(len(node_list)):
                tree_node_list.append(TreeNode(node_list[i], None, None))
            for i in range(0, len(node_list)):
                if i * 2 + 1 <= len(tree_node_list) - 1 and i * 2 + 2 <= len(tree_node_list) - 1:
                    tree_node_list[i].left = tree_node_list[i * 2 + 1]
                    tree_node_list[i].right = tree_node_list[i * 2 + 2]
                elif i * 2 + 1 <= len(tree_node_list) - 1 < i * 2 + 2:
                    tree_node_list[i].left = tree_node_list[i * 2 + 1]
                else:
                    break
            self.root = tree_node_list[0]

    def pre_order_travel(self, sub_tree):
        """
        先序遍历
        :param sub_tree:
        :return:
        """
        if sub_tree:
            print(sub_tree.val)
            self.pre_order_travel(sub_tree.left)
            self.pre_order_travel(sub_tree.right)

    def in_order_travel(self, sub_tree):
        """
        中序遍历
        :param sub_tree:
        :return:
        """
        if sub_tree:
            self.in_order_travel(sub_tree.left)
            print(sub_tree.val)
            self.in_order_travel(sub_tree.right)

    def post_order_travel(self, sub_tree):
        """
        后序遍历
        :param sub_tree:
        :return:
        """
        if sub_tree:
            self.post_order_travel(sub_tree.left)
            self.post_order_travel(sub_tree.right)
            print(sub_tree.val)

    def level_order_travel(self, root):
        """
        层次遍历
        :param root:
        :return: list[list[int]]
        """
        if not root:
            print([])
            return
        cur_nodes = [root]
        next_nodes = []
        print([i.val for i in cur_nodes])
        while cur_nodes or next_nodes:
            for node in cur_nodes:
                if node.left:
                    next_nodes.append(node.left)
                if node.right:
                    next_nodes.append(node.right)
            if next_nodes:
                print([i.val for i in next_nodes])
            cur_nodes = next_nodes
            next_nodes = []


class ReverseBinaryTree:
    """
    反转二叉树
               1                         1
            3     4                   4     3
         5   7  8   9      =>       9  8  7   5
      10                                        10
    """

    def __init__(self, node_list=None):
        if node_list and isinstance(node_list, list):
            tree_node_list = []
            for i in range(len(node_list)):
                tree_node_list.append(TreeNode(node_list[i], None, None))
            for i in range(0, len(node_list)):
                if i * 2 + 1 <= len(tree_node_list) - 1 and i * 2 + 2 <= len(tree_node_list) - 1:
                    tree_node_list[i].left = tree_node_list[i * 2 + 1]
                    tree_node_list[i].right = tree_node_list[i * 2 + 2]
                elif i * 2 + 1 <= len(tree_node_list) - 1 < i * 2 + 2:
                    tree_node_list[i].left = tree_node_list[i * 2 + 1]
                else:
                    break
            self.root = tree_node_list[0]

    def reverse_tree(self, root):
        if root and root.left and root.right:
            root.left, root.right = root.right, root.left
            self.reverse_tree(root.left)
            self.reverse_tree(root.right)
        self.root = root


class FindTreeKthSmallest:
    """
    二叉搜索树中第 K 小的元素
    二叉搜索树按照中序遍历的顺序打印出来正好就是排序好的顺序。所以对其遍历一个节点就进行计数，计数达到 k 的时候就结束。
    """
    count = 0
    node_val = None

    def kth_smallest(self, root: TreeNode, k):
        """
        ：type root： TreeNode
        ：type k： int
        ：rtype： int
        """
        self.dfs(root, k)
        return self.node_val

    def dfs(self, node: TreeNode, k):
        if not node:
            return
        self.dfs(node.left, k)
        self.count = self.count + 1
        if self.count == k:
            self.node_val = node.val
            # 将该节点的左右子树置为 None,来结束递归，减少时间复杂度
            node.left = None
            node.right = None
        self.dfs(node.right, k)


if __name__ == "__main__":
    bt = BinTreeTravel([1, 3, 4, 5, 7, 8, 9, 10])
    bt.pre_order_travel(bt.root)
    bt.in_order_travel(bt.root)
    bt.post_order_travel(bt.root)
    bt.level_order_travel(bt.root)
    print("#################################################")
    it = ReverseBinaryTree()
    it.reverse_tree(bt.root)
    bt.pre_order_travel(bt.root)
