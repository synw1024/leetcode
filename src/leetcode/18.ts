function fourSum(nums: number[], target: number): number[][] {
  nums.sort((a, b) => a - b)
  const res: number[][] = []
  for (let i = 0; i < nums.length;) {
    for (let j = i + 1; j < nums.length;) {
      for (let k = j + 1, l = nums.length - 1; k < l;) {
        const sum = nums[i] + nums[j] + nums[k] + nums[l]
        if (sum === target) {
          res.push([nums[i], nums[j], nums[k], nums[l]])
          do {
            k++
          } while (nums[k] === nums[k-1]);
          do {
            l--
          } while(nums[l] === nums[l+1]);
        } else if (sum < target) {
          do {
            k++
          } while (nums[k] === nums[k-1]);
        } else {
          do {
            l--
          } while(nums[l] === nums[l+1]);
        }
      }
      do {
        j++
      } while(nums[j] === nums[j-1])
    }
    do {
      i++
    } while(nums[i] === nums[i-1])
  }
  return res
};
