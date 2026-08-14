function countTasks(tasks: number[], shifts: number[]): number[] {
  const sums = tasks.reduce((res, t) => {
    res.push(t + (res[res.length - 1] || 0))
    return res
  }, [] as number[])
  let nextStart = 0, remaining = 0, res: number[] = []
  for (let i = 0; i < shifts.length; i++) {
    if (shifts[i] < remaining) {
      remaining = remaining - shifts[i]
      res.push(tasks.length - nextStart)
      continue
    } else if (shifts[i] === remaining) {
      remaining = 0
      nextStart++
      res.push(tasks.length - nextStart)
      nextStart %= tasks.length
      continue
    }
    const start = remaining ? nextStart + 1 : nextStart;
    if (start >= tasks.length) {
      res.push(0)
      nextStart = 0
      remaining = 0
      continue
    }
    [nextStart, remaining] = recurse(shifts[i] - remaining + (sums[start-1] || 0), start, tasks.length - 1)
    res.push(tasks.length - nextStart)
    if (nextStart === tasks.length) {
      nextStart = 0
    }
  }
  function recurse(shift: number, start: number, end: number): number[] {
    if (start === end) {
      if (sums[start] <= shift) {
        return [start + 1, 0]
      }
      return [start, sums[start] - shift]
    }

    const mid = Math.floor((end - start) / 2) + start
    if (sums[mid] < shift) {
      return recurse(shift, mid + 1, end)
    } else if (sums[mid] > shift && sums[mid] - shift >= tasks[mid]) {
      return recurse(shift, start, mid - 1)
    } else if (sums[mid] > shift) {
      return [mid, sums[mid] - shift]
    } else {
      return [mid + 1, 0]
    }
  }
  return res
};

// console.log(countTasks([1,4,4], [9,1,4]))
// console.log(countTasks([2,3,4], [20,4,5]))
// console.log(countTasks([4,2], [3,6,1]))
console.log(countTasks([1,1,3,3,8], [2,9,5,3,9]))
