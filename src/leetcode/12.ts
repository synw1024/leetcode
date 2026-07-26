function intToRoman(num: number): string {
  let res = ''
  const symbols = ['I', 'V', 'X', 'L', 'C', 'D', 'M']
  for (let i = 0; i < 4; i++) {
    const n = num % 10
    if (n < 4) {
      res = Array(n).fill(symbols[i * 2]).join('') + res
    } else if (n === 4) {
      res = `${symbols[i * 2]}${symbols[i * 2 + 1]}${res}`
    } else if (n === 5) {
      res = `${symbols[i * 2 + 1]}${res}`
    } else if (n < 9) {
      res = `${symbols[i * 2 + 1]}${Array(n - 5).fill(symbols[i * 2]).join('')}${res}`
    } else {
      res = `${symbols[i * 2]}${symbols[(i + 1) * 2]}${res}`
    }
    num = Math.floor(num / 10)
  }
  return res
};
