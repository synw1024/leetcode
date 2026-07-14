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
  }

  delete() {
    delete this.map[this.tail!.key]
    this.tail = this.tail!.prev
    if (this.tail) {
      this.tail.next = null
    } else {
      this.head = null
    }
  }

  unshift(key: number, value: number) {
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
  }
}
