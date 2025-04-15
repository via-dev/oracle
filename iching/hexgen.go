package iching

import (
	"math/rand/v2"
)

type Hexagram struct {
	lines   uint8
	changes uint8
}

func YarrowStalk() (hexagram Hexagram) {
	odds := [2][8]uint8{
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1, 1, 1},
	}
	hexagram.lines = uint8(rand.UintN(64))
	for i := range 6 {
		line := Index(hexagram.lines, i)
		hexagram.changes |= odds[line][rand.UintN(8)] << i
	}
	return
}

func ThreeCoins() (hexagram Hexagram) {
	odds := [8]uint8{0, 0, 0, 0, 0, 0, 1, 1}
	hexagram.lines = uint8(rand.UintN(64))
	for i := range 6 {
		hexagram.changes |= odds[rand.UintN(8)] << i
	}
	return
}

func SingleLine() (hexagram Hexagram) {
	hexagram.lines = uint8(rand.UintN(64))
	line := rand.IntN(6)
	hexagram.changes |= 1 << line
	return
}

func Index(num uint8, index int) uint8 {
	return (num << (7 - index)) >> 7
}

func SecondaryHexagram(primary Hexagram) (secondary Hexagram) {
	secondary.lines = primary.lines ^ primary.changes
	return
}

func ReversedHexagram(primary Hexagram) (reversed Hexagram) {
	reversed.lines = (^primary.lines << 2) >> 2
	reversed.changes = primary.changes
	return
}

func AntiHexagram(primary Hexagram) (anti Hexagram) {
	anti.lines = ReversedHexagram(primary).lines ^ primary.changes
	anti.changes = (^primary.changes << 2) >> 2
	return
}
