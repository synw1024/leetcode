function minRotations(n: number, s: string): number {
  let leftPrefixsum = [0], counter = 0, point = 0
  for (let i = 0; i < s.length; i++) {
    const n = Number(s[i])
    const diff = Math.min(Math.abs(n - point), Math.min(n, point) + 10 - Math.max(n, point))
    counter += diff
    leftPrefixsum.push(counter)
    point = n
  }
  let rightPrefixsum = [0, 0]
  counter = 0, point = Number(s[s.length - 1])
  for (let i = s.length - 2; i >= 0; i--) {
    const n = Number(s[i])
    const diff = Math.min(Math.abs(n - point), Math.min(n, point) + 10 - Math.max(n, point))
    counter += diff
    rightPrefixsum.push(counter)
    point = n
  }
  let min = Number.MAX_SAFE_INTEGER
  for (let i = 0; i <= s.length; i++) {
    const n = Number(s[i-1] || 0)
    const last = Number(s[s.length - 1])
    const diff = Math.min(Math.abs(n - last), Math.min(n, last) + 10 - Math.max(n, last))
    min = Math.min(rightPrefixsum[s.length - i] + diff + leftPrefixsum[i], min)
  }
  return min
}

console.log(minRotations(4, '2916'))
