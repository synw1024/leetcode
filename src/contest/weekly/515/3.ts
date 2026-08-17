function maximumGap(skill: string, station: string): number {
  const left: number[] = []
  for (let i = 0; i < skill.length; i++) {
    left[i] = station.indexOf(skill[i], (left[i-1] ?? -1) + 1)
  }
  
  let max = 0, last = station.length
  for (let i = skill.length - 1; i > 0; i--) {
    last = station.lastIndexOf(skill[i], last-1)
    max = Math.max(max, last - left[i-1])
  }
  return max
};

// console.log(maximumGap('aa', 'aaaa'))
// console.log(maximumGap('xyz', 'xyzz'))
// console.log(maximumGap('cbc', 'cbcdbc'))
// console.log(maximumGap('adkz', 'adykroez'))
// console.log(maximumGap('cc', 'acc'))
// console.log(maximumGap('acc', 'cacc'))
console.log(maximumGap('aaa', 'aaa'))