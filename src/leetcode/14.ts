function longestCommonPrefix(strs: string[]): string {
  if (strs.length === 1) return strs[0]

  let j = 0
  while (true) {
    for (let i = 1; i < strs.length; i++) {
      if (strs[i][j] !== strs[i - 1][j] || strs[i][j] === undefined) return strs[i].slice(0, j)
    }
    j++
  }
};
