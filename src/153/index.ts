function findMin(nums: number[]): number {
  let start = 0
  let end = nums.length - 1
  while(start < end) {
    const midIndex = Math.floor((end - start) / 2) + start
    const midValue = nums[midIndex]
    if (nums[0] <= midValue) {
      start = midIndex + 1
    } else {
      end = midIndex
    }
  }
  return nums[start] < nums[0] ? nums[start] : nums[0]
};

findMin([2,3,4,5,6,1])