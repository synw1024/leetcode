class LRUCache2 {
  map: {
    [key: number]: number
  } = {}
  accessList: number[] = []
  constructor(private capacity: number) { }

  get(key: number): number {
    const val = this.map[key]
    if (val === undefined) {
      return -1
    }
    const index = this.accessList.indexOf(key)
    this.accessList.splice(index, 1)
    this.accessList.push(key)
    return val
  }

  put(key: number, value: number): void {
    const val = this.map[key]

    if (val === undefined && this.accessList.length >= this.capacity) {
      const k = this.accessList.shift()!
      delete this.map[k]
    }

    if (val !== undefined) {
      const index = this.accessList.indexOf(key)
      this.accessList.splice(index, 1)
    }
    this.accessList.push(key)
    this.map[key] = value
  }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */