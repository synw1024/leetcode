function countValidPrefixes(s: string): number {
  const map: {[key: string]: number} = {
    '0': 0,
    '1': 0
  }
  let count = 0
  for (let i = 0; i < s.length; i++) {
    map[s[i]]++
    const diff = Math.abs(map['1'] - map['0'])
    if (diff <= 1) {
      count++
    }
  }
  return count
};
