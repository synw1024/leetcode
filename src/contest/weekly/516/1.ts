function isPalindromic(s: string): boolean {
  let n = ''
  for (let i = 0; i < s.length; i++) {
    n += '0' + s[i].charCodeAt(0).toString(2)
  }
  let i = 0, j = n.length - 1
  while(i < j) {
    if (n[i] !== n[j]) return false
    i++
    j--
  }
  return true
};
