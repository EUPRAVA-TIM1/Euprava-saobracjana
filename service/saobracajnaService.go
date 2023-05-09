package service

import (
	"EuprvaSaobracajna/data"
	"errors"
	"time"
)

type SaobracajnaService interface {
	GetGradjaninPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error)
	GetPolcajacPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error)
	GetStanice() ([]data.PolicijskaStanica, error)
	SendKradjaPrijava(prijava data.PrijavaKradjeVozila) error
	SaveNalog(nalog data.PrekrsajniNalog) (*data.PrekrsajniNalogDTO, error)
	GetPdfNalog(nalogId string) (*data.FileDto, error)
}

type saobracjanaServiceImpl struct {
	saobracjanaRepo data.SaobracajnaRepo
	sudService      SudService
	MupService      MupService
	FileService     FilesService
}

func NewSaobracjanaService(repo data.SaobracajnaRepo, ms MupService, ss SudService, fs FilesService) SaobracajnaService {
	return saobracjanaServiceImpl{saobracjanaRepo: repo, sudService: ss, MupService: ms, FileService: fs}
}

func (s saobracjanaServiceImpl) SendKradjaPrijava(prijava data.PrijavaKradjeVozila) error {
	return s.MupService.SendKradjaPrijava(prijava)
}

func (s saobracjanaServiceImpl) GetPolcajacPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error) {
	return s.saobracjanaRepo.GetPolcajacPrekrsajneNaloge(JMBG)

}

func (s saobracjanaServiceImpl) GetGradjaninPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error) {
	return s.saobracjanaRepo.GetGradjaninPrekrsajneNaloge(JMBG)
}
func (s saobracjanaServiceImpl) SaveNalog(noviNalog data.PrekrsajniNalog) (*data.PrekrsajniNalogDTO, error) {
	noviNalog.Datum = time.Now()
	nalog, err := s.saobracjanaRepo.SaveNalog(noviNalog)
	if err != nil {
		return nil, errors.New("There was problem while saving nalog")
	}
	return &data.PrekrsajniNalogDTO{
		Id:             nalog.Id,
		Datum:          nalog.Datum,
		Opis:           nalog.Opis,
		IzdatoOdStrane: nalog.IzdatoOdStrane,
		IzdatoZa:       nalog.IzdatoZa,
		JMBGZapisanog:  nalog.JMBGZapisanog,
		TipPrekrsaja:   nalog.TipPrekrsaja,
		JedinicaMere:   nalog.JedinicaMere,
		Vrednost:       nalog.Vrednost,
		Slike:          nalog.Slike,
	}, nil
}

func (s saobracjanaServiceImpl) GetStanice() ([]data.PolicijskaStanica, error) {
	return s.saobracjanaRepo.GetStanice()
}

func (s saobracjanaServiceImpl) GetPdfNalog(nalogId string) (*data.FileDto, error) {
	nalog, err := s.saobracjanaRepo.GetPrekrajniNalog(nalogId)
	if err != nil || nalog == nil {
		return nil, errors.New("Cant find nalog with specified id")
	}
	pdf, err := GeneratePdf(*nalog)
	if err != nil {
		return nil, errors.New("There has been problem with generating pdf")
	}
	fileDto, err := s.FileService.SavePdf(pdf)
	if err != nil {
		return nil, errors.New("There has been problem with saving pdf")
	}
	return &fileDto, nil
}
