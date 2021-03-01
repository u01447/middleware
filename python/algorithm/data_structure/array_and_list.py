# -*- coding: utf-8 -*- 
"""
Module Description:
Time : 2020/11/20 16:59 
Author : Wang P
File : array_and_list.py
version: 
"""
import ctypes


class Matrix:
    """ 最好用pandas的DataFrame
    Matrix(rows, ncols): constructor
    numCols()
    getitem(row, col)
    setitem(row, col, val)
    scaleBy(scalar): 每个元素乘scalar
    transpose(): 返回transpose转置
    add(rhsMatrix):    size must be the same
    subtract(rhsMatrix)
    multiply(rhsMatrix)
    """

    def __init__(self, num_rows, num_cols):
        self._the_grid = Array2D(num_rows, num_cols)
        self._the_grid.clear(0)

    @property
    def num_rows(self):
        return self._the_grid.num_rows

    @property
    def NumCols(self):
        return self._the_grid.numCols

    def __getitem__(self, ndx_tuple):
        return self._the_grid[ndx_tuple[0], ndx_tuple[1]]

    def __setitem__(self, ndx_tuple, scalar):
        self._the_grid[ndx_tuple[0], ndx_tuple[1]] = scalar

    def scaleBy(self, scalar):
        for r in range(self.num_rows):
            for c in range(self.Num_cols):
                self[r, c] *= scalar

    def __add__(self, rhsMatrix):
        assert (rhsMatrix.num_rows == self.num_rows and
                rhsMatrix.numCols == self.Num_cols)
        newMartrix = Matrix(self.num_rows, self.Num_cols)
        for r in range(self.num_rows):
            for c in range(self.Num_cols):
                newMartrix[r, c] = self[r, c] + rhsMatrix[r, c]


class Array2D:
    """ 要实现的方法
    Array2D(nrows, ncols):    constructor
    numRows()
    numCols()
    clear(value)
    getitem(i, j)
    setitem(i, j, val)
    """

    def __init__(self, num_rows, num_cols):
        self._the_rows = Array(num_rows)  # 数组的数组
        for i in range(num_rows):
            self._the_rows[i] = Array(num_cols)

    @property
    def num_rows(self):
        return len(self._the_rows)

    @property
    def num_cols(self):
        return len(self._the_rows[0])

    def clear(self, value):
        for row in self._the_rows:
            row.clear(value)

    def __getitem__(self, ndx_tuple):  # ndx_tuple: (x, y)
        assert len(ndx_tuple) == 2
        row, col = ndx_tuple[0], ndx_tuple[1]
        assert (0 <= row < self.num_rows and 0 <= col < self.num_cols)
        the_1d_array = self._the_rows[row]
        return the_1d_array[col]

    def __setitem__(self, ndx_tuple, value):
        assert len(ndx_tuple) == 2
        row, col = ndx_tuple[0], ndx_tuple[1]
        assert (0 <= row < self.num_rows and 0 <= col < self.num_cols)
        the_1d_array = self._the_rows[row]
        the_1d_array[col] = value


class Array:
    def __init__(self, size):
        assert size > 0, 'array size must be > 0'
        self._size = size
        PyArrayType = ctypes.py_object * size
        self._elements = PyArrayType()
        self.clear(None)

    def __len__(self):
        return self._size

    def __getitem__(self, index):
        assert 0 <= index < len(self), 'out of range'
        return self._elements[index]

    def __setitem__(self, index, value):
        assert 0 <= index < len(self), 'out of range'
        self._elements[index] = value

    def clear(self, value):
        """ 设置每个元素为value """
        for i in range(len(self)):
            self._elements[i] = value

    def __iter__(self):
        return _ArrayIterator(self._elements)


class _ArrayIterator:
    def __init__(self, items):
        self._items = items
        self._idx = 0

    def __iter__(self):
        return self

    def __next__(self):
        if self._idx < len(self._items):
            val = self._items[self._idx]
            self._idx += 1
            return val
        else:
            raise StopIteration


a = [1, 2, 3]
print(a[-2])
