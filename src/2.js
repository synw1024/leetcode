/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {
  const l3 = { val: 0 }
  let temp = l3
  while (l1 || l2) {
    if (l1 && l2) {
      temp.val += l1.val + l2.val
    } else if (l1) {
      temp.val += l1.val
    } else {
      temp.val += l2.val
    }
    l1 = l1 ? l1.next : l1
    l2 = l2 ? l2.next : l2
    if (temp.val >= 10) {
      temp.val %= 10
      temp = temp.next = { val: 1, next: null }
    } else if (l1 || l2) {
      temp = temp.next = { val: 0, next: null }
    } else {
      temp.next = null
    }
  }
  return l3
};