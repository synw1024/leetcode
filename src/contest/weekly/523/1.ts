function threeFibonacciSum(n: number): boolean {
  if (n === 2) return true
  
  let x = 0, y = 1
  while (x < n / 2 && y < n / 2) {
    if (x + y === n / 2) return true
    const z = x + y
    x = y
    y = z
  }
  return false
};