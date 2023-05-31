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
	UpdateSudNalogStatus(id string, status data.SudStatusDTO) error
	GetPolicajacNeIzvrseniNalozi(jmbg string) ([]data.PrekrsajniNalogDTO, error)
	UpdatePrekrsajNalogIzvrsen(id, tokenSluzbenika string) error
	GetAktivniSlucajvei(jmbg string) ([]data.SudskiSlucaj, error)
}

type saobracjanaServiceImpl struct {
	saobracjanaRepo data.SaobracajnaRepo
	sudService      SudService
	mupService      MupService
	fileService     FilesService
	jwtService      JwtService
}

func NewSaobracjanaService(repo data.SaobracajnaRepo, ms MupService, ss SudService, fs FilesService, js JwtService) SaobracajnaService {
	return saobracjanaServiceImpl{saobracjanaRepo: repo, sudService: ss, mupService: ms, fileService: fs, jwtService: js}
}

func (s saobracjanaServiceImpl) SendKradjaPrijava(prijava data.PrijavaKradjeVozila) error {
	return s.mupService.SendKradjaPrijava(prijava)
}

func (s saobracjanaServiceImpl) GetPolcajacPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error) {
	return s.saobracjanaRepo.GetPolcajacPrekrsajneNaloge(JMBG)
}

func (s saobracjanaServiceImpl) GetPolicajacNeIzvrseniNalozi(JMBG string) ([]data.PrekrsajniNalogDTO, error) {
	return s.saobracjanaRepo.GetPolcajacNeIzvrsenePrekrsajneNaloge(JMBG)
}

func (s saobracjanaServiceImpl) GetGradjaninPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error) {
	return s.saobracjanaRepo.GetGradjaninPrekrsajneNaloge(JMBG)
}

func (s saobracjanaServiceImpl) GetPolcajacSudskeNaloge(JMBG string) ([]data.SudskiNalogDTO, error) {
	return s.saobracjanaRepo.GetPolicajacSudskeNaloge(JMBG)
}

func (s saobracjanaServiceImpl) SaveNalog(noviNalog data.PrekrsajniNalog) (*data.PrekrsajniNalogDTO, error) {
	noviNalog.Datum = time.Now()
	if noviNalog.KaznaIzvrsena {
		points := CalculatePointsForTicket(noviNalog)
		if points != 0 {
			err := s.mupService.SendPoints(points)
			if err != nil {
				log.Fatal(err)
				return nil, errors.New("There was problem while sending nalog points to MUP")
			}
		}
	}
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
	nalog, err := s.saobracjanaRepo.GetPrekrsajniNalog(nalogId)
	if err != nil || nalog == nil {
		log.Fatal(err)
		return nil, errors.New("Cant find nalog with specified id")
	}
	pdf, err := GeneratePdf(*nalog)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with generating pdf")
	}
	fileDto, err := s.fileService.SavePdf(pdf)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with saving pdf")
	}
	return &fileDto, nil
}

func (s saobracjanaServiceImpl) GetVozacka(jmbg string) (*data.VozackaDozvola, error) {
	vozacka, err := s.mupService.GetVozacka(jmbg)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem while getting vozacka from mup")
	}
	return vozacka, nil
}

func (s saobracjanaServiceImpl) GetSaobracjana(tablica string) (*data.SaobracjanaDozvola, error) {
	saobracjana, err := s.mupService.GetSaobracjana(tablica)
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
	nalog.OpstinaPTT = zaposleni.RadiU.Opstina.PTT
	err = s.sudService.SendNalog(nalog)
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

func (s saobracjanaServiceImpl) UpdateSudNalogStatus(id string, status data.SudStatusDTO) error {
	SlucajEnumMap := map[int]string{
		data.Odbijen:   "ODBIJEN",
		data.Presudjen: "PRESUDJEN",
	}
	err := s.saobracjanaRepo.UpdateSudNalogStatus(id, SlucajEnumMap[status.Status])
	if err != nil {
		log.Fatal(err.Error())
		return errors.New("There is problem with saving nalog status to db")
	}
	return nil
}

func (s saobracjanaServiceImpl) UpdatePrekrsajNalogIzvrsen(id, tokenSluzbenika string) error {
	nalog, err := s.saobracjanaRepo.GetPrekrsajniNalog(id)
	if err != nil || nalog == nil {
		log.Fatal(err)
		return errors.New("Cant find nalog with specified id")
	}
	jmbgZaposlenog, _ := GetPrincipal(tokenSluzbenika, s.jwtService.GetSecret().Secret)
	if nalog.JMBGSluzbenika != jmbgZaposlenog {
		return errors.New("Cant update someone elsess nalog")
	}
	err = s.saobracjanaRepo.UpdatePrekrsajNalogIzvrsen(id)
	if err != nil {
		log.Fatal(err.Error())
		return errors.New("There is problem with saving nalog status to db")
	}
	return nil
}

func (s saobracjanaServiceImpl) GetAktivniSlucajvei(jmbg string) ([]data.SudskiSlucaj, error) {
	slucajevi, err := s.sudService.GetGradjaninSlucajevi(jmbg)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There was problem while getting aktivni slucajevi")
	}
	return slucajevi, nil
}
