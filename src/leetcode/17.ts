function letterCombinations(digits: string): string[] {
  const map = [['a', 'b', 'c'], ['d', 'e', 'f'], ['g', 'h', 'i'], ['j', 'k', 'l'],
  ['m', 'n', 'o'], ['p', 'q', 'r', 's'], ['t', 'u', 'v'], ['w', 'x', 'y', 'z']]

  const res: string[] = []
  function recurse(digits: string, prev: string) {
    const a = map[Number(digits[0]) - 2]
    if (digits.length === 1) {
      for (let i = 0; i < a.length; i++) {
        res.push(prev + a[i])
      }
    } else {
      const nextDigits = digits.slice(1)
      for (let i = 0; i < a.length; i++) {
        recurse(nextDigits, prev + a[i])
      }
    }
  }
  recurse(digits, '')
  return res
};
