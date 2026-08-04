function minInitialStrength(monsters: number[], boosts: number[][]): number {
  const diff: number[] = Array(monsters.length).fill(0)
  for (const [start, end, val] of boosts) {
    diff[start - 1] -= val
    diff[end] += val
  }

  let sum = 0, bonus = 0
  for (let i = monsters.length - 1; i >= 0; i--) {
    const m = monsters[i]
    bonus += diff[i]

    if (sum > 0) {
      sum += m
    } else {
      sum += Math.max(m - bonus, 0)
    }

  }
  return sum
};

console.log(minInitialStrength([5, 10, 15], [[1, 2, 10], [1, 2, 5]]))