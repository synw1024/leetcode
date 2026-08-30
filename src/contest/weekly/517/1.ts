function countSpecialIntegers(nums: number[]): number {
  const exist: {[key: number]: number} = {}
  for (let i = 0; i < nums.length; i++) {
    if (exist[nums[i]] === undefined || exist[nums[i]] + 1 === i) {
      exist[nums[i]] = i
    } else {
      exist[nums[i]] = -1
    }
  }
  return Object.values(exist).filter(v => v !== -1).length
};
