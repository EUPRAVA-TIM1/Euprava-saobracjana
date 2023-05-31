package service

import (
	"EuprvaSaobracajna/data"
	"time"
)

type SudService interface {
	SendNalog(nalog data.SudskiNalog, PTT string) error
	GetGradjaninSlucajevi(jmbg string) (slucajevi []*data.SudskiSlucaj, err error)
}

type sudServiceImpl struct {
	serviceUrl string
}

func NewSudService(url string) SudService {
	return sudServiceImpl{serviceUrl: url}
}

func (s sudServiceImpl) SendNalog(nalog data.SudskiNalog, PTT string) error {
	//TODO implement me
	return nil
}

func (s sudServiceImpl) GetGradjaninSlucajevi(jmbg string) (slucajevi []*data.SudskiSlucaj, err error) {
	//TODO implement me
	nalozi := make([]*data.SudskiSlucaj, 0)
	nalozi = append(nalozi, &data.SudskiSlucaj{
		Datum:  time.Now(),
		Naslov: "Prekoracenje brzine u zoni skole",
		Opis:   "lorem ipsum neki opis nmg sad da se setim bolje",
		Status: "U_PROCESU",
	})
	nalozi = append(nalozi, &data.SudskiSlucaj{
		Datum:  time.Now(),
		Naslov: "Voznja pod dejstvom alkohola",
		Opis:   "lorem ipsum neki opis nmg sad da se setim bolje",
		Status: "U_PROCESU",
	})
	return nalozi, nil
}
