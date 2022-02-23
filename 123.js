/**
 * 0-买过一次，1-卖过一次，2-买过两次，3-卖过两次
 * dp[0][i] = max(-prices[i], dp[0][i-1])
 * dp[1][i] = max(dp[0][i-1]+prices[i], dp[1][i-1])
 * dp[2][i] = max(dp[1][i-1] - prices[i], dp[2][i-1])
 * dp[3][i] = max(dp[2][i-1] + prices[i], dp[3][i-1])
 */
var maxProfit = function (prices) {
  const dp = [
    [-prices[0]],
    [-Number.MAX_SAFE_INTEGER],
    [-Number.MAX_SAFE_INTEGER, -Number.MAX_SAFE_INTEGER],
    [-Number.MAX_SAFE_INTEGER, -Number.MAX_SAFE_INTEGER, -Number.MAX_SAFE_INTEGER]
  ]

  for (let i = 1; i < prices.length; i++) {
    dp[0][i] = Math.max(-prices[i], dp[0][i - 1])
    dp[1][i] = Math.max(dp[0][i - 1] + prices[i], dp[1][i - 1])
    if (i > 1) {
      dp[2][i] = Math.max(dp[1][i - 1] - prices[i], dp[2][i - 1])
      if (i > 2) {
        dp[3][i] = Math.max(dp[2][i - 1] + prices[i], dp[3][i - 1])
      }
    }
  }

  const i = prices.length - 1
  return Math.max(dp[1][i], dp[3][i], 0)
};

const res = maxProfit([1, 2])
console.log(res)
