package data

type SaobracajnaRepo interface {
	GetGradjaninPrekrsajneNaloge(JMBG string) ([]PrekrsajniNalogDTO, error)
	GetPolcajacPrekrsajneNaloge(JMBG string) ([]PrekrsajniNalogDTO, error)
	GetStanice() ([]PolicijskaStanica, error)
	IsAWorker(jmbg string) (bool, error)
	SaveNalog(nalog PrekrsajniNalog) (*PrekrsajniNalog, error)
	GetPrekrajniNalog(nalogId string) (*PrekrsajniNalog, error)
}
