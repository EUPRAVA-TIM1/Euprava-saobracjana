package service

import "EuprvaSaobracajna/data"

type MupService interface {
	SendKradjaPrijava(prijava data.PrijavaKradjeVozila) error
}

type mupServiceImpl struct {
	serviceUrl string
}

func NewMupService(url string) MupService {
	return mupServiceImpl{serviceUrl: url}
}

func (m mupServiceImpl) SendKradjaPrijava(prijava data.PrijavaKradjeVozila) error {
	//TODO implement me
	panic("implement me")
}
