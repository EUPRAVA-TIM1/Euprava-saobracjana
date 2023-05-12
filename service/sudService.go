package service

import "EuprvaSaobracajna/data"

type SudService interface {
	SendNalog(nalog data.SudskiNalog, PTT string) error
	SendDokazi(idNaloga string, dto data.DokaziDTO) error
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

func (s sudServiceImpl) SendDokazi(idNaloga string, dto data.DokaziDTO) error {
	//TODO implement me
	return nil
}
