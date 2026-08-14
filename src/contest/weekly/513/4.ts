function countRatioSubarrays(nums: number[], a: number, b: number): number {
  const sorted = [0]
  let last = 0, res = 0
  for (let i = 0; i < nums.length; i++) {
    const val = nums[i] % 2 > 0 ? a : -b
    last += val
    res += insert(sorted, last)
  }
  return res
}

function insert(sorted: number[], n: number) {
  let start = 0, end = sorted.length - 1
  while (start <= end) {
    const mid = Math.floor((end - start) / 2) + start
    if (sorted[mid] <= n) {
      start = mid + 1
    } else {
      end = mid - 1
    }
  }
  sorted.splice(start, 0, n)
  return start
}

// console.log(countRatioSubarrays([1, 2, 1, 2], 3, 2))
// console.log(countRatioSubarrays([2,2,1], 2, 1))
console.log(countRatioSubarrays([2,2,2], 1, 1))
