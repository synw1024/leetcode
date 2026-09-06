function countRotations(s: string, k: number): number {
  let res = 0
  for (let i = 0; i < s.length; i++) {
    const ss = s.slice(i) + s.slice(0, i)
    let last = ss[0]
    let kk = 0
    for (let j = 1; j < ss.length; j++) {
      if (ss[j] === last) {
        kk++
      } else {
        last = ss[j]
      }
    }
    if (kk === k) {
      res++
    }
  }
  return res
};

console.log(countRotations('aab', 1))
