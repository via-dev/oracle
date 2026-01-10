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

get_mothers :: proc() -> (figures: [4]u8) {
	for i in 0 ..< 4 {
		for n in 0 ..< 4 {
			figures[i] |= u8(rand.uint_max(2)) << u8(n)
		}
	}
	return
}

get_figure :: proc() -> u8 {
	return u8(rand.uint_max(1 << 4))
}

index :: proc(num: u8, index: int) -> u8 {
	return (num << u8(7 - index)) >> 7
}

get_daughters :: proc(mothers: [4]u8) -> (figures: [4]u8) {
	for f in 0 ..< 4 {
		figures[f] |= index(mothers[0], f) << 0
		figures[f] |= index(mothers[1], f) << 1
		figures[f] |= index(mothers[2], f) << 2
		figures[f] |= index(mothers[3], f) << 3
	}
	return
}

combine :: proc(fig1, fig2: u8) -> (newfig: u8) {
	newfig = fig1 ~ fig2
	return
}

get_nieces :: proc(mothers, daughters: [4]u8) -> (figures: [4]u8) {
	figures[0] = combine(mothers[0], mothers[1])
	figures[1] = combine(mothers[2], mothers[3])
	figures[2] = combine(daughters[0], daughters[1])
	figures[3] = combine(daughters[2], daughters[3])
	return
}

get_winesses :: proc(nieces: [4]u8) -> (lwitness, rwitness: u8) {
	rwitness = combine(nieces[0], nieces[1])
	lwitness = combine(nieces[2], nieces[3])
	return
}

gen_chart :: proc(mothers: [4]u8) -> (chart: Chart) {
	chart.Mothers = mothers
	chart.Daughters = get_daughters(chart.Mothers)
	chart.Nieces = get_nieces(chart.Mothers, chart.Daughters)
	chart.Lwitness, chart.Rwitness = get_winesses(chart.Nieces)
	chart.Judge = combine(chart.Rwitness, chart.Lwitness)
	chart.Reconciler = combine(chart.Judge, chart.Mothers[0])
	return
}
