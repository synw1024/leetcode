function threeSumClosest(nums: number[], target: number): number {
  nums.sort((a, b) => a - b)
  let closest = nums[0] + nums[1] + nums[nums.length - 1]
  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1, k = nums.length - 1; j < k;) {
      const sum = nums[i] + nums[j] + nums[k]
      const diff = Math.abs(target - sum)
      const prevDiff = Math.abs(target - closest)
      if (diff < prevDiff) {
        closest = sum
      }
      if (sum < target) {
        j++
      } else if(sum > target) {
        k--
      } else {
        return sum
      }
    }
  }
  return closest
};

threeSumClosest([1, 2, 7, 13], 12)
