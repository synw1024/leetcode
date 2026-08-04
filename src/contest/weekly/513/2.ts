/**
 * fn(i, j) = fn()
 */
function countRatioSubarrays(nums: number[], a: number, b: number): number {
  let count = 0
  for (let start = 0; start < nums.length; start++) {
    let x = 0, y = 0
    for (let end = start; end < nums.length; end++) {
      if (nums[end] % 2 > 0) {
        y++
      } else {
        x++
      }
      if (y === 0) continue

      if (x / y <= a / b) count++
    }
  }
  return count
}

console.log(countRatioSubarrays([1, 2, 1, 2], 3, 2))