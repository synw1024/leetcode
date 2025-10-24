function shuffle(arr) {
  const tempArr = arr.slice()
  return function(n) {
    const res = []
    for (let i = 0; i < n; i++) {
      const index = Math.floor(Math.random() * tempArr.length)
      res.push(tempArr.splice(index, 1))
    }
    console.log(res)
  }
}

let random = shuffle([0, 1, 2, 3, 4, 5, 6])
random(1); // [1]
random(1); // [0]
random(1); // [2]
random(1); // [3]
random(1); // [5]
random(1); // [4]
random(1); // [6]
random(1);

random = shuffle([0,1,2,3,4,5,6]);
random(1); // [1]
random(2); // [0,6]
random(1); // [2]
random(4); // [3,4,5,2]
