package service

import (
	"EuprvaSaobracajna/data"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

type SudService interface {
	SendNalog(nalog data.SudskiNalog) error
	GetGradjaninSlucajevi(jmbg string) (slucajevi []data.SudskiSlucaj, err error)
}

type sudServiceImpl struct {
	serviceUrl string
}

func NewSudService(url string) SudService {
	return sudServiceImpl{serviceUrl: url}
}

func (s sudServiceImpl) SendNalog(nalog data.SudskiNalog) (err error) {
	u, err := url.Parse(s.serviceUrl + "/api/prekrsajnaprijava")
	if err != nil {
		return
	}
	nalog.StatusSlucaja = "0"
	json, err := json.Marshal(nalog)
	if err != nil {
		log.Fatal(err)
		return
	}
	client := http.Client{}
	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(json))
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return
	}

	if response.StatusCode == http.StatusOK {
		return
	}
	err = errors.New("can not reach Sud service")
	return
}

type SlucajeviWrapper struct {
	Values []data.SudskiSlucaj `json:"$values"`
}

func (s sudServiceImpl) GetGradjaninSlucajevi(jmbg string) (slucajevi []data.SudskiSlucaj, err error) {
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
		var retVal SlucajeviWrapper
		err = json.NewDecoder(response.Body).Decode(&retVal)
		if err != nil {
			log.Fatal(err)
		}
		slucajevi = retVal.Values
		return
	}
	err = errors.New("can not reach Sud service")
	return
}
