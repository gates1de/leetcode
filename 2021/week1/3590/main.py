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
