function sumDecoded(nums: number[]): number {
  const mod = 1000000007n
  const res = nums.reduce((sum, cur) => {
    const width = cur % 10
    const d = Math.floor(cur / 10)
    const s = d.toString()
    let x = BigInt(parseInt(s.slice(0, width)))
    let y = parseInt(s.slice(width))

    let res = 1n
    while (y > 0) {
      if (y & 1) {
        res = (res * x) % mod
      }
      x = x * x % mod
      y = y >> 1
    }
    return (sum + res) % mod
  }, 0n)
  return Number(res)
};

// console.log(sumDecoded([59412]))
// console.log(sumDecoded([379723]))
console.log(sumDecoded([17752335231789]))
