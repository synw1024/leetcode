type Link = {
  key: number,
  prev: Link | null
  next: Link | null
  value: number
}

class LRUCache {
  map: {
    [key: number]: Link
  } = {}
  head: Link | null = null
  tail: Link | null = null
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
    const link = this.map[key]

    if (!link && this.capacity <= 0) {
      delete this.map[this.tail!.key]
      this.tail = this.tail!.prev
      if (this.tail) {
        this.tail.next = null
      } else {
        this.head = null
      }
    }

    if (!link) {
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
    } else {
      link.value = value
      this.move(link)
    }
  }

  move(link: Link) {
    if (!link.prev) return

    link.prev.next = link.next

    if (link.next) {
      link.next.prev = link.prev
    } else {
      this.tail = link.prev
    }

    link.next = this.head 
    this.head = link
  }
}

const obj = new LRUCache(2)
const params = [[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
for (let i = 0; i < params.length; i++) {
  const p = params[i]
  if (p.length === 1) {
    obj.get(p[0])
  } else {
    obj.put(p[0], p[1])
  }
}
