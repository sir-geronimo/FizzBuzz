// https://leetcode.com/problems/sum-of-left-leaves/

/* Definition for a binary tree node. */
class TreeNode {
  constructor(
    public readonly val: number = 0,
    public readonly left: TreeNode | null = null,
    public readonly right: TreeNode | null = null,
  ) {
  }
}

function sumOfLeftLeaves(root?: TreeNode | null): number {
  let result = 0

  const isLeaf = (node: TreeNode) => node.left === null && node.right === null

  if (root) {
    if (root.left) {
      result += isLeaf(root.left)
        ? root.left.val
        : sumOfLeftLeaves(root.left)
    }

    result += sumOfLeftLeaves(root.right)
  }

  return result
}

const root = new TreeNode(
  3,
  new TreeNode(9),
  new TreeNode(
    20,
    new TreeNode(15),
    new TreeNode(7),
  )
)

console.log(sumOfLeftLeaves(root) === 24)
