"""
Module Description: 拓扑排序
Problem: 对应于该图的拓扑排序,每一个有向无环图都至少存在一种拓扑排序
Solution：
Date: 2020/3/23 
Author: Wang P
"""
import pysnooper
from typing import Mapping


class TopologicalSort:

    @pysnooper.snoop()  # pysnooper 调试神器
    def topological_sort(self, graph: Mapping):
        # in_degrees = {'a'： 0, 'b'： 0, 'c'： 0, 'd'： 0, 'e'： 0, 'f'： 0}
        in_degrees = dict((u, 0) for u in graph)
        for u in graph:
            for v in graph[u]:  # 根据键找出值也就是下级节点
                in_degrees[v] += 1  # 对获取到的下级节点的入度加 1
        # 循环结束之后的结果： {'a'： 0, 'b'： 1, 'c'： 1, 'd'： 2, 'e'： 1, 'f'： 4}
        Q = [u for u in graph if in_degrees[u] == 0]  # 入度为 0 的节点
        in_degrees_zero = []
        while Q:
            u = Q.pop()  # 默认从最后一个移除
            in_degrees_zero.append(u)  # 存储入度为 0 的节点
            for v in graph[u]:
                in_degrees[v] -= 1  # 删除入度为 0 的节点，以及移除其指向
                if in_degrees[v] == 0:
                    Q.append(v)
        return in_degrees_zero


if __name__ == '__main__':
    # 用字典的键值表示图的节点之间的关系，键当前节点。值是后续节点。
    graph_dict = {
        'a': 'bf',  # 表示 a 指向 b 和 f
        'b': 'cdf',
        'c': 'd',
        'd': 'ef',
        'e': 'f',
        'f': ''
    }
    ts = TopologicalSort()
    t = ts.topological_sort(graph_dict)
    print(t)
