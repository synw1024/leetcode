function strStr(haystack: string, needle: string): number {
  let i = 0, j = 0
  for (; i < haystack.length && j < needle.length; i++) {
    if (haystack[i] === needle[j]) {
      j++
    } else {
      i -= j
      j = 0
    }
  }
  return j === needle.length ? i - j : -1
};