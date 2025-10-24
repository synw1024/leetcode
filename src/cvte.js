function solve(strArr) {
  let s = ''
  for (let j = 0; true; j++) {
    for (let i = 0; i < strArr.length-1; i++) {
      if (!strArr[i][j] || !strArr[i+1][j]) {
        return s
      }
      if (strArr[i][j] !== strArr[i+1][j]) {
        return s
      }

      if (i === strArr.length - 2) {
        s += strArr[i][j]
      }
    }
  }
}

const res = solve(['flower', 'flow', 'alight', 'f', 'false'])
console.log(res)