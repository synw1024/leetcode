function twoSum(nums: number[], target: number): number[] {
  const map: { [key: number]: number } = {}
  for (let i = 0; i < nums.length; i++) {
    const prevIndex = map[target - nums[i]]
    if (prevIndex !== undefined) {
      return [prevIndex, i]
    }
    map[nums[i]] = i
  }
};

console.log(twoSum([2,7,11,15], 9));