package solo

import "math/rand/v2"

func roll(num, sides int) (dice []int, res int) {
	for i := 1; i <= num; i++ {
		die := rand.IntN(sides) + 1
		dice = append(dice, die)
		res += die
	}
	return
}

func d3(n int) int {
	_, res := roll(n, 3)
	return res
}

func d4(n int) int {
	_, res := roll(n, 4)
	return res
}

func d6(n int) int {
	_, res := roll(n, 6)
	return res
}

func d8(n int) int {
	_, res := roll(n, 8)
	return res
}

func d10(n int) int {
	_, res := roll(n, 10)
	return res
}

func d12(n int) int {
	_, res := roll(n, 12)
	return res
}

func d20(n int) int {
	_, res := roll(n, 20)
	return res
}

func d100(n int) int {
	_, res := roll(n, 100)
	return res
}

func dF(n int) int {
	return d3(n) - (n * 2)
}
