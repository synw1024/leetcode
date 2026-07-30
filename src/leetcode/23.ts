function mergeKLists(lists: Array<ListNode | null>): ListNode | null {
  if (!lists.length) return null

  lists = lists.filter(Boolean)
  if (!lists.length) return null

  let minIndex = 0
  for (let i = 1; i < lists.length; i++) {
    if (lists[i]!.val < lists[minIndex]!.val) {
      minIndex = i
    }
  }
  const res = lists[minIndex]!
  let pointer = res
  lists[minIndex] = lists[minIndex]!.next
  if (!lists[minIndex]) {
    lists.splice(minIndex, 1)
  }

  while (lists.length) {
    let minIndex = 0
    for (let i = 1; i < lists.length; i++) {
      if (lists[i]!.val < lists[minIndex]!.val) {
        minIndex = i
      }
    }
    pointer.next = lists[minIndex]!
    pointer = pointer.next
    lists[minIndex] = lists[minIndex]!.next
    if (!lists[minIndex]) {
      lists.splice(minIndex, 1)
    }
  }
  return res
};

type ListNode = {
  val: number
  next: ListNode | null
}
