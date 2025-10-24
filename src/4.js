/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function (nums1, nums2) {
  const nums = nums1.concat(nums2).sort((a, b) => a - b)
  const median = Math.floor(nums.length / 2)
  if (nums.length % 2 === 1) {
    return nums[median]
  }

  return (nums[median-1] + nums[median]) / 2
};

console.log(findMedianSortedArrays([1,3], [2]))
console.log(findMedianSortedArrays([1,2], [3,4]))
console.log(findMedianSortedArrays([3], [-2,-1]))