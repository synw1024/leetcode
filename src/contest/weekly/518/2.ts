function countGoodRotations(nums: number[]): number {
  const prevSums = nums.reduce((res, cur) => {
    if (res.length === 0) {
      res.push(cur)
    } else {
      res.push(res[res.length - 1] + cur)
    }
    return res
  }, [] as number[])

  const half = nums.length / 2
  const increse = half - 1
  let res = 0
  for (let i = 0; i < nums.length; i++) {
    const prev = ((prevSums[i + increse] || prevSums[prevSums.length-1]) - (prevSums[i - 1] || 0)) + (prevSums[i - half - 1] || 0)
    const post = prevSums[prevSums.length - 1] - prev
    if (prev > post) {
      res++
    }
  }
  return res
};

console.log(countGoodRotations([1, 2, 3, 4, 5, 6]))
