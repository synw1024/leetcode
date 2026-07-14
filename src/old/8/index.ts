function myAtoi(s: string): number {
  const automaton = [
    [0, 1, 2, 3],
    [3, 3, 2, 3],
    [3, 3, 2, 3],
    [3, 3, 3, 3],
  ]
  let start = -1
  let end = -1
  let state = 0
  let isNegative = false
  for (let i = 0; i < s.length; i++) {
    const c = getCol(s[i])
    state = automaton[state][c]
    if (state === 1) {
      if (s[i] === '-') isNegative = true
    } else if (state === 2 && start === -1) {
      start = i
    } else if (state === 3) {
      end = i
      break
    }
  }
  if (end === -1) end = s.length
  if (start === -1) s = ''
  let n = Number(s.slice(start, end))

  if (isNegative) n *= -1

  const min = -Math.pow(2, 31)
  if (n < min) n = min

  const max = Math.pow(2, 31) - 1
  if (n > max) n = max

  return n
}

function getCol(c: string) {
  if (c === ' ') {
    return 0
  } else if (['-', '+'].includes(c)) {
    return 1
  } else if (/\d/.test(c)) {
    return 2
  } else {
    return 3
  }
}
