function generateParenthesis(n: number): string[] {
  const res: string[] = []
  function recurse(s: string, open: number, close: number) {
    if (!open && !close) {
      res.push(s)
      return
    }
    if (open > 0) {
      recurse(s + '(', open - 1, close)
    }
    if (close > 0 && close > open) {
      recurse(s + ')', open, close - 1)
    }
  }
  recurse('(', n - 1, n)
  return res
};