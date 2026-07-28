function threeSum(nums: number[]): number[][] {
  const positiveMap: { [key: number]: number } = {}
  const negativeMap: { [key: number]: number } = {}
  const positiveKeys: number[] = []
  const negativeKeys: number[] = []
  let zeroCount = 0
  for (let i = 0; i < nums.length; i++) {
    const n = nums[i]
    if (n > 0) {
      if (positiveMap[n]) {
        positiveMap[n]++
      } else {
        positiveMap[n] = 1
        positiveKeys.push(n)
      }
    } else if (n < 0) {
      if (negativeMap[n]) {
        negativeMap[n]++
      } else {
        negativeMap[n] = 1
        negativeKeys.push(n)
      }
    } else {
      zeroCount++
    }
  }
  const res: number[][] = []

  if (zeroCount >= 3) {
    res.push([0, 0, 0])
  }
  if (zeroCount > 0) {
    const [short, long] = positiveKeys.length <= negativeKeys.length ? [positiveKeys, negativeMap] : [negativeKeys, positiveMap]
    for (let i = 0; i < short.length; i++) {
      if (long[-short[i]]) {
        res.push([0, short[i], -short[i]])
      }
    }
  }

  for (let i = 0; i < positiveKeys.length; i++) {
    const target = -positiveKeys[i]
    const usedKeys: number[] = []
    for (let j = 0; j < negativeKeys.length; j++) {
      const diff = target - negativeKeys[j]
      if (!usedKeys.includes(negativeKeys[j]) && negativeMap[diff] && (diff !== negativeKeys[j] || negativeMap[diff] > 1)) {
        usedKeys.push(diff)
        res.push([positiveKeys[i], negativeKeys[j], diff])
      }
    }
  }
  for (let i = 0; i < negativeKeys.length; i++) {
    const target = -negativeKeys[i]
    const usedKeys: number[] = []
    for (let j = 0; j < positiveKeys.length; j++) {
      const diff = target - positiveKeys[j]
      if (!usedKeys.includes(positiveKeys[j]) && positiveMap[diff] && (diff !== positiveKeys[j] || positiveMap[diff] > 1)) {
        usedKeys.push(diff)
        res.push([negativeKeys[i], positiveKeys[j], diff])
      }
    }
  }

  return res
};

threeSum([-100, -70, -60, 110, 120, 130, 160])