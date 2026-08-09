function minPrice(prices: number[], discounts: number[]): number {
  prices.sort((a, b) => b - a)
  discounts.sort((a, b) => b - a)

  return prices.reduce((sum, price, index) => {
    return sum + price * (100 - (discounts[index] || 0)) / 100
  }, 0)
}

minPrice([10,30,21], [50, 60])