class BusBooking {
  private seats: {[key: string]: number} = {}
  private total = 0
  private min = 0
  constructor(n: number) {
  }

  toggle(seat: string): void {
    const c = seat.slice(0, 1)
    const n = seat.slice(1)
    if (c === 'A') {
      if (this.seats[n] & 1) {
        this.seats[n] &= ~1
        this.total -= 3
        this.min -= 1
      } else if (this.seats[n] & 2) {
        this.seats[n] &= ~2
        this.total -= 1
        this.min -= 1
      } else if (this.seats[n] & 4) {
        this.seats[n] |= 1
        this.total += 3
        this.min += 1
      } else {
        this.seats[n] |= 2
        this.total += 1
        this.min += 1
      }
    } else if (c === 'B') {
      if (this.seats[n] & 4) {
        this.total -= 1
        this.min -= 1
      } else {
        this.total += 1
        this.min += 1
      }
      this.seats[n] ^= 4
    } else if (c === 'C') {
      if (this.seats[n] & 8) {
        this.total -= 1
        this.min -= 1
      } else {
        this.total += 1
        this.min += 1
      }
      this.seats[n] ^= 8
    } else {
      if (this.seats[n] & 16) {
        this.seats[n] &= ~16
        this.total -= 1
        this.min -= 1
      } else if (this.seats[n] & 32) {
        this.seats[n] &= ~32
        this.total -= 3
        this.min -= 1
      } else if (this.seats[n] & 8) {
        this.seats[n] |= 32
        this.total += 3
        this.min += 1
      } else {
        this.seats[n] |= 16
        this.total += 1
        this.min += 1
      }
    }
  }

  getTotalTime(): number {
    return this.total
  }

  getMinTime(): number {
    return this.min
  }
}

const b = new BusBooking(20)
b.toggle('B3')
b.toggle('A3')
b.getTotalTime()
b.getMinTime()