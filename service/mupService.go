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

type MupService interface {
	SendKradjaPrijava(prijava data.PrijavaKradjeVozila) error
	GetVozacka(jmbg string) (*data.VozackaDozvola, error)
	GetSaobracjana(tablica string) (*data.SaobracjanaDozvola, error)
	SendPoints(jmbg string, points int) error
}

type mupServiceImpl struct {
	serviceUrl string
}

func NewMupService(url string) MupService {
	return mupServiceImpl{serviceUrl: url}
}

func (m mupServiceImpl) SendPoints(jmbg string, points int) (err error) {
	u, err := url.Parse(m.serviceUrl + "/api/driving_licenses/" + jmbg)
	if err != nil {
		return
	}
	dto := data.BodoviDto{
		OduzimanjeVozacke: false,
		OduzimanjeBodova:  points,
	}
	json, err := json.Marshal(dto)
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

func (m mupServiceImpl) SendKradjaPrijava(prijava data.PrijavaKradjeVozila) error {
	//TODO implement me
	return nil
}

func (m mupServiceImpl) GetVozacka(jmbg string) (vozacka *data.VozackaDozvola, err error) {
	u, err := url.Parse(m.serviceUrl + "/api/driving_licenses/" + jmbg)
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
		err = json.NewDecoder(response.Body).Decode(&vozacka)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	err = errors.New("can not reach mup service")
	return
}

func (m mupServiceImpl) GetSaobracjana(tablica string) (saobracjana *data.SaobracjanaDozvola, err error) {
	u, err := url.Parse(m.serviceUrl + "/api/vehicles/" + tablica)
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
		err = json.NewDecoder(response.Body).Decode(&saobracjana)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	err = errors.New("can not reach mup service")
	return
}
