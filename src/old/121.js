/**
 * 0-没买，1-买了没卖，2-卖了
 * dp[0][i] = 0
 * dp[1][i] = max(dp[0][i-1]-prices[i], dp[1][i-1])
 * dp[2][i] = max(dp[1][i-1]+prices[i], dp[2][i-1])
 * 
 * 因为 dp[0][i] 全部为 0，这个状态不需要记录，所以：
 * 
 * 0-买了没卖，1-卖了
 * dp[0][i] = max(-prices[i], dp[0][i-1])
 * dp[1][i] = max(dp[0][i-1]+prices[i], dp[1][i-1])
 */

var maxProfit = function(prices) {
  const dp = [[-prices[0]], [-Number.MAX_SAFE_INTEGER]]
  for (let i = 1; i < prices.length; i++) {
    dp[0][i] = Math.max(-prices[i], dp[0][i-1])
    dp[1][i] = Math.max(dp[0][i-1]+prices[i], dp[1][i-1])
  }
  const i = prices.length-1
  return dp[1][i] > 0 ? dp[1][i] : 0
};

maxProfit([7,1,5,3,6,4])
