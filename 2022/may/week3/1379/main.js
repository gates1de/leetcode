var TreeNode = /** @class */ (function () {
    function TreeNode(val, left, right) {
        this.val = (val === undefined ? 0 : val);
        this.left = (left === undefined ? null : left);
        this.right = (right === undefined ? null : right);
    }
    return TreeNode;
}());
function getTargetCopy(original, cloned, target) {
    return helper(target, cloned);
}
function helper(target, cloned) {
    if (!cloned) {
        return null;
    }
    if (target.val === cloned.val) {
        return cloned;
    }
    var left = helper(target, cloned.left);
    if (left) {
        return left;
    }
    return helper(target, cloned.right);
}
function makeTree1() {
    var root = new TreeNode(7, null, null);
    root.left = new TreeNode(4, null, null);
    root.right = new TreeNode(3, null, null);
    root.right.left = new TreeNode(6, null, null);
    root.right.right = new TreeNode(19, null, null);
    return root;
}
function main() {
    // result: 3
    var original = makeTree1();
    var target = original.right;
    var cloned = makeTree1();
    // result: 
    // const original = makeTree()
    // const target = original
    // const cloned = makeTree()
    var result = getTargetCopy(original, cloned, target);
    console.dir(result);
}
main();
