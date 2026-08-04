function minInitialStrength(monsters: number[], boosts: number[][]): number {
  const diff: number[] = []
  for (let i = 0; i < boosts.length; i++) {
    const [start, end, val] = boosts[i]
    diff[start] = diff[start] ? diff[start] + val : val
    diff[end+1] = diff[end+1] ? diff[end+1] - val : -val
  }
  const bonus: number[] = []
  bonus[0] = diff[0] || 0
  for (let i = 1; i < monsters.length; i++) {
    bonus[i] = bonus[i - 1] + (diff[i] || 0)
  }

  let sum = 0
  for (let i = monsters.length - 1; i >= 0; i--) {
    const m = monsters[i]

    if (sum > 0 && i < monsters.length - 1) {
      sum += m
    }

    const diff = m - (bonus[i] || 0) - sum
    if (diff > 0) {
      sum += diff
    }

  }
  return sum
};

console.log(minInitialStrength([5, 10, 15], [[1, 2, 10], [1, 2, 5]]))