package service

import (
	"EuprvaSaobracajna/data"
	"log"
	"time"
)

type JwtService interface {
	Validate(jwt string) bool
	IsAWorker(jwt string) bool
	GetSecret() *data.Secret
}

type jwtServiceImpl struct {
	repo            data.SecretRepo
	saobracjanaRepo data.SaobracajnaRepo
	ssoService      SsoService
}

func NewJwtService(repo data.SecretRepo, saobracajnaRepo data.SaobracajnaRepo, service SsoService) JwtService {
	return jwtServiceImpl{repo: repo, saobracjanaRepo: saobracajnaRepo, ssoService: service}
}

func (j jwtServiceImpl) Validate(jwt string) bool {
	secret := j.GetSecret()
	err := ValidateJwt(jwt, secret.Secret)

	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}

func (j jwtServiceImpl) IsAWorker(jwt string) bool {
	secret := j.GetSecret()
	jmbg, err := GetPrincipal(jwt, secret.Secret)
	if err != nil {
		panic(err)
		return false
	}
	isAWorker, err := j.saobracjanaRepo.IsAWorker(jmbg)
	if err != nil {
		panic(err)
		return false
	}
	return isAWorker
}

func (j jwtServiceImpl) GetSecret() *data.Secret {
	secret, err := j.repo.GetSecret()
	if err != nil {
		panic(err)
		return nil
	}
	if time.Now().After(secret.ExpiresAt.Time) {
		secret, err := j.ssoService.GetSecret()
		if err != nil {
			panic(err)
		}
		err = j.repo.UpdateSecret(secret)
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		panic(err)
		return nil
	}
	return secret
}
