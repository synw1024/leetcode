const tree = {
  val: 1,
  left: {
    val: 2,
    left: {
      val: 4
    },
    right: {
      val: 5
    }
  },
  right: {
    val: 3,
    left: {
      val: 6
    },
    right: {
      val: 7
    }
  }
}

function solve(node) {
  if (node.left) {
    solve(node.left)
  }
  console.log(node.val)
  if (node.right) {
    solve(node.right)
  }
}

solve(tree)