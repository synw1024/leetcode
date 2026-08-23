/**
 * 
 */
function longestSubarray(nums: number[], k: number): number {
  
};

function primeFactor(n: number) {
  const result: {[key: number]: number} = {};
  // 处理因子2
  while (n % 2 === 0) {
    result[2] = (result[2] || 0) + 1;
    n = n / 2;
  }
  // 从3开始试除，只算奇数
  for (let i = 3; i * i <= n; i += 2) {
    while (n % i === 0) {
      result[i] = (result[i] || 0) + 1;
      n = n / i;
    }
  }
  // 剩下大于1，本身是质数
  if (n > 1) {
    result[n] = 1;
  }
  return Object.keys(result);
}
