package iching

import "core:math/rand"

Hexagram :: struct {
	lines:   u8,
	changes: u8,
}

YarrowStalk :: proc() -> (hexagram: Hexagram) {
	odds := [2][8]u8{{0, 0, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 1, 1, 1}}
	hexagram.lines = u8(rand.uint_max(64))
	for i in 0 ..< 6 {
		line := Index(hexagram.lines, i)
		hexagram.changes |= odds[line][u8(rand.uint_max(8))] << u8(i)
	}
	return
}

ThreeCoins :: proc() -> (hexagram: Hexagram) {
	odds := [8]u8{0, 0, 0, 0, 0, 0, 1, 1}
	hexagram.lines = u8(rand.uint_max(64))
	for i in 0 ..< 6 {
		hexagram.changes |= odds[rand.uint_max(8)] << u8(i)
	}
	return
}

SingleLine :: proc() -> (hexagram: Hexagram) {
	hexagram.lines = u8(rand.uint_max(64))
	line := rand.int_max(6)
	hexagram.changes |= 1 << u8(line)
	return
}

Index :: proc(num: u8, index: int) -> u8 {
	return (num << u8(7 - index)) >> 7
}

SecondaryHexagram :: proc(primary: Hexagram) -> (secondary: Hexagram) {
	secondary.lines = primary.lines ~ primary.changes
	return
}

ReversedHexagram :: proc(primary: Hexagram) -> (reversed: Hexagram) {
	reversed.lines = (~primary.lines << 2) >> 2
	reversed.changes = primary.changes
	return
}

AntiHexagram :: proc(primary: Hexagram) -> (anti: Hexagram) {
	anti.lines = ReversedHexagram(primary).lines ~ primary.changes
	anti.changes = (~primary.changes << 2) >> 2
	return
}
