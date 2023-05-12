package service

import (
	"EuprvaSaobracajna/data"
	"math"
)

const dozvoljeniPromili = 0.3
const poeniZaRegularnePrekrsaje = 1
const poneiZaVecePrekrsaje = 1

func CalculatePointsForTicket(kazna data.PrekrsajniNalog) int {
	switch kazna.TipPrekrsaja {
	case "POJAS":
		return poeniZaRegularnePrekrsaje
	case "PREKORACENJE_BRZINE":
		points := (*kazna.Vrednost / 10) * 2
		if *kazna.Vrednost/10 > 0 {
			points += 1
		}
		return int(points)
	case "PIJANA_VOZNJA":
		if *kazna.Vrednost > 0.3 {
			points := int(math.Floor((*kazna.Vrednost - dozvoljeniPromili) / 0.2))
			return points
		}
		return 0
	case "TEHNICKA_NEISPRAVNOST":
		return poneiZaVecePrekrsaje
	case "REGISTRACIJA":
		return poneiZaVecePrekrsaje
	}
	return 0
}
