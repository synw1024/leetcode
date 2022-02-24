var rotate = function(nums, k) {
  if (k === 0) return
  if (k > nums.length) k = k % nums.length
  const a = nums.slice(nums.length-k, nums.length)
  nums.splice(nums.length-k, k)
  nums.unshift(...a)
};