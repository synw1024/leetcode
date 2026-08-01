/**
 * fn(n) = const(n) + max()
 */
function maximumWidth(planks: number[]): number {
  const map: { [key: string]: number } = {}
  for (let i = 0; i < planks.length; i++) {
    if (map[planks[i]]) {
      map[planks[i]]++
    } else {
      map[planks[i]] = 1
    }
  }
  const heights = Object.keys(map)
  const countMap = {...map}

  for (let i = 0; i < heights.length; i++) {
    const n = Number(heights[i])
    for (let j = map[n] > 1 ? i : i + 1; j < heights.length; j++) {
      const m = Number(heights[j])
      if (j === i) {
        const min = Math.floor(map[n] / 2)
        if (countMap[2 * n]) {
          countMap[2 * n] += min
        } else {
          countMap[2 * n] = min
        }
        continue
      }

      const min = Math.min(map[n], map[m])
      if (countMap[n + m]) {
        countMap[n + m] += min
      } else {
        countMap[n + m] = min
      } 
    }
  }
  const values = Object.values(countMap)
  return Math.max(...values)
};

console.log(maximumWidth([1, 3, 2, 5, 7, 5, 4, 2, 1]))