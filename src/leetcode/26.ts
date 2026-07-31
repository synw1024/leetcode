function removeDuplicates(nums: number[]): number {
  let removed = 0
  for (let i = 0; i < nums.length; i++) {
    if (nums[i + 1] !== undefined && nums[i] === nums[i + 1]) {
      removed++
    } else {
      nums.splice(i - removed, removed)
      i -= removed
      removed = 0
    }
  }
  return nums.length
};
