package iching

import "core:math/rand"

Hexagram :: struct {
	lines:   u8,
	changes: u8,
}

yarrow_stalk :: proc() -> (hexagram: Hexagram) {
	odds := [2][8]u8{{0, 0, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 1, 1, 1}}
	hexagram.lines = u8(rand.uint_max(64))
	for i in 0 ..< 6 {
		line := index(hexagram.lines, i)
		hexagram.changes |= odds[line][u8(rand.uint_max(8))] << u8(i)
	}
	return
}

three_coins :: proc() -> (hexagram: Hexagram) {
	odds := [8]u8{0, 0, 0, 0, 0, 0, 1, 1}
	hexagram.lines = u8(rand.uint_max(64))
	for i in 0 ..< 6 {
		hexagram.changes |= odds[rand.uint_max(8)] << u8(i)
	}
	return
}

single_line :: proc() -> (hexagram: Hexagram) {
	hexagram.lines = u8(rand.uint_max(64))
	line := rand.int_max(6)
	hexagram.changes |= 1 << u8(line)
	return
}

index :: proc(num: u8, index: int) -> u8 {
	return (num << u8(7 - index)) >> 7
}

secondary_hexagram :: proc(primary: Hexagram) -> (secondary: Hexagram) {
	secondary.lines = primary.lines ~ primary.changes
	return
}

reversed_hexagram :: proc(primary: Hexagram) -> (reversed: Hexagram) {
	reversed.lines = (~primary.lines << 2) >> 2
	reversed.changes = primary.changes
	return
}

anti_hexagram :: proc(primary: Hexagram) -> (anti: Hexagram) {
	anti.lines = reversed_hexagram(primary).lines ~ primary.changes
	anti.changes = (~primary.changes << 2) >> 2
	return
}
