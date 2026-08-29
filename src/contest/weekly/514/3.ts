function maxArea(mat: number[][]): number {
  const m = mat.length, n = mat[0].length
  let maxSide = Math.floor(Math.max(m, n) / 2)
  const prevSides: number[][] = mat.reduce((res) => {
    res.push([])
    return res
  }, [] as number[][])
  const sides: { i: number, j: number }[][] = []
  for (let i = m - 1; i >= 0; i--) {
    for (let j = n - 1; j >= 0; j--) {
      if (!mat[i][j] || !prevSides[i + 1]?.[j + 1]) {
        prevSides[i][j] = mat[i][j]
        if (sides[prevSides[i][j]]) {
          sides[prevSides[i][j]].push({ i, j })
        } else {
          sides[prevSides[i][j]] = [{ i, j }]
        }
        continue
      }
      const side = Math.min(prevSides[i + 1][j + 1], prevSides[i + 1][j], prevSides[i][j + 1]) + 1
      prevSides[i][j] = side
      if (sides[side]) {
        sides[side].push({ i, j })
      } else {
        sides[side] = [{ i, j }]
      }
    }
  }

  let minSide = 1
  maxSide = Math.min(sides.length - 1, maxSide)
  let res = 0
  while (minSide <= maxSide) {
    const mid = minSide + maxSide >> 1
    let minR = m, maxR = -1
    let minC = n, maxC = -1
    for (let j = mid; j <= maxSide; j++) {
      for (let i = 0; i < sides[j].length; i++) {
        const item = sides[mid][i]
        if (item.i < minR) minR = item.i
        if (item.i > maxR) maxR = item.i
        if (item.j < minC) minC = item.j
        if (item.j > maxC) maxC = item.j
      }
    }

    if (maxR !== -1 && (maxR - minR >= mid || maxC - minC >= mid)) {
      res = mid
      minSide = mid + 1
    } else {
      maxSide = mid - 1
    }
  }

  return res * res
};

// console.log(maxArea([[0,1],[1,0]]))
// console.log(maxArea([[0,0],[0,1]]))
// console.log(maxArea([[0,1,1,0]]))
// console.log(maxArea([[1, 1, 1, 1], [1, 1, 1, 1]]))
// console.log(maxArea([[0,1,1,1,1,1,1,0],[1,1,1,1,0,1,0,1],[1,1,0,0,1,1,1,1]]))
console.log(maxArea([[0, 1, 1, 1, 1, 1], [1, 1, 1, 1, 1, 0], [1, 0, 1, 1, 1, 1]]))
