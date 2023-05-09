package service

import (
	"EuprvaSaobracajna/data"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type SsoService interface {
	GetSecret() (data.Secret, error)
}

const issuerHeader = "X-Service-Name"

type ssoServiceImpl struct {
	serviceUrl string
	issuer     string
}

func NewSsoService(url, issuer string) SsoService {
	return ssoServiceImpl{serviceUrl: url, issuer: issuer}
}

func (s ssoServiceImpl) GetSecret() (secret data.Secret, err error) {
	u, err := url.Parse(s.serviceUrl + "/sso/Secret")
	if err != nil {
		return
	}

	client := http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Add(issuerHeader, s.issuer)

	response, err := client.Do(req)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&secret)
		return
	}
	err = errors.New("can not reach SSOService")
	return
}
