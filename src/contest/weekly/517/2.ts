function sumDecoded(nums: number[]): number {
  return nums.reduce((sum, cur) => {
    const width = cur % 10
    const d = Math.floor(cur / 10)
    const s = d.toString()
    const x = parseInt(s.slice(0, width))
    const y = parseInt(s.slice(width))

    let n = y, times = 1, remain = 0
    while (Math.pow(x, n) > Math.pow(Number.MAX_SAFE_INTEGER, 0.5)) {
      if (n % 2 > 0) remain += times
      n = Math.floor(n / 2)
      times *= 2
    }

    let c = Math.pow(x, remain) % 1000000007
    for (let i = 0; i < times; i++) {
      c = (c * Math.pow(x, n)) % 1000000007
    }

    return (sum + c) % 1000000007
  }, 0)
};

// console.log(sumDecoded([59412]))
console.log(sumDecoded([379723]))
