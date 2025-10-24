type Link2 = {
  key: number,
  prev: Link | null
  next: Link | null
  value: number
}

class LRUCache2 {
  map: {
    [key: number]: Link
  } = {}
  head: Link | null = null
  tail: Link | null = null
  accessList: number[] = []
  constructor(private capacity: number) { }

  get(key: number): number {
    const link = this.map[key]
    if (!link) {
      return -1
    }

    this.move(link)

    return link.value
  }

  put(key: number, value: number): void {
    if (key === 9 && value === 3) {
      debugger
    }
    const link = this.map[key]

    if (link) {
      link.value = value
      this.move(link)
      return
    }

    if (this.capacity <= 0) {
      this.delete()
    }

    this.unshift(key, value)
  }

  move(link: Link) {
    if (!link.prev) return

    const index = this.accessList.indexOf(link.key)
    this.accessList.splice(index, 1)
    this.accessList.unshift(link.key)

    link.prev.next = link.next

    if (link.next) {
      link.next.prev = link.prev
    } else {
      this.tail = link.prev
    }

    link.next = this.head
    this.head!.prev = link
    this.head = link
    this.head.prev = null

    // this.test()
  }

  delete() {
    this.accessList.pop()

    delete this.map[this.tail!.key]
    this.tail = this.tail!.prev
    if (this.tail) {
      this.tail.next = null
    } else {
      this.head = null
    }

    // this.test()
  }

  unshift(key: number, value: number) {
    this.accessList.unshift(key)


    const l: Link = {
      key,
      prev: null,
      next: this.head,
      value
    }
    this.map[key] = l
    if (this.head) {
      this.head.prev = l
    } else {
      this.tail = l
    }
    this.head = l
    this.capacity--

    // this.test()
  }

  test() {
    let pointer = this.head!
    const keys = []
    let counter = 0
    while(pointer) {
      if (counter === this.accessList.length) {
        debugger
      }
      keys.push(pointer.key)
      pointer = pointer.next!
      counter++
    }

    const a = this.accessList.join('-')
    const k = keys.join('-')
    if (a !== k) {
      debugger
    }
  }
}

const obj = new LRUCache(10)
const params = [[10, 13], [3, 17], [6, 11], [10, 5], [9, 10], [13], [2, 19], [2], [3], [5, 25], [8], [9, 22], [5, 5], [1, 30], [11], [9, 12], [7], [5], [8], [9], [4, 30], [9, 3], [9], [10], [10], [6, 14], [3, 1], [3], [10, 11], [8], [2, 14], [1], [5], [4], [11, 4], [12, 24], [5, 18], [13], [7, 23], [8], [12], [3, 27], [2, 12], [5], [2, 9], [13, 4], [8, 18], [1, 7], [6], [9, 29], [8, 21], [5], [6, 30], [1, 12], [10], [4, 15], [7, 22], [11, 26], [8, 17], [9, 29], [5], [3, 4], [11, 30], [12], [4, 29], [3], [9], [6], [3, 4], [1], [10], [3, 29], [10, 28], [1, 20], [11, 13], [3], [3, 12], [3, 8], [10, 9], [3, 26], [8], [7], [5], [13, 17], [2, 27], [11, 15], [12], [9, 19], [2, 15], [3, 16], [1], [12, 17], [9, 1], [6, 19], [4], [5], [5], [8, 1], [11, 7], [5, 2], [9, 28], [1], [2, 2], [7, 4], [4, 22], [7, 24], [9, 26], [13, 28], [11, 26]]
for (let i = 0; i < params.length; i++) {
  const p = params[i]
  if (p.length === 1) {
    obj.get(p[0])
  } else {
    obj.put(p[0], p[1])
  }
}
