package data

type SaobracajnaRepo interface {
	GetGradjaninPrekrsajneNaloge(JMBG string) ([]PrekrsajniNalogDTO, error)
	GetStanice() ([]PolicijskaStanica, error)
}
