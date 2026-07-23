/**
 * f(i, j) = f(i - 1, j - *) && match(i, j)
 */
function isMatch(s: string, p: string): boolean {
  let i = s.length - 1, j = p.length - 1

  function match(i: number, j: number) {
    if (p[j] === '.' || s[i] === p[j]) return true

    if (p[j] === '*' && (p[j-1] === '.' || p[j-1] === s[i])) return true

    return false
  }

  function recursion(i: number, j: number): boolean {
    if (i === 0 && j === 0) return match(i, j)

    if (i < 0 && j < 0) return true

    if (i < 0 && j > 0) return p[j] === '*' ? recursion(i, j - 2) : false

    if (i > 0 && j <= 0) return false

    if (p[j] !== '*') return match(i, j) && recursion(i - 1, j - 1)

    return (match(i, j) && recursion(i - 1, j)) || recursion(i, j - 2)
  }

  return recursion(i, j)
};
