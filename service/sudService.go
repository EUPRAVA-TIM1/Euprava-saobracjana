package service

import (
	"EuprvaSaobracajna/data"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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

type slucajeviWrapper struct {
	id     int                  `json:"$id"`
	values []*data.SudskiSlucaj `json:"$values"`
}

func (s sudServiceImpl) GetGradjaninSlucajevi(jmbg string) (slucajevi []*data.SudskiSlucaj, err error) {
	u, err := url.Parse(s.serviceUrl + "/api/predmet/gradjanin/" + jmbg)
	if err != nil {
		return
	}

	client := http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(response.Body)

	if response.StatusCode == http.StatusOK {
		var retVal slucajeviWrapper
		err = json.NewDecoder(response.Body).Decode(&retVal)
		if err != nil {
			log.Fatal(err)
		}
		slucajevi = retVal.values
		fmt.Println(slucajevi)
		return
	}
	err = errors.New("can not reach Sud service")
	return
}
