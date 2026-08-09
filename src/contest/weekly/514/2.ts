function weightedSum(parent: number[], nums: number[]): number {
  const deepMap: {[key: number | string]: number} = {0: 1}
  for (let i = 1; i < parent.length; i++) {
    if (deepMap[i]) continue
    deepMap[i] = recurse(i)
  }
  function recurse(i: number) {
    const p = parent[i]
    if (deepMap[p]) return deepMap[p] + 1
    deepMap[p] = recurse(p)
    return deepMap[p] + 1
  }

  const h = Math.max(...Object.values(deepMap))
  const keys = Object.keys(deepMap)
  return keys.reduce((sum, key) => {
    const d = deepMap[key]
    return sum + nums[Number(key)] * (h - d + 1)
  }, 0)
}
