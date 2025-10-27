function lengthOfLIS(nums: number[]): number {
  let longest = 0
  const map: {
    [index: number]: number
  } = {}
  function optimization(n: number) {
    if (map[n]) return map[n]

    const slice = nums.slice(0, n)
    let localLongest = 0
    for (let i = slice.length - 1; i >= 0; i--) {
      if (nums[i] < nums[n]) {
        const res = optimization(i)
        if (localLongest < res) {
          localLongest = res
        }
      }
    }
    map[n] = localLongest + 1
    return map[n]
  }
  for (let i = nums.length - 1; i >= 0; i--) {
    const res = optimization(i)
    if (res > longest) longest = res
  }
  return longest
};
