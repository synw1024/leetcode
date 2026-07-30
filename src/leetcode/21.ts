function mergeTwoLists(list1: ListNode | null, list2: ListNode | null): ListNode | null {
  if (list1 === null) return list2
  if (list2 === null) return list1

  let pointer1: ListNode, pointer2: ListNode | null
  if (list1.val <= list2.val) {
    pointer1 = list1
    pointer2 = list2
  } else {
    pointer1 = list2
    pointer2 = list1
  }
  let res = pointer1

  while (pointer2) {
    if (!pointer1.next) {
      pointer1.next = pointer2
      break
    }
    if (pointer2.val < pointer1.next.val) {
      const temp: ListNode | null = pointer2.next
      pointer2.next = pointer1.next
      pointer1.next = pointer2
      pointer1 = pointer1.next
      pointer2 = temp
    } else {
      pointer1 = pointer1.next
    }
  }
  return res
};

type ListNode = {
  val: number
  next: ListNode | null
}
