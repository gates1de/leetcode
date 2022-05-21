class TreeNode {
  val: number
  left: TreeNode | null
  right: TreeNode | null
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = (val===undefined ? 0 : val)
    this.left = (left===undefined ? null : left)
    this.right = (right===undefined ? null : right)
  }
}

function getTargetCopy(original: TreeNode | null, cloned: TreeNode | null, target: TreeNode | null): TreeNode | null {
  return helper(target, cloned)
}

function helper(target: TreeNode | null, cloned: TreeNode | null): TreeNode | null {
  if (!cloned) {
    return null
  }
  
  if (target.val === cloned.val) {
    return cloned
  }
  
  const left = helper(target, cloned.left)
  if (left) {
    return left
  }
  
  return helper(target, cloned.right)
}

function makeTree1(): TreeNode {
  const root = new TreeNode(7, null, null)
  root.left = new TreeNode(4, null, null)
  root.right = new TreeNode(3, null, null)
  root.right.left = new TreeNode(6, null, null)
  root.right.right = new TreeNode(19, null, null)
  return root
}

function main() {
  // result: 3
  const original = makeTree1()
  const target = original.right
  const cloned = makeTree1()

  // result: 
  // const original = makeTree()
  // const target = original
  // const cloned = makeTree()

  const result = getTargetCopy(original, cloned, target)
  console.dir(result)
}

main()

