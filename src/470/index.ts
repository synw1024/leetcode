function rand10(): number {
  let n = 0
  do {
    n = rand7() + 7 * (rand7() - 1)
  } while(n > 40);
  return (n - 1) % 10 + 1
};

const rand7 = () => 0