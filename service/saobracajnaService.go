package service

import (
	"EuprvaSaobracajna/data"
	"errors"
	"log"
	"time"
)

type SaobracajnaService interface {
	GetGradjaninPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error)
	GetPolcajacPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error)
	GetPolcajacSudskeNaloge(JMBG string) ([]data.SudskiNalogDTO, error)

	GetStanice() ([]data.PolicijskaStanica, error)
	SendKradjaPrijava(prijava data.PrijavaKradjeVozila) error
	SaveNalog(nalog data.PrekrsajniNalog) (*data.PrekrsajniNalogDTO, error)
	GetPdfNalog(nalogId string) (*data.FileDto, error)
	GetVozacka(jmbg string) (*data.VozackaDozvola, error)
	GetSaobracjana(tablica string) (*data.SaobracjanaDozvola, error)
	SaveSudskiNalog(nalog data.SudskiNalog) (*data.SudskiNalog, error)
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

func (s saobracjanaServiceImpl) GetPolcajacSudskeNaloge(JMBG string) ([]data.SudskiNalogDTO, error) {
	return s.saobracjanaRepo.GetPolicajacSudskeNaloge(JMBG)
}

func (s saobracjanaServiceImpl) SaveNalog(noviNalog data.PrekrsajniNalog) (*data.PrekrsajniNalogDTO, error) {
	noviNalog.Datum = time.Now()
	nalog, err := s.saobracjanaRepo.SaveNalog(noviNalog)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
		return nil, errors.New("Cant find nalog with specified id")
	}
	pdf, err := GeneratePdf(*nalog)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with generating pdf")
	}
	fileDto, err := s.FileService.SavePdf(pdf)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with saving pdf")
	}
	return &fileDto, nil
}

func (s saobracjanaServiceImpl) GetVozacka(jmbg string) (*data.VozackaDozvola, error) {
	vozacka, err := s.MupService.GetVozacka(jmbg)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem while getting vozacka from mup")
	}
	return vozacka, nil
}

func (s saobracjanaServiceImpl) GetSaobracjana(tablica string) (*data.SaobracjanaDozvola, error) {
	saobracjana, err := s.MupService.GetSaobracjana(tablica)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem while getting saobracajna from mup")
	}
	return saobracjana, nil
}

func (s saobracjanaServiceImpl) SaveSudskiNalog(nalog data.SudskiNalog) (*data.SudskiNalog, error) {
	zaposleni, err := s.saobracjanaRepo.GetZaposleni(nalog.JMBGSluzbenika)
	if err != nil {
		log.Fatal(err.Error())
		return nil, errors.New("There is no policeman with that jmbg")
	}
	nalog.Datum = time.Now()
	err = s.sudService.SendNalog(nalog, zaposleni.RadiU.Opstina.PTT)
	if err != nil {
		log.Fatal(err.Error())
		return nil, errors.New("There is problem with sending nalog to sudService")
	}
	nalog.StatusSlucaja = "U_PROCESU"
	savedNalog, err := s.saobracjanaRepo.SaveSudskiNalog(nalog)
	if err != nil {
		log.Fatal(err.Error())
		return nil, errors.New("There is problem with saving nalog to db")
	}
	return savedNalog, nil
}
