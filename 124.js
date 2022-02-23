const root = {
  val: -1,
  right: {
    val: 9,
    left: {
      val: -6
    },
    right: {
      val: 3,
      right: {
        val: -2
      }
    },
  }
}

var maxPathSum = function (root) {
  let outMax = -Number.MAX_SAFE_INTEGER

  function traversal(node) {
    if (!node.left && !node.right) return [node.val, 0]

    let [left, leftbackup, right, rightbackup] = [0, -Number.MAX_SAFE_INTEGER, 0, -Number.MAX_SAFE_INTEGER]
    if (node.left) {
      [left, leftbackup] = traversal(node.left)
    }
    if (node.right) {
      [right, rightbackup] = traversal(node.right)
    }
    if (leftbackup > right + node.val) {
      outMax = Math.max(left + leftbackup, outMax)
    }
    if (rightbackup > left + node.val) {
      outMax = Math.max(right + rightbackup, outMax)
    }

    let [max, min] = left > right ? [left, right] : [right, left]

    if (max <= 0) {
      return [node.val, 0]
    }

    if (min < 0) {
      min = 0
    }

    return [node.val + max, min]
  }
  let [sum, backup] = traversal(root)

  return Math.max(outMax, sum+backup, sum)
};

const res = maxPathSum(root)
console.log(res)
