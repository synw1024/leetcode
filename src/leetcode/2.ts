function addTwoNumbers(l1?: ListNode | null, l2?: ListNode | null): ListNode | null {
  const node = listNode(l1?.val, l2?.val)
  if (l1?.next || l2?.next) {
    if (node.next) {
      (l1?.next || l2?.next)!.val++
    }
    const next = addTwoNumbers(l1?.next, l2?.next)
    node.next = next
  }
  return node
};

function listNode(val1 = 0, val2 = 0): ListNode {
  let val = val1 + val2
  let next = null
  if (val >= 10) {
    val %= 10
    next = {
      val: 1,
      next: null
    }
  }
  return {
    val,
    next
  }
}

type ListNode = {
  val: number
  next: ListNode | null
}

