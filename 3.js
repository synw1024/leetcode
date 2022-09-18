/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
  let max = 0, newMax = 0, start = 0
  const map = {}
  for (let i = 0; i < s.length; i++) {
    const letter = s[i]
    if (map[letter] === undefined || map[letter] < start) {
      map[letter] = i
      newMax++
    } else {
      max = newMax > max ? newMax : max
      newMax = i - map[letter]
      start = map[letter] + 1
      map[letter] = i
    }
  }
  return newMax > max ? newMax : max
};

console.log(lengthOfLongestSubstring("abcabcbb"))
console.log(lengthOfLongestSubstring("bbbbb"))
console.log(lengthOfLongestSubstring("pwwkew"))
console.log(lengthOfLongestSubstring("abba"))
console.log(lengthOfLongestSubstring("tmmzuxt"))