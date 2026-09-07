function minCost(grid: number[][], k: number): number {
  const m = grid.length
  const n = grid[0].length

  const cache: { [key: string]: (number | string)[] } = {}
  function recurse(i: number, j: number, visited: string, direction: number, turn: number) {
    if (i === m - 1 && j === n - 1) return [grid[i][j], `[${i},${j}]`]

    visited = newVisited(visited, i * n + j)
    const key = i + ',' + j + ',' + direction + ',' + turn

    // if (i === 3 && j === 4 && visited.startsWith('11100000000100000000100000000111')) {
    //   debugger
    // }

    if (cache[key] !== undefined) return cache[key]

    let min = Number.MAX_SAFE_INTEGER
    let stack = ''
    let d = 0

    // if (i === 3 && j === 2 && visited.startsWith('110000000010000000010000000011')) {
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
            d = 1
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
            d = 2
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
            d = 3
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
            d = 4
          }
        }
      }
    }

    if (key === '6,6,3,2') {
      debugger
    }

    if (min === Number.MAX_SAFE_INTEGER) {
      return [-1, '']
    } else {
      return cache[key] = [min + grid[i][j], `[${i},${j}](${d},${min + grid[i][j]}) - ` + stack]
    }
  }

  function checkVisited(visited: string, i: number, j: number) {
    const res = visited[i * n + j] === '1'
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

const p1 = [[72,13,80,20,35,45,21,29,43],[81,9,87,50,10,37,61,15,49],[2,45,75,5,74,82,26,21,65],[70,15,55,26,47,15,92,5,80],[86,71,7,9,87,33,82,5,20],[64,56,55,83,57,36,1,42,15],[55,29,71,85,76,63,18,1,68],[50,4,57,78,12,78,22,68,61]]

minCost(p1, 7)

/**
 * 515 [0,0](2,515) - [0,1](3,443) - [1,1](3,430) - [2,1](3,421) - [3,1](2,376) - [3,2](2,361) - 
 * [3,3](2,306) - [3,4](2,280) - [3,5](3,233) - [4,5](3,218) - [5,5](2,185) - [5,6](3,149) - [6,6](2,148) - [6,7](2,130) - [6,8](3,129) - [7,8]
 * 
 * [3,3](2,354) - [3,4](2,328) - [3,5](2,281) - [3,6](2,266) - [3,7](3,174) - [4,7](2,169) - [4,8](3,164) - [5,8](3,144) - [6,8](3,129) - [7,8]
 * 
 *                [3,4](2,301) - [3,5](3,254) - [4,5](3,239) - [5,5](2,206) - [5,6](3,170) - [6,6](3,169) - [7,6](2,151) - [7,7](2,129) - [7,8]
 * 
 *                                                                                           [6,6](3,169) - [7,6](2,151) - [7,7](2,129) - [7,8]
 * 
 * 110000000
 * 010000000
 * 010000000
 * 011111000
 * 000001000
 * 000001100
 * 000000111
 * 000000001
 * 
 * 110000000010000000010000000011111000000001000000001100000000111
 * 
 * 
 * 
 * 111111111
 * 000000111
 * 000000111
 * 000000111
 * 000000111
 * 000000111
 * 000000111
 * 000000000
 * 
 * 
 */
