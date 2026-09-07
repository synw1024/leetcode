function minCost(grid: number[][], k: number): number {
  const m = grid.length
  const n = grid[0].length

  const cache: { [key: string]: number } = {}
  function recurse(i: number, j: number, direction: number, turn: number): number {
    if (i === m - 1 && j === n - 1) return grid[i][j]

    const key = i + ',' + j + ',' + direction + ',' + turn
    if (cache[key] !== undefined) return cache[key]

    let min = Number.MAX_SAFE_INTEGER
    if (i - 1 >= 0) {
      const newTurn = direction === 1 || !direction ? turn : turn - 1
      if (newTurn >= 0) {
        const next = recurse(i - 1, j, 1, newTurn)
        if (next !== -1) {
          min = Math.min(min, next)
        }
      }
    }
    if (j + 1 < n) {
      const newTurn = direction === 2 || !direction ? turn : turn - 1
      if (newTurn >= 0) {
        const next = recurse(i, j + 1, 2, newTurn)
        if (next !== -1) {
          min = Math.min(min, next)
        }
      }
    }
    if (i + 1 < m) {
      const newTurn = direction === 3 || !direction ? turn : turn - 1
      if (newTurn >= 0) {
        const next = recurse(i + 1, j, 3, newTurn)
        if (next !== -1) {
          min = Math.min(min, next)
        }
      }
    }
    if (j - 1 >= 0) {
      const newTurn = direction === 4 || !direction ? turn : turn - 1
      if (newTurn >= 0) {
        const next = recurse(i, j - 1, 4, newTurn)
        if (next !== -1) {
          min = Math.min(min, next)
        }
      }
    }

    cache[key] = min === Number.MAX_SAFE_INTEGER ? -1 : min + grid[i][j]

    return cache[key]
  }

  return recurse(0, 0, 0, k)
};

console.log(minCost([[39,53,65,3,45,72,25,3],[64,48,38,64,17,24,24,53],[73,30,69,33,58,40,28,74],[76,8,6,6,28,61,72,17],[57,22,22,65,40,56,4,55],[75,15,7,25,62,20,57,37],[15,2,48,54,25,61,49,28]], 6))
