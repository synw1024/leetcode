function minInitialStrength(monsters: number[], boosts: number[][]): number {
  const bonus: { [key: number]: number } = {}

  function findBonus() {
    
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