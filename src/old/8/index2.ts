function myAtoi(s: string): number {
  s = s.trimStart()

  let symbol = 0
  let start = -1
  let end = -1
  let metZero = false
  for (let i = 0; i < s.length; i++) {
    const c = s[i]

    if (c === '-' && !symbol && start === -1 && !metZero) {
      symbol = -1
    } else if (c === '-') {
      end = i
      break
    } else if (c === '+' && !symbol && start === -1 && !metZero) {
      symbol = 1
    } else if (c === '+') {
      end = i
      break
    } else if (c === '0' && start < 0) {
      metZero = true
    } else if (/\d/.test(c) && start < 0) {
      start = i
    } else if (/[^\d]/.test(c)) {
      end = i
      break
    }
  }
  if (end === -1) end = s.length

  if (start === -1) s = ''

  s = s.slice(start, end)
  console.log(s)
  let n = 0
  for (let i = 0; i < s.length; i++) {
    n += Number(s[i]) * Math.pow(10, s.length - i - 1)
  }

  if (!symbol) symbol = 1
  n *= symbol

  const min = -Math.pow(2, 31)
  if (n < min) n = min

  const max = Math.pow(2, 31) - 1
  if (n > max) n = max

  return n
};
