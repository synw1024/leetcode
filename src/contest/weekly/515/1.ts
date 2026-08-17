function nearestDrone(drones: number[][], target: number[]): number {
  let min = Number.MAX_SAFE_INTEGER, index = -1
  for (let i = 0; i < drones.length; i++) {
    const drone = drones[i]
    const d = Math.abs(drone[0] - target[0]) + Math.abs(drone[1] - target[1])
    if (min > d && d <= drone[2]) {
      min = d
      index = i
    }
  }
  return index
}
