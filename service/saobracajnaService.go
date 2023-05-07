package service

import "EuprvaSaobracajna/data"

type SaobracajnaService interface {
	GetGradjaninPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error)
	GetStanice() ([]data.PolicijskaStanica, error)
}

type saobracjanaServiceImpl struct {
	saobracjanaRepo data.SaobracajnaRepo
}

func NewSaobracjanaService(repo data.SaobracajnaRepo) SaobracajnaService {
	return saobracjanaServiceImpl{repo}
}

func (s saobracjanaServiceImpl) GetGradjaninPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error) {
	return s.saobracjanaRepo.GetGradjaninPrekrsajneNaloge(JMBG)
}

func (s saobracjanaServiceImpl) GetStanice() ([]data.PolicijskaStanica, error) {
	return s.saobracjanaRepo.GetStanice()
}
