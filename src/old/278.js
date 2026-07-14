var solution = function (isBadVersion) {
  /**
   * @param {integer} n Total versions
   * @return {integer} The first bad version
   */
  return function (n) {
    function binarySearch(start, end) {
      const mid = Math.round((start + end) / 2)
      if (!isBadVersion(mid)) {
        return binarySearch(mid+1, end)
      } else if (!isBadVersion(mid-1)) {
        return mid
      } else {
        return binarySearch(start, mid-1)
      }
    }
    return binarySearch(1, n)
  };
};