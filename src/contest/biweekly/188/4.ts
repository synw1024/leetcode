/**
 * 
 */
function minMaxWaitingTime(demand: number[], fuel: number[]): number {
  const cache: {[key: string]: number[]} = {}
  function recurse(n: number, f0: number, f1: number, w0: number, w1: number) {
    const key = n + ',' + f0 + ',' + f1 + ',' + w0 + ',' + w1
    if (cache[key]) return cache[key]

    let res = [n, 0]

    if (n === demand.length) {
      return [n, 0]
    }

    const d = demand[n]
    if (f0 - d >= 0) {
      const next = recurse(n + 1, f0 - d, f1, d, Math.max(0, w1 - w0))
      const cand = [next[0], Math.max(next[1], w0)]
      if (next[0] > res[0] || (next[0] === res[0] && cand[1] < res[1])) {
        res = cand
      }
    }

    if (f1 - d >= 0) {
      const next = recurse(n + 1, f0, f1 - d, Math.max(0, w0 - w1), d)
      const cand = [next[0], Math.max(next[1], w1)]
      if (next[0] > res[0] || (next[0] === res[0] && cand[1] < res[1])) {
        res = cand
      }
    }

    cache[key] = res

    return res
  }
  const res = recurse(0, fuel[0], fuel[1], 0, 0)
  return res[0] ? res[1] : -1
};

minMaxWaitingTime([3, 2, 4, 4], [4, 5])
