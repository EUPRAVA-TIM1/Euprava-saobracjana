package data

type SaobracajnaRepo interface {
	GetGradjaninPrekrsajneNaloge(JMBG string) ([]PrekrsajniNalogDTO, error)
	GetPolcajacPrekrsajneNaloge(JMBG string) ([]PrekrsajniNalogDTO, error)
	GetPolcajacNeIzvrsenePrekrsajneNaloge(JMBG string) ([]PrekrsajniNalogDTO, error)
	UpdatePrekrsajNalogIzvrsen(id string) error
	GetPolicajacSudskeNaloge(JMBG string) ([]SudskiNalogDTO, error)
	GetStanice() ([]PolicijskaStanica, error)
	IsAWorker(jmbg string) (bool, error)
	SaveNalog(nalog PrekrsajniNalog) (*PrekrsajniNalog, error)
	GetPrekrsajniNalog(nalogId string) (*PrekrsajniNalog, error)
	GetZaposleni(jmbg string) (*Zaposleni, error)
	SaveSudskiNalog(nalog SudskiNalog) (*SudskiNalog, error)
	UpdateSudNalogStatus(id, status string) error
	RemoveSudskiNalog(id int64) error
}
