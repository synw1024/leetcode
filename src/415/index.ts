function addStrings(num1: string, num2: string): string {
  let s = ''
  for (let i = 0; i < num1.length || i < num2.length; i++) {
    const index1 = num1.length - i - 1
    const n1 = Number(index1 >= 0 ? num1[index1] : 0)
    const index2 = num2.length - i - 1
    const n2 = Number(index2 >= 0 ? num2[index2] : 0)
    const sum = n1 + n2 + (s.length > i ? 1 : 0)

    if (sum >= 10) {
      s = '1' + String(sum % 10) + (s.length > i ? s.slice(1) : s)
    } else {
      s = String(sum) + (s.length > i ? s.slice(1) : s)
    }
  }
  return s
};

