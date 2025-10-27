function addStrings(num1: string, num2: string): string {
  let s = ''
  for (let i = 0; i < num1.length || i < num2.length; i++) {
    const index1 = num1.length - i - 1
    const n1 = Number(index1 >= 0 ? num1[index1] : 0)
    const index2 = num2.length - i - 1
    const n2 = Number(index2 >= 0 ? num2[index2] : 0)
    const indexS = s.length - i - 1
    const ns = Number()
  }
  // const arr1 = num1.split('').reverse().map(n => Number(n))
  // const arr2 = num2.split('').reverse().map(n => Number(n))
  // const arr: number[] = []
  // let i = 0
  // while (i < arr1.length || i < arr2.length) {
  //   const sum = (arr1[i] || 0) + (arr2[i] || 0)
  //   arr[i] = sum % 10
  //   if (sum > 9) {
  //     arr1[i+1] = arr1[i+1] ? arr1[i+1] + 1 : 1
  //   }
  //   i++
  // }
  // return arr.reverse().join('')
};
