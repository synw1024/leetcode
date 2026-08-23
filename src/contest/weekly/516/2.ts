function findDisappearedNumbers(nums: number[], lower: number, upper: number): number[][] {
  nums = Array.from(new Set(nums)).filter(n => n >= lower && n <= upper).sort((a, b) => a - b)
  if (!nums.length) return [[lower, upper]]

  const res: number[][] = []
  for (let i = 0; i < nums.length; i++) {
    if (lower === nums[i]) {
      lower++
      continue
    }
    res.push([lower, nums[i]-1])
    lower = nums[i]+1
  }
  if (upper > nums[nums.length-1]) {
    res.push([nums[nums.length-1]+1, upper])
  }
  return res
}
