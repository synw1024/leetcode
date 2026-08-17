function maximumGap(skill: string, station: string): number {
  const left: number[] = []
  
  let max = 0, last = station.length
  for (let i = skill.length - 1; i > 0; i--) {
    last = station.lastIndexOf(skill[i], last-1)
    max = Math.max(max, last - getLeft(i-1))
  }

  function getLeft(i: number) {
    if (left[i] !== undefined) return left[i]
    if (i === 0) {
      left[i] = station.indexOf(skill[i])
    } else {
      left[i] = station.indexOf(skill[i], getLeft(i - 1) + 1)
    }
    return left[i]
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