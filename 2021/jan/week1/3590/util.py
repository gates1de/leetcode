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
        parent_index = 0
        if i % 2 == 0:
            parent_index = int(i / 2)
        elif i % 2 == 1:
            parent_index = int(i / 2) + 1
        insert(parent, i + 1, 1, val)

    return parent

def insert(parent: TreeNode, i: int, parent_index: int, val: int) -> bool:
    # print('i = {}, parent.index = {}, parent.val: {}'.format(i, parent_index, parent.val if parent else None))
    if parent is None or val is None:
        # print('Cannot insert index: {}, val: {}'.format(i, val))
        return False

    if parent_index * 2 == i:
        if parent.left is None:
            parent.left = TreeNode(val)
            # print('i: {}, parent.val: {}, parent.left.val: {}'.format(i, parent.val, parent.left.val))
            return True
    elif parent_index * 2 + 1 == i:
        if parent.right is None:
            parent.right = TreeNode(val)
            # print('i: {}, parent.val: {}, parent.right.val: {}'.format(i, parent.val, parent.right.val))
            return True

    if parent_index * 2 > i:
        # print('Cannot insert index: {}'.format(i))
        return False

    is_inserted = insert(parent.left, i, parent_index * 2, val)
    if is_inserted:
        return True

    return insert(parent.right, i, parent_index * 2 + 1, val)

def print_result(target: TreeNode, index: int, is_left: bool) -> bool:
    if index == 1:
        print('root node val: {}'.format(target.val))
    else:
        print('current node index: {}, val: {}'.format(index, target.val))

    if target.left:
        print_result(target.left, index * 2, True)

    if target.right:
        print_result(target.right, index * 2 + 1, False)

