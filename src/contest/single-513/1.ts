function maxPairStrength(nums: number[]): number {
  let max = 0
  for (let i = 0; i < nums.length - 1; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      max = Math.max(max, nums[i] * nums[j] / gcd2(nums[i], nums[j]))
    }
  }
  return max
};

function gcd2(a: number, b: number) {
  a = Math.abs(a);
  b = Math.abs(b);
  while (b !== 0) {
    const temp = b;
    b = a % b;
    a = temp;
  }
  return a * a;
}