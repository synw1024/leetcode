/**
 * start(n) = start(n - 1) + demand(n - 1) || max(anotherQueue.lastStart + anotherQueue.lastDemand, start(n - 1))
 * wait(n) = demand(n - 1) || max(anotherQueue.lastStart + anotherQueue.lastDemand, start(n - 1)) - start(n - 1)
 * left, right: [[start, demandIndex]]
 */
function minMaxWaitingTime(demand: number[], fuel: number[]): number {
  let minimumWait: number[] = []
  function recurse(n: number, left: number[][], right: number[][], wait: number[]) {
    const d = demand[n]
    const totalLeft = left.reduce((prev, [_, index]) => prev + demand[index], 0)
    const totalRight = right.reduce((prev, [_, index]) => prev + demand[index], 0)
    const totalWait = wait.reduce((prev, cur) => prev + cur, 0)
    let totalminimumWait = minimumWait.reduce((prev, cur) => prev + cur, 0)

    if (n === demand.length - 1) {
      if (left[left.length - 1][1] === n - 1) {
        if (totalLeft + d <= fuel[0] && (totalWait + demand[n - 1] < totalminimumWait || wait.length + 1 > minimumWait.length)) {
          minimumWait = [...wait, demand[n - 1]]
        }

        const [lastRightStart, lastRightDemandIndex] = right[right.length - 1]
        const lastLeftStart = left[left.length - 1][0]
        const anotherQueueWait = lastRightStart + demand[lastRightDemandIndex] - lastLeftStart
        totalminimumWait = minimumWait.reduce((prev, cur) => prev + cur, 0)
        if (totalRight + d <= fuel[1] && (totalWait + anotherQueueWait < totalminimumWait || wait.length + 1 > minimumWait.length)) {
          minimumWait = [...wait, anotherQueueWait]
        }
      } else {
        if (totalRight + d <= fuel[1] && (totalWait + demand[n - 1] < totalminimumWait || wait.length + 1 > minimumWait.length)) {
          minimumWait = [...wait, demand[n - 1]]
        }

        const [lastLeftStart, lastLeftDemandIndex] = left[left.length - 1]
        const lastRightStart = right[right.length - 1][0]
        const anotherQueueWait = lastLeftStart + demand[lastLeftDemandIndex] - lastRightStart
        totalminimumWait = minimumWait.reduce((prev, cur) => prev + cur, 0)
        if (totalLeft + d <= fuel[0] && (totalWait + anotherQueueWait < totalminimumWait || wait.length + 1 > minimumWait.length)) {
          minimumWait = [...wait, anotherQueueWait]
        }
      }
      return
    }

    if (totalLeft + d > fuel[0] && totalRight + d > fuel[1]) {
      if (wait.length > minimumWait.length) {
        minimumWait = [...wait]
      }
      return
    }

    if (totalLeft + d <= fuel[0]) {
      if (left[left.length - 1][1] === n - 1) {
        const [lastLeftStart, lastLeftDemandIndex] = left[left.length - 1]
        const start = lastLeftStart + demand[lastLeftDemandIndex]
        recurse(n + 1, [...left, [start, n]], [...right], [...wait, demand[lastLeftDemandIndex]])
      } else {
        const [lastLeftStart] = left[left.length - 1]
        const [lastRightStart, lastRightDemandIndex] = right[right.length - 1]
        const start = lastRightStart + demand[lastRightDemandIndex]
        recurse(n + 1, [...left], [...right, [start, n]], [...wait, start - lastLeftStart])
      }
    }

    if (totalRight + d <= fuel[1]) {
      if (right[right.length - 1][1] === n - 1) {
        const [lastRightStart, lastRightDemandIndex] = right[right.length - 1]
        const start = lastRightStart + demand[lastRightDemandIndex]
        recurse(n + 1, [...left], [...right, [start, n]], [...wait, demand[lastRightDemandIndex]])
      } else {
        const [lastRightStart] = right[right.length - 1]
        const [lastLeftStart, lastLeftDemandIndex] = left[left.length - 1]
        const start = lastLeftStart + demand[lastLeftDemandIndex]
        recurse(n + 1, [...left, [start, n]], [...right], [...wait, start - lastRightStart])
      }
    }
  }
  recurse(0, [], [])
};
