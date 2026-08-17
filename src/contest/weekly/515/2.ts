function minPenalty(period: number, lights: number[], arrivalTime: number[]): number {
  const maxLignt = lights.reduce((max, l) => max < l ? l : max, 0)
  let max = 0
  for (let i = 0; i < arrivalTime.length; i++) {
    const arrival = arrivalTime[i] % period
    if (arrival < maxLignt) continue
    const w = period - arrival
    max = w > max ? w : max
  }
  return max
};

