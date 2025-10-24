const inorder = [9, 3, 15, 20, 7]
const postorder = [9, 15, 7, 20, 3]

var buildTree = function (inorder, postorder) {
  const indexMap = {}
  inorder.forEach((val, index) => {
    indexMap[val] = index
  });

  function build(leftIndex, rightIndex) {
    if (leftIndex > rightIndex) return null
    const val = postorder.pop()
    const inorderIndex = indexMap[val]
    const node = { val }
    node.right = build(inorderIndex + 1, rightIndex)
    node.left = build(leftIndex, inorderIndex - 1)
    return node
  }
  return build(0, inorder.length - 1)
};
