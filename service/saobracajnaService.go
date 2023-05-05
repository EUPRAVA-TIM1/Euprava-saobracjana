package service

import "EuprvaSaobracajna/data"

type SaobracajnaService interface {
	GetGradjaninPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error)
	GetStanice() ([]data.PolicijskaStanica, error)
}

type saobracjanaService struct {
	saobracjanaRepo data.SaobracajnaRepo
}

func (s saobracjanaService) GetGradjaninPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error) {
	return s.saobracjanaRepo.GetGradjaninPrekrsajneNaloge(JMBG)
}

func (s saobracjanaService) GetStanice() ([]data.PolicijskaStanica, error) {
	return s.saobracjanaRepo.GetStanice()
}

func NewSaobracjanaService(repo data.SaobracajnaRepo) SaobracajnaService {
	return saobracjanaService{repo}
}
