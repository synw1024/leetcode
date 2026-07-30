import { arrayToListNode } from "./utils"

function reverseKGroup(head: ListNode | null, k: number): ListNode | null {
  if (!head) return null
  if (!head.next || k === 1) return head

  function recurse(node: ListNode, n: number): (ListNode | null)[] {
    if (!n) {
      return [node, node.next]
    }

    if (node.next) {
      const [last, next] = recurse(node.next, n - 1)
      if (last) {
        node.next.next = node
      }
      return [last, next]
    }

    return [null, null]
  }

  function groupRecurse(node: ListNode): ListNode | null {
    const [last, next] = recurse(node, k - 1)
    if (next) {
      const groupLast = groupRecurse(next)
      if (groupLast) {
        node.next = groupLast
      } else {
        node.next = next
      }
    } else if (last) {
      node.next = null
    }
    
    return last
  }
  return groupRecurse(head)
};

reverseKGroup(arrayToListNode([1,2,3,4,5]), 3)

type ListNode = {
  val: number
  next: ListNode | null
}
