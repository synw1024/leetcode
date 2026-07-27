function romanToInt(s: string): number {
  let res = 0
  for (let i = 0; i < s.length; i++) {
    const c = s[i]
    if (c === 'I' && s[i + 1] !== 'V' && s[i + 1] !== 'X') {
      res += 1
    } else if (c === 'I') {
      res -= 1
    } else if (c === 'V') {
      res += 5
    } else if (c === 'X' && s[i + 1] !== 'L' && s[i + 1] !== 'C') {
      res += 10
    } else if (c === 'X') {
      res -= 10
    } else if (c === 'L') {
      res += 50
    } else if (c === 'C' && s[i + 1] !== 'D' && s[i + 1] !== 'M') {
      res += 100
    } else if (c === 'C') {
      res -= 100
    } else if (c === 'D') {
      res += 500
    } else {
      res += 1000
    }
  }
  return res
};
