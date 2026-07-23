function myAtoi(s: string): number {
  let res = 0
  let leadingWhitespace = true
  let sign = 0
  let leadingZero = true
  let startRead = false
  for (let i = 0; i < s.length; i++) {
    const c = s[i]

    if (c !== ' ') {
      leadingWhitespace = false
    }
    if (leadingWhitespace && c === ' ') continue

    if (!startRead) {
      if (c === '-') {
        sign = -1
        startRead = true
        continue
      } else if (c === '+'){
        sign = 1
        startRead = true
        continue
      } else {
        sign = 1
      }
    }
    startRead = true
    
    if (c !== '0') {
      leadingZero  = false
    }
    if (leadingZero && c === '0') continue

    switch(c) {
      case '0':
        res = res * 10
        break
      case '1':
        res = res * 10 + 1
        break
      case '2':
        res = res * 10 + 2
        break
      case '3':
        res = res * 10 + 3
        break
      case '4':
        res = res * 10 + 4
        break
      case '5':
        res = res * 10 + 5
        break
      case '6':
        res = res * 10 + 6
        break
      case '7':
        res = res * 10 + 7
        break
      case '8':
        res = res * 10 + 8
        break
      case '9':
        res = res * 10 + 9
        break
      default: return res * sign
    }
    if (res * sign > Math.pow(2, 31) - 1) return Math.pow(2, 31) - 1
    else if (res * sign < -Math.pow(2, 31)) return -Math.pow(2, 31)
  }
  return res * sign
};
