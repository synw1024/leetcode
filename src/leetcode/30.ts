function findSubstring(s: string, words: string[]): number[] {
  const wl = words[0].length
  const wordMap: { [key: string]: number } = {}
  for (let i = 0; i < words.length; i++) {
    if (wordMap[words[i]]) {
      wordMap[words[i]]++
    } else {
      wordMap[words[i]] = 1
    }
  }

  const res: number[] = []
  for (let h = 0; h < wl; h++) {
    let countMap: { [key: string]: number } = {}
    let count = 0
    for (let i = h; i <= s.length - (words.length - count); i += wl) {
      const sw = s.slice(i, i + wl)

      if (!wordMap[sw]) {
        countMap = {}
        count = 0
        continue
      }

      if (countMap[sw]) {
        countMap[sw]++
      } else {
        countMap[sw] = 1
      }
      count++
      let j = i - (count - 1) * wl
      while (countMap[sw] > wordMap[sw]) {
        const swj = s.slice(j, j + wl)
        countMap[swj]--
        count--
        j += wl
      }

      if (count < words.length) {
        continue
      }

      const start = i - wl * (count - 1)
      res.push(start)
      const startWord = s.slice(start, start + wl)
      countMap[startWord]--
      count--
    }
  }

  return res
};

console.log(findSubstring('barfoothefoobarman', ["foo","bar"]))