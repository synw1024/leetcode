function resolve2(arr: number[]) {
  const map: {[key: number]: number} = {}
  arr.forEach(a => {
    if (!map[a]) {
      map[a] = 1
    } else {
      map[a]++
    }
  })
  const keys: string[] = []
  const arr2 = Object.entries(map)
  arr2.forEach(([key, val]) => {
    if (val > 1) keys.push(key)
  })
  console.log(keys)
}
