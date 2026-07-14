type TreeNode = {
  val: number
  left: TreeNode | null
  right: TreeNode | null
}

function hasPathSum(root: TreeNode | null, targetSum: number): boolean {
  if (!root) return false

  let existed = false
  function dfs(node: TreeNode, sum: number) {
    if (existed) return

    sum += node.val
    if (sum === targetSum && !node.left && !node.right) {
      existed = true
      return
    }

    if (node.left) {
      dfs(node.left, sum)
    }

    if (node.right) {
      dfs(node.right, sum)
    }
  }
  dfs(root, 0)
  return existed
};
