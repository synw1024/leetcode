type TreeNode = {
  val: number
  left: TreeNode | null
  right: TreeNode | null
}

function kthSmallest(root: TreeNode | null, k: number): number {
  let res = -1
  function dfs(node: TreeNode | null, prev: number): number {
    if (!node) return 0

    const left = dfs(node.left, prev)
    if (left + prev + 1 === k) res = node.val
    const right = dfs(node.right, left + prev + 1)

    return left + 1 + right
  }
  dfs(root, 0)
  return res
};