package geomancy

import "core:math/rand"

Chart :: struct {
	Mothers:    [4]u8,
	Daughters:  [4]u8,
	Nieces:     [4]u8,
	Lwitness:   u8,
	Rwitness:   u8,
	Judge:      u8,
	Reconciler: u8,
}

GetMothers :: proc() -> (figures: [4]u8) {
	for i in 0 ..< 4 {
		for n in 0 ..< 4 {
			figures[i] |= u8(rand.uint_max(2)) << u8(n)
		}
	}
	return
}

GetFigure :: proc() -> u8 {
	return u8(rand.uint_max(1 << 4))
}

Index :: proc(num: u8, index: int) -> u8 {
	return (num << u8(7 - index)) >> 7
}

GetDaughters :: proc(mothers: [4]u8) -> (figures: [4]u8) {
	for f in 0 ..< 4 {
		figures[f] |= Index(mothers[0], f) << 0
		figures[f] |= Index(mothers[1], f) << 1
		figures[f] |= Index(mothers[2], f) << 2
		figures[f] |= Index(mothers[3], f) << 3
	}
	return
}

Combine :: proc(fig1, fig2: u8) -> (newfig: u8) {
	newfig = fig1 ~ fig2
	return
}

GetNieces :: proc(mothers, daughters: [4]u8) -> (figures: [4]u8) {
	figures[0] = Combine(mothers[0], mothers[1])
	figures[1] = Combine(mothers[2], mothers[3])
	figures[2] = Combine(daughters[0], daughters[1])
	figures[3] = Combine(daughters[2], daughters[3])
	return
}

GetWinesses :: proc(nieces: [4]u8) -> (lwitness, rwitness: u8) {
	rwitness = Combine(nieces[0], nieces[1])
	lwitness = Combine(nieces[2], nieces[3])
	return
}

GenChart :: proc(mothers: [4]u8) -> (chart: Chart) {
	chart.Mothers = mothers
	chart.Daughters = GetDaughters(chart.Mothers)
	chart.Nieces = GetNieces(chart.Mothers, chart.Daughters)
	chart.Lwitness, chart.Rwitness = GetWinesses(chart.Nieces)
	chart.Judge = Combine(chart.Rwitness, chart.Lwitness)
	chart.Reconciler = Combine(chart.Judge, chart.Mothers[0])
	return
}
