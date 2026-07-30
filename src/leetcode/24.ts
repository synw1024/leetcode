function swapPairs(head: ListNode | null): ListNode | null {
  if (!head) return null
  if (!head.next) return head

  const next = head.next
  head.next = head.next.next
  next.next = head
  head = next

  let last = head.next!
  let pointer: ListNode | null = last.next
  while(pointer && pointer.next) {
    const next = pointer.next
    pointer.next = pointer.next.next
    next.next = pointer
    last.next = next
    last = pointer
    pointer = pointer.next
  }
  return head
};

type ListNode = {
  val: number
  next: ListNode | null
}
