function elevatorRequests(n: number, start: number, requests: number[][]): number {
  const cache: {[key: string]: number} = {}
  function recurse(startTime: number, startFloor: number, remainingFloor: number) {
    const key = startTime + ',' + startFloor + ',' + remainingFloor
    if (cache[key]) return cache[key]

    if (!remainingFloor) return startTime

    const up: number[] = [], down: number[] = []
    for (let i = 0; i < requests.length; i++) {
      const val = 1 << i
      if (!(remainingFloor & val)) continue

      const r = requests[i]
      const walk = startFloor - r[1]
      const wait = r[0] - (startTime + Math.abs(walk))
      if (wait <= 0) {
        if (walk > 0) {
          down.push(i)
        } else if (walk < 0) {
          up.push(i)
        } else {
          remainingFloor = remainingFloor & ~val
        }
      }
    }

    let filterRemaining = remainingFloor

    const minUp = Math.min(...up.map(u => requests[u][1]))
    up.forEach(u => {
      if (requests[u][1] > minUp) {
        const val = 1 << u
        filterRemaining = filterRemaining & ~val
      }
    })

    const maxDown = Math.max(...down.map(d => requests[d][1]))
    down.forEach(d => {
      if (requests[d][1] < maxDown) {
        const val = 1 << d
        filterRemaining = filterRemaining & ~val
      }
    })

    let min = Number.MAX_SAFE_INTEGER
    for (let i = 0; i < requests.length; i++) {
      const val = 1 << i
      if (!(filterRemaining & val)) continue

      const r = requests[i]
      const newRemaining = remainingFloor & ~val
      const res = recurse(Math.max(startTime + floorDiff(startFloor, r[1]), r[0]), r[1], newRemaining)
      min = Math.min(res, min)
    }

    cache[key] = min === Number.MAX_SAFE_INTEGER ? startTime : min
    return cache[key]
  }
  return recurse(0, start, (1 << requests.length) - 1)
};

function floorDiff(a: number, b: number) {
  return Math.abs(a - b)
}

console.log(elevatorRequests(9, 0, [[0, 8], [6, 5]]))
