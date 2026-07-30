function isValid(s: string): boolean {
  const stack = []
  const closeBracketMap: {[key: string]: string} = {
    ')': '(',
    '}': '{',
    ']': '['
  }
  for (let i = 0; i < s.length; i++) {
    if (!closeBracketMap[s[i]]) {
      stack.push(s[i])
    } else {
      if (!stack[stack.length - 1] || stack[stack.length - 1] !== closeBracketMap[s[i]]) return false
      stack.pop()
    }
  }
  return stack.length ? false : true
};

isValid('()[]{}')