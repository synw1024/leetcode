function lengthOfLongestSubstring(s: string): number {
  const map: { [key: string]: number } = {}
  let max = 0
  for (let i = 0, start = 0; i < s.length && s.length - start > max; i++) {
    const c = s[i]
    if (map[c] === undefined || map[c] < start) {
      max = Math.max(max, i - start + 1)
    } else {
      start = map[c] + 1
    }
    map[c] = i
  }
  return max
};
