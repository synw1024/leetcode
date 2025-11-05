function resolve(arr: number[]) {
  let res1 = 0
  for (let i = arr.length - 1; i >= 1; i = i - 2) {
    res1 += arr[i]
  }

  let res2 = 0
  for (let i = arr.length - 2; i >= 0; i = i - 2) {
    res2 += arr[i]
  }
  
  
  
  // function fn(n: number): number {
  //   let sum = 0
  //   let end = 0
  //   if (n === arr.length - 1) {
  //     end = 1
  //   }
  //   for (let i = n - 2; i >= end;i = i - 2) {
  //     sum += fn(i) + arr[n]
  //   }
  //   return sum
  // }
  // const res1 = fn(arr.length - 1)
  // const res2 = fn(arr.length - 2)
  return res1 > res2 ? res1 : res2
}

const test = [1, 2, 3, 4, 5, 6, 7]
const res = resolve(test)
console.log(res)