type ListNode = {
  val: number
  next: ListNode | null
}

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
  const res = listNode()
  let pointer: ListNode | null = res
  while (l1 || l2) {
    const sum = (l1?.val || 0) + (l2?.val || 0) + pointer!.val
    pointer!.val = sum % 10
    if (sum > 9) {
      pointer!.next = listNode(1)
    } else if (l1?.next || l2?.next) {
      pointer!.next = listNode()
    }

    l1 = l1?.next || null
    l2 = l2?.next || null
    pointer = pointer!.next
  }
  return res
};

function listNode(val = 0): ListNode {
  return {
    val,
    next: null
  }
}
