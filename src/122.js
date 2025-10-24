/**
 * 0-没股票，1-有股票
 * dp[0][i] = max(dp[0][i-1], dp[1][i-1]+prices[i])
 * dp[1][i] = max(dp[1][i-1], dp[0][i-1]-prices[i])
 */
var maxProfit = function(prices) {
  const dp = [[0], [-prices[0]]]
  for (let i = 1; i < prices.length; i++) {
    dp[0][i] = Math.max(dp[0][i-1], dp[1][i-1]+prices[i])
    dp[1][i] = Math.max(dp[1][i-1], dp[0][i-1]-prices[i])
  }
  return dp[0][prices.length-1]
};

const res = maxProfit([7,1,5,3,6,4])
console.log(res)
