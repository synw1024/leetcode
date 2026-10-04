function minRotations(s: string): number {
  let point = 0, counter = 0
  for (let i = 0; i < s.length; i++) {
    const n = Number(s[i])
    const diff = Math.min(Math.abs(n - point), Math.min(n, point) + 10 - Math.max(n, point))
    counter += diff
    point = n
  }
  return counter
};

minRotations('0192837465')
