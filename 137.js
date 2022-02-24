var singleNumber = function (nums) {
  nums.sort()
  for (let i = 0; i < nums.length; i++) {
    if (i === 0 && nums[i] !== nums[i + 1]) return nums[i]
    if (i === nums.length - 1) return nums[i]
    if (nums[i] !== nums[i + 1] && nums[i] !== nums[i - 1]) return nums[i]
  }
};

const res = singleNumber([0, 1, 0, 1, 0, 1, 99])