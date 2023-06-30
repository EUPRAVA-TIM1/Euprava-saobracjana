package service

import (
	"EuprvaSaobracajna/data"
	"math"
)

const dozvoljeniPromili = 0.3
const promiliIncrement = 0.2
const brzinaIncrement = 10
const brzinaIncrementPoeni = 2
const poeniZaRegularnePrekrsaje = 1
const poneiZaVecePrekrsaje = 1

func CalculatePointsForTicket(kazna data.PrekrsajniNalog) int {
	switch kazna.TipPrekrsaja {
	case "POJAS":
		return poeniZaRegularnePrekrsaje
	case "PREKORACENJE_BRZINE":
		points := (*kazna.Vrednost / brzinaIncrement) * brzinaIncrementPoeni
		if int(*kazna.Vrednost)%brzinaIncrement > 0 {
			points += 1
		}
		return int(points)
	case "PIJANA_VOZNJA":
		if *kazna.Vrednost > dozvoljeniPromili {
			return int(math.Floor((*kazna.Vrednost - dozvoljeniPromili) / promiliIncrement))
		}
		return 0
	case "TEHNICKA_NEISPRAVNOST":
		return poneiZaVecePrekrsaje
	case "REGISTRACIJA":
		return poneiZaVecePrekrsaje
	}
	return 0
}
