package service

import (
	"EuprvaSaobracajna/data"
	"time"
)

type MupService interface {
	SendKradjaPrijava(prijava data.PrijavaKradjeVozila) error
	GetVozacka(jmbg string) (*data.VozackaDozvola, error)
	GetSaobracjana(tablica string) (*data.SaobracjanaDozvola, error)
	SendPoints(points int) error
}

type mupServiceImpl struct {
	serviceUrl string
}

func NewMupService(url string) MupService {
	return mupServiceImpl{serviceUrl: url}
}

func (m mupServiceImpl) SendPoints(points int) error {
	//TODO implement me
	return nil
}

func (m mupServiceImpl) SendKradjaPrijava(prijava data.PrijavaKradjeVozila) error {
	//TODO implement me
	return nil
}

func (m mupServiceImpl) GetVozacka(jmbg string) (*data.VozackaDozvola, error) {
	switch jmbg {
	case "1":
		return &data.VozackaDozvola{
			BrojVozackeDozvole:   "1",
			KategorijeVozila:     append(make([]string, 0), "A", "B"),
			DatumIzdavavanja:     time.Now(),
			DatumIsteka:          time.Now().Add(time.Hour * 1000),
			BrojKaznenihPoena:    13,
			StatusVozackeDozvole: "AKTIVNA",
		}, nil
	case "2":
		return &data.VozackaDozvola{
			BrojVozackeDozvole:   "2",
			KategorijeVozila:     append(make([]string, 0), "A"),
			DatumIzdavavanja:     time.Now().Add(time.Hour * -1000),
			DatumIsteka:          time.Now(),
			BrojKaznenihPoena:    8,
			StatusVozackeDozvole: "ISTEKLA",
		}, nil
	default:
		return &data.VozackaDozvola{
			BrojVozackeDozvole:   "3",
			KategorijeVozila:     append(make([]string, 0), "A", "B", "C"),
			DatumIzdavavanja:     time.Now(),
			DatumIsteka:          time.Now().Add(time.Hour * 1000),
			BrojKaznenihPoena:    0,
			StatusVozackeDozvole: "ODUZETA",
		}, nil
	}
}

func (m mupServiceImpl) GetSaobracjana(tablica string) (*data.SaobracjanaDozvola, error) {
	switch tablica {
	case "222":
		time := time.Now()
		return &data.SaobracjanaDozvola{
			Marka:              "Audi",
			Model:              "A6",
			GodinaProizvodnje:  2012,
			Boja:               "Plava",
			RegBroj:            "222",
			SnagaMotora:        120,
			MaksimalnaBrzina:   240,
			BrojSedista:        5,
			Tezina:             4,
			TipVozila:          "PUTNICKO_VOZILO",
			StatusRegistracije: "ODOBRENA",
			PrijavljenaKradja:  &time,
		}, nil
	case "111":
		return &data.SaobracjanaDozvola{
			Marka:              "Wolkswagen Golf",
			Model:              "6",
			GodinaProizvodnje:  2010,
			Boja:               "Crna",
			RegBroj:            "111",
			SnagaMotora:        106,
			MaksimalnaBrzina:   200,
			BrojSedista:        5,
			Tezina:             2,
			TipVozila:          "PUTNICKO_VOZILO",
			StatusRegistracije: "ODOBRENA",
		}, nil
	default:
		return &data.SaobracjanaDozvola{
			Marka:              "Honda",
			Model:              "R6",
			GodinaProizvodnje:  2016,
			Boja:               "Crvena",
			RegBroj:            "111",
			SnagaMotora:        300,
			MaksimalnaBrzina:   360,
			BrojSedista:        1,
			Tezina:             1,
			TipVozila:          "SKUTER",
			StatusRegistracije: "ODOBRENA",
		}, nil
	}
}
