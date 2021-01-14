from math import sqrt

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

def maketree(list: [int]) -> TreeNode:
    if not list:
        return None

    parent = TreeNode(list[0])
    for i in range(1, len(list)):
        val = list[i]
        parentIndex = 0
        if i % 2 == 0:
            parentIndex = i / 2
        elif i % 2 == 1:
            parentIndex = i // 2 + 1
        parent = insert(parent, i, parentIndex, val)

    return parent

def insert(parent: TreeNode, i: int, parentIndex: int, val: int) -> TreeNode:
    # print('i = {}, parent.index = {}, parent.val: {}'.format(i, parentIndex, parent.val))

    if parentIndex == i / 2:
        if parent.left is None:
            parent.left = TreeNode(val)
            print('i: {}, parent.val: {}, parent.left.val: {}'.format(i, parent.val, parent.left.val))
        else:
            return insert(parent.left, i, parentIndex, val)
    elif parentIndex == i // 2 + 1:
        if parent.right is None:
            parent.right = TreeNode(val)
            print('i: {}, parent.val: {}, parent.right.val: {}'.format(i, parent.val, parent.right.val))
        else:
            return insert(parent.right, i, parentIndex, val)

    return parent

def main():
    # original = maketree([7, 4, 3, None, None, 6, 19])
    original = maketree([1, 2, 3, 4, 5, 6, 7, 8, 9, 10])

if __name__ == '__main__':
    main()

