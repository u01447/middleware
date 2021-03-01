"""
Module Description: 基于collections.OrderedDict实现LRU_cache(最近最少使用缓存算法)
Problem:
Solution：
Date: 2020/2/13 
Author: Wang P
"""
from collections import OrderedDict


class LURCache:
    def __init__(self, capacity=128):
        self._od = OrderedDict()
        self._capacity = capacity

    def get(self, key):
        if key in self._od:
            val = self._od[key]
            self._od.move_to_end(key)
            return val
        else:
            return -1

    def push(self, key, value):
        if key in self._od:
            del self._od[key]
            self._od[key] = value
        else:  # insert
            self._od[key] = value
            if len(self._od) > self._capacity:
                self._od.popitem(last=False)

    def get_items(self):
        return self._od.items()


if __name__ == "__main__":
    lru_cache = LURCache(3)
    lru_cache.push("1", 1)
    lru_cache.push("2", 2)
    lru_cache.push("3", 3)
    lru_cache.push("4", 4)
    lru_cache.get("2")
    print(lru_cache.get_items())
