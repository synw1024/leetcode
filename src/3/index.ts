function lengthOfLongestSubstring(s: string): number {
  const map: {
    [key: string]: number
  } = {}
  let max = 0
  let start = 0
  for (let i = 0; i < s.length; i++) {
    const c = s[i]
    if (map[c] !== undefined && map[c] >= start) {
      // 存在
      const len = i - start
      if (len > max) max = len
      start = map[c] + 1
    }
    map[c] = i
  }
  const len = s.length - start
  return max > len ? max : len
};
