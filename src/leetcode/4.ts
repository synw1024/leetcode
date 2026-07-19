function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
  const isOdd = (nums1.length + nums2.length) % 2 > 0
  let index1, index2
  if (nums1.length % 2 > 0) {
    index1 = Math.floor(nums1.length / 2)
    if (nums2.length === 0) {
      return nums1[index1]
    }
    index2 = Math.floor(nums2.length / 2) - 1
  } else {
    if (nums1.length === 0) {
      const mid2 = Math.floor(nums2.length / 2)
      if (nums2.length % 2 > 0) {
        return nums2[mid2]
      } else {
        return (nums2[mid2 - 1] + nums2[mid2]) / 2
      }
    }
    index1 = Math.floor(nums1.length / 2) - 1
    if (nums2.length === 0) {
      return (nums1[index1] + nums1[index1 + 1]) / 2
    }
    index2 = Math.floor(nums2.length / 2) - 1
    if (nums2.length % 2 > 0) {
      index2 = Math.floor(nums2.length / 2)
    }
  }

  if (nums1[index1] > (nums2[index2 + 1] ?? Number.MAX_SAFE_INTEGER)) {
    if (nums1.length === 1 && isOdd) {
      return nums2[index2 + 1]
    } else if (nums1.length === 1) {
      return (Math.min(nums1[index1], nums2[index2 + 2] ?? Number.MAX_SAFE_INTEGER) + nums2[index2 + 1]) / 2
    }

    if (nums2.length === 1) {
      return (Math.max(nums1[index1 - 1], nums2[index2 + 1]) + nums1[index1]) / 2
    }

    const removed1 = nums1.length - (index1 + 1)
    const removed2 = index2 + 1
    const min = Math.min(removed1, removed2)
    return findMedianSortedArrays(nums1.slice(0, nums1.length - min), nums2.slice(min))
  } else if (nums2[index2] > (nums1[index1 + 1] ?? Number.MAX_SAFE_INTEGER)) {
    if (nums2.length === 1) {
      return nums1[index1 + 1]
    }
    const removed1 = index1 + 1
    const removed2 = nums2.length - (index2 + 1)
    const min = Math.min(removed1, removed2)
    return findMedianSortedArrays(nums1.slice(min), nums2.slice(0, nums2.length - min))
  } else {
    if (isOdd) {
      return Math.max(nums1[index1], nums2[index2])
    }
    return (Math.max(nums1[index1], nums2[index2] ?? Number.MIN_SAFE_INTEGER) + Math.min(nums1[index1 + 1] ?? Number.MAX_SAFE_INTEGER, nums2[index2 + 1] ?? Number.MAX_SAFE_INTEGER)) / 2
  }
};

const res = findMedianSortedArrays([2, 3, 4], [1])
console.log(res)