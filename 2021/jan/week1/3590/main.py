from copy import deepcopy

from util import maketree

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def getTargetCopy(self, original: TreeNode, cloned: TreeNode, target: TreeNode) -> TreeNode:
        if cloned.val == target.val:
            return cloned

        cloned_left = None
        if cloned.left:
            cloned_left = self.getTargetCopy(original, cloned.left, target)

        if cloned_left:
            return cloned_left

        cloned_right = None
        if cloned.right:
            cloned_right = self.getTargetCopy(original, cloned.right, target)

        return cloned_right

def main():
    # list = [7, 4, 3, None, None, 6, 19]
    list = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    original = maketree(list)
    cloned = deepcopy(original)
    solution = Solution()
    target = original.left.left
    print('target id = {}, val = {}'.format(id(target), target.val))
    result = solution.getTargetCopy(original, cloned, target)
    print('result id = {}, val = {}'.format(id(result), result.val))

if __name__ == '__main__':
    main()

