/**
 * 0-不偷，1-偷
 * dp[0][i] = max(dp[0][i-1], dp[1][i-1])
 * dp[1][i] = max(dp[0][i-1]+nums[i])
 */
var rob = function (nums) {
  const dp = [[0], [nums[0]]]
  for (let i = 1; i < nums.length; i++) {
    dp[0][i] = Math.max(dp[0][i - 1], dp[1][i - 1])
    dp[1][i] = Math.max(dp[0][i - 1] + nums[i])
  }
  const i = nums.length - 1
  return Math.max(dp[0][i], dp[1][i])
};
