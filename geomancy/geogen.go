package geomancy

import (
	"math/rand/v2"
)

type Chart struct {
	Mothers    [4]uint8
	Daughters  [4]uint8
	Nieces     [4]uint8
	Lwitness   uint8
	Rwitness   uint8
	Judge      uint8
	Reconciler uint8
}

func GetMothers() (figures [4]uint8) {
	for i := range figures {
		for n := range 4 {
			figures[i] |= uint8(rand.UintN(2)) << n
		}
	}
	return
}

func GetFigure() uint8 {
	return uint8(rand.UintN(1 << 4))
}

func Index(num uint8, index int) uint8 {
	return (num << (7 - index)) >> 7
}

func GetDaughters(mothers [4]uint8) (figures [4]uint8) {
	for f := range 4 {
		figures[f] |= Index(mothers[0], f) << 0
		figures[f] |= Index(mothers[1], f) << 1
		figures[f] |= Index(mothers[2], f) << 2
		figures[f] |= Index(mothers[3], f) << 3
	}
	return
}

func Combine(fig1, fig2 uint8) (newfig uint8) {
	newfig = fig1 ^ fig2
	return
}

func GetNieces(mothers, daughters [4]uint8) (figures [4]uint8) {
	figures[0] = Combine(mothers[0], mothers[1])
	figures[1] = Combine(mothers[2], mothers[3])
	figures[2] = Combine(daughters[0], daughters[1])
	figures[3] = Combine(daughters[2], daughters[3])
	return
}

func GetWinesses(nieces [4]uint8) (lwitness, rwitness uint8) {
	rwitness = Combine(nieces[0], nieces[1])
	lwitness = Combine(nieces[2], nieces[3])
	return
}

func GenChart(mothers [4]uint8) (chart Chart) {
	chart.Mothers = mothers
	chart.Daughters = GetDaughters(chart.Mothers)
	chart.Nieces = GetNieces(chart.Mothers, chart.Daughters)
	chart.Lwitness, chart.Rwitness = GetWinesses(chart.Nieces)
	chart.Judge = Combine(chart.Rwitness, chart.Lwitness)
	chart.Reconciler = Combine(chart.Judge, chart.Mothers[0])
	return
}
