function findSubstring(s: string, words: string[]): number[] {
  const wl = words[0].length
  const res: number[] = []
  for (let h = 0; h < wl; h++) {
    let tempWords = [...words]
    for (let i = h; i <= s.length - wl * tempWords.length;) {
      const sw = s.slice(i, i + wl)
      const existingInWords = words.includes(sw)

      if (!existingInWords) {
        i += wl
        tempWords = [...words]
        continue
      }

      const index = tempWords.indexOf(sw)
      if (index === -1) {
        const removed = words.length - tempWords.length
        let j = i - removed * wl
        while (j < i) {
          const swj = s.slice(j, j + wl)
          tempWords.push(swj)
          j += wl
          if (swj === sw) {
            break
          }
        }
      } else {
        tempWords.splice(index, 1)
        if (!tempWords.length) {
          const start = i - wl * (words.length - 1)
          res.push(start)
          const startWord = s.slice(start, start + wl)
          tempWords.push(startWord)
        }
        i += wl
      }
    }
  }

  return res
};