function countGroups(position: number[], speed: number[], distance: number): number {
  let counter = 1
  let lastSpeed = speed[speed.length - 1]
  for (let i = speed.length - 1; i > 0; i--) {
    if (position[i] - position[i-1] <= distance || lastSpeed < speed[i-1]) {
      continue
    } else {
      counter++
      lastSpeed = speed[i-1]
    }
  }
  return counter
};

console.log(countGroups([259,349,604,944], [173,505,468,659], 203))
