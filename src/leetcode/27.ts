function removeElement(nums: number[], val: number): number {
  let removed = 0
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] === val) {
      removed++
    } else {
      nums.splice(i - removed, removed)
      i -= removed
      removed = 0
    }
  }
  if (removed > 0) {
    nums.splice(nums.length - removed, removed)
  }
  return nums.length
};
