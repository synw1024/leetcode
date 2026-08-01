/**
 * fn(n) = const(n) + max()
 */
function maximumWidth(planks: number[]): number {
  const sorted = Array.from(new Set(planks)).sort((a, b) => b - a)

  const map: {[key: string]: number} = {}
  for (let i = 0; i < planks.length; i++) {
    if (map[planks[i]]) {
      map[planks[i]]++
    } else {
      map[planks[i]] = 1
    }
  }
  
  function recurse(height: number, sorted: number[]) {
    let max = 0
    for (let i = 0; i < sorted.length; i++) {
      
    }
  }
  recurse(sorted[0], sorted.slice(1))



};

console.log(maximumWidth([1,3,2,5,7,5,4,2,1]))