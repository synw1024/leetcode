function maxArea(height: number[]): number {
  let i = 0, j = height.length - 1
  let area = (j - i) * Math.min(height[i], height[j])
  while (i < j) {
    if (height[i] <= height[j]) {
      let p = i + 1
      while (p < j) {
        if (height[p] > height[i]) {
          i = p
          area = Math.max((j - i) * Math.min(height[i], height[j]), area)
          break
        }
        p++
      }
      if (p >= j) return area
    } else {
      let p = j - 1
      while (p > i) {
        if (height[p] > height[j]) {
          j = p
          area = Math.max((j - i) * Math.min(height[i], height[j]), area)
          break
        }
        p--
      }
      if (p <= i) return area
    }
  }
  return area
};
