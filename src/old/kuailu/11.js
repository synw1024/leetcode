function factorial2(n) {
  if (n === 1) return 1
  return n * factorial2(n - 1)
}

function factorial3(n) {
  let product = 1
  for (let i = n; i > 1; i--) {
    product *= i
  }
  return product
}

const res = factorial(5)
console.log(res)