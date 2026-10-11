function maxPrimes(n: number, s: number): number[] {
  const res: number[] = []
  let sum = 0
  function isPrime(n: number): boolean {
    // 2 是唯一偶质数
    if (n === 2) return true;
    // 偶数直接排除
    if (n % 2 === 0) return false;

    // 只循环到 sqrt(n)，步长2只校验奇数
    const sqrtVal = Math.sqrt(n);
    for (let i = 3; i <= sqrtVal; i += 2) {
      if (n % i === 0) {
        return false;
      }
    }
    return true;
  }
  for (let i = 2; i <= n; i++) {
    if (!isPrime(i)) continue

    if (sum + i > s) break

    res.push(i)
    sum += i
  }
  return res
}
