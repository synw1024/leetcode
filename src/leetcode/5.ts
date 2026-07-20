function longestPalindrome(s: string): string {
  let start = 0
  let end = 0
  for (let i = 0; s.length - 1 - i >= (end - start + 1) / 2; i++) {
    let j = i
    let k = i
    for (; j >= 0 && k < s.length && s[j] === s[k]; j--, k++) {
      if (k - j > end - start) {
        end = k
        start = j
      }
    }
    j = i
    k = i + 1
    for (; j >= 0 && k < s.length && s[j] === s[k]; j--, k++) {
      if (k - j > end - start) {
        end = k
        start = j
      }
    }
  }
  return s.slice(start, end + 1)
};

const res = longestPalindrome('ccc')

