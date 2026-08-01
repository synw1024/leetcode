export function arrayToListNode(nums: number[]) {
  let pointer: ListNode | null
  let head: ListNode
  for (let i = 0; i < nums.length; i++) {
    const node = {
      val: nums[i],
      next: null
    }
    if (i === 0) {
      pointer = node
      head = node
    } else {
      pointer!.next = node
      pointer = node
    }
  }
  return head!
}

type ListNode = {
  val: number
  next: ListNode | null
}