function removeNthFromEnd(head: ListNode | null, n: number): ListNode | null {
  function recurse(pointer: ListNode) {
    if (pointer.next) {
      if (!recurse(pointer.next)) {
        pointer.next = pointer.next.next
      }
    }
    return --n
  }
  if (!recurse(head!)) {
    return head!.next
  }
  return head
};

type ListNode = {
  val: number
  next: ListNode | null
}
