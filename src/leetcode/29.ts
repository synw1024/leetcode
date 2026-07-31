function divide(dividend: number, divisor: number): number {
  if (dividend === -Math.pow(2, 31) && divisor === -1) {
    return Math.pow(2, 31) - 1
  }

  if (Math.abs(dividend) < Math.abs(divisor) || dividend === 0) return 0

  let sign = 1
  if ((dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0)) {
    sign = -1
  }

  dividend = Math.abs(dividend)
  divisor = Math.abs(divisor)

  const s1 = dividend.toString()
  const s2 = divisor.toString()
  let res = ''
  function recurse(s1: string, s2: string, r: string) {
    console.log(s1, r)
    const n0 = Number(r + s1)
    let n2 = Number(s2)
    if (n0 < n2) {
      res += Array(s1.length).fill(0).join('')
      return
    }
    let n1 = Number(r + s1.slice(0, s2.length))
    if (n1 < n2) {
      res += '0'
      n1 = Number(s1.slice(0, s2.length + 1))
    }
    const [quotient, remainder] = division(n1, n2)
    res += quotient
    recurse(s1.slice(r ? String(n1).length - r.length : String(n1).length), s2, remainder ? String(remainder) : '')
  }

  recurse(s1, s2, '')


  return sign < 0 ? -Number(res) : Number(res)
};

function division(a: number, b: number) {
  let times = 0
  while (a >= b) {
    a -= b
    times++
  }
  return [times, a]
}

console.log(divide(-1060849722, 99958928))

