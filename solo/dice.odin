package solo

import "core:math/rand"

roll :: proc(num, sides: int) -> (res: int) {
	for i := 1; i <= num; i += 1 {
		die := rand.int_max(sides) + 1
		res += die
	}
	return
}

d3 :: proc(n: int) -> int {
	return roll(n, 3)
}

d4 :: proc(n: int) -> int {
	return roll(n, 4)
}

d6 :: proc(n: int) -> int {
	return roll(n, 6)
}

d8 :: proc(n: int) -> int {
	return roll(n, 8)
}

d10 :: proc(n: int) -> int {
	return roll(n, 10)
}

d12 :: proc(n: int) -> int {
	return roll(n, 12)
}

d20 :: proc(n: int) -> int {
	return roll(n, 20)
}

d100 :: proc(n: int) -> int {
	return roll(n, 100)
}

dF :: proc(n: int) -> int {
	return d3(n) - (n * 2)
}
