package main

func isBadVersion(version int) bool

func main() {

}

func firstBadVersion(n int) int {
	min, max := 1, n
	for true {
		target := (min + max) / 2
		if isBadVersion(target) {
			if !isBadVersion(target - 1) {
				return target
			}
			max = target - 2
		} else {
			if isBadVersion(target + 1) {
				return target + 1
			}
			min = target + 2
		}
	}
	return 0
}
