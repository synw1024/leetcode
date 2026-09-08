function minCost(grid: number[][], k: number): number {
  const m = grid.length
  const n = grid[0].length
  const max = 1000 * m * n

  const cache: { [key: string]: number } = {}
  const dirs = [[0, 1], [1, 0], [0, -1], [-1, 0]]
  function recurse(i: number, j: number, direction: number, turn: number): number {
    if (i === m - 1 && j === n - 1) return grid[i][j]

    const key = i + ',' + j + ',' + direction + ',' + turn
    if (cache[key] !== undefined) return cache[key]

    let min = max
    for (let d = 0; d < dirs.length; d++) {
      const [x, y] = [i + dirs[d][0], j + dirs[d][1]]
      if (x < 0 || x >= m || y < 0 || y >= n || (direction !== -1 && (d + 2) % 4 === direction)) continue

      const t = d === direction || direction === -1 ? turn : turn - 1
      if (t < 0) continue

      min = Math.min(min, recurse(x, y, d, t) + grid[i][j])
      if (!min) break
    }

    return cache[key] = min
  }

  const res = recurse(0, 0, -1, k)
  return res >= max ? -1 : res
};

console.log(minCost([[0,0,0],[0,1,0],[0,0,0]], 2))
