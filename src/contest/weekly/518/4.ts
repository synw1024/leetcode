function minCost(grid: number[][], k: number): number {
  const m = grid.length
  const n = grid[0].length

  const cache: number[][][][] = []
  for (let i = 0; i < m; i++) {
    const aa: number[][][] = []
    for (let j = 0; j < n; j++) {
      const a: number[][] = []
      for (let l = 0; l < 4; l++) {
        a.push(Array(k+1).fill(-1))
      }
      aa.push(a)
    }
    cache.push(aa)
  }

  const dirs = [[0, 1], [1, 0], [0, -1], [-1, 0]]
  function recurse(i: number, j: number, direction: number, turn: number): number {
    if (i === m - 1 && j === n - 1) return grid[i][j]

    if (direction !== -1 && cache[i][j][direction][turn] !== -1) return cache[i][j][direction][turn]

    let min = Number.MAX_SAFE_INTEGER
    for (let d = 0; d < dirs.length; d++) {
      const [x, y] = [i + dirs[d][0], j + dirs[d][1]]
      if (x < 0 || x >= m || y < 0 || y >= n || (direction !== -1 && (d + 2) % 4 === direction)) continue

      const t = d === direction || direction === -1 ? turn : turn + 1
      if (t > k) continue

      min = Math.min(min, recurse(x, y, d, t) + grid[i][j])
      if (!min) break
    }

    if (direction === -1) {
      return min
    } else {
      return cache[i][j][direction][turn] = min
    }
  }

  const res = recurse(0, 0, -1, 0)
  return res >= Number.MAX_SAFE_INTEGER ? -1 : res
};

console.log(minCost([[2, 7, 3], [1, 4, 5]], 1))
