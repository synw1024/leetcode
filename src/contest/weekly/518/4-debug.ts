function minCost(grid: number[][], k: number): number {
  const m = grid.length
  const n = grid[0].length

  const cache: { [key: string]: (number | string)[] } = {}
  function recurse(i: number, j: number, visited: string, direction: number, turn: number) {
    if (i === m - 1 && j === n - 1) return [grid[i][j], `[${i},${j}]`]

    visited = newVisited(visited, i * n + j)
    const key = i + ',' + j + ',' + direction + ',' + turn
    if (cache[key] !== undefined) return cache[key]

    let min = Number.MAX_SAFE_INTEGER
    let stack = ''

    // if (i === 3 && j === 2 && direction === 2 && visited.startsWith('110000000100000001000000011')) {
    //   debugger
    // }

    if (i - 1 >= 0 && !checkVisited(visited, i - 1, j)) {
      const newTurn = direction === 1 || !direction ? turn : turn - 1
      if (newTurn >= 0) {
        const [next, lastStack] = recurse(i - 1, j, visited, 1, newTurn)
        if (next !== -1) {
          min = Math.min(min, next)
          if (min === next) {
            stack = lastStack
          }
        }
      }
    }
    if (j + 1 < n && !checkVisited(visited, i, j + 1)) {
      const newTurn = direction === 2 || !direction ? turn : turn - 1
      if (newTurn >= 0) {
        const [next, lastStack] = recurse(i, j + 1, visited, 2, newTurn)
        if (next !== -1) {
          min = Math.min(min, next)
          if (min === next) {
            stack = lastStack
          }
        }
      }
    }
    if (i + 1 < m && !checkVisited(visited, i + 1, j)) {
      const newTurn = direction === 3 || !direction ? turn : turn - 1
      if (newTurn >= 0) {
        const [next, lastStack] = recurse(i + 1, j, visited, 3, newTurn)
        if (next !== -1) {
          min = Math.min(min, next)
          if (min === next) {
            stack = lastStack
          }
        }
      }
    }
    if (j - 1 >= 0 && !checkVisited(visited, i, j - 1)) {
      const newTurn = direction === 4 || !direction ? turn : turn - 1
      if (newTurn >= 0) {
        const [next, lastStack] = recurse(i, j - 1, visited, 4, newTurn)
        if (next !== -1) {
          min = Math.min(min, next)
          if (min === next) {
            stack = lastStack
          }
        }
      }
    }

    cache[key] = [min === Number.MAX_SAFE_INTEGER ? -1 : min + grid[i][j], `[${i},${j}] - ` + stack]

    return cache[key]
  }

  function checkVisited(visited: string, i: number, j: number) {
    const res = visited[i * n + j] === '1' && false
    return res
  }

  function newVisited(visited: string, index: number) {
    const prev = visited.slice(0, index)
    const post = visited.slice(index + 1)
    return prev + '1' + post
  }

  const [res, stack] = recurse(0, 0, Array(m * n).fill(0).join(''), 0, k)

  console.log(res, stack)
};

// const p1 = [[20, 53, 26, 55, 36, 12, 6, 38], [0, 54, 4, 6, 69, 62, 61, 1], [56, 23, 63, 33, 63, 15, 68, 26], [32, 74, 31, 61, 46, 75, 6, 31], [11, 60, 73, 58, 34, 2, 64, 20], [19, 41, 41, 63, 61, 60, 14, 4], [5, 67, 59, 60, 58, 0, 47, 31]]
// const p1 = [[2,7,3],[1,4,5]]
const p1 = [[39,53,65,3,45,72,25,3],[64,48,38,64,17,24,24,53],[73,30,69,33,58,40,28,74],[76,8,6,6,28,61,72,17],[57,22,22,65,40,56,4,55],[75,15,7,25,62,20,57,37],[15,2,48,54,25,61,49,28]]

minCost(p1, 6)

/**
 * 438 [0,0] - [0,1] - [1,1] - [2,1] - [3,1] - [3,2] - [3,3] - [3,4] - [4,4] - [4,5] - [4,6] - [4,7] - [5,7] - [6,7]
 * 442 [0,0] - [0,1] - [1,1] - [2,1] - [3,1] - [3,2] - [4,2] - [5,2] - [5,3] - [5,4] - [5,5] - [5,6] - [5,7] - [6,7]
 * 
 * 11111111
 * 01111111
 * 01000011
 * 01100011000000110000001100000000
 * 
 * 11000000
 * 01000000
 * 01000000
 * 011
 */
