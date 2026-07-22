function reverse(x: number): number {
  let skipZero = true
  let res = 0
  if (x >= 0) {
    while (x) {
      const n = x % 10
      if (skipZero && n === 0) {
        x = Math.floor(x / 10)
        continue
      }

      skipZero = false
      res = res * 10 + n
      if (res > Math.pow(2, 31) - 1) return 0
      x = Math.floor(x / 10)
    }
  } else {
    while (x) {
      const n = x % 10
      if (skipZero && n === 0) {
        x = Math.ceil(x / 10)
        continue
      }

      skipZero = false
      res = res * 10 + n
      if (res < -Math.pow(2, 31)) return 0
      x = Math.ceil(x / 10)
    }
  }
  return res
};
