function convert(s: string, numRows: number): string {
  if (numRows === 1) return s
  let res = ''
  for (let row = 0; row < numRows; row++) {
    let v = true
    if (s[row]) {
      res += s[row]
    } else {
      break
    }
    if (row === numRows - 1) {
      v = false
    }
    let i = row
    while (true) {
      if (v) {
        i += (numRows - (row + 1)) * 2
      } else {
        i += (row) * 2
      }
      if (!s[i]) {
        break
      } else {
        res += s[i]
      }
      if (row !== 0 && row !== numRows - 1) {
        v = !v
      }

    }
  }
  return res
};

convert('PAYPALISHIRING', 3)