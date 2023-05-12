package data

import "time"

/*Secret is struct that contains the secret key used for signing JWT tokens and ExpiresAt that represents until when JWT would be used and valid
 */
type Secret struct {
	Secret    string     `json:"secret"`
	ExpiresAt CustomTime `json:"expiresAt"`
}

type SudStatusDTO struct {
	Status string `json:"status"`
}

type DokaziDTO struct {
	Dokumenti []string `json:"dokumenti"`
}

type FileDto struct {
	Name string `json:"name"`
}

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	// Custom parsing logic for  date format
	// Example: "2006-01-02T15:04:05Z"
	parsedTime, err := time.Parse(`"2006-01-02T15:04:05Z"`, string(b))
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

type Opstina struct {
	PTT   string `json:"PTT"`
	Naziv string `json:"Naziv"`
}

type PolicijskaStanica struct {
	Id              string  `json:"id"`
	Adresa          string  `json:"adresa"`
	BrojTelefona    string  `json:"brojTelefona"`
	Email           string  `json:"email"`
	VremeOtvaranja  string  `json:"vremeOtvaranja"`
	VremeZatvaranja string  `json:"vremeZatvaranja"`
	Opstina         Opstina `json:"opstina"`
}

type Zaposleni struct {
	JMBG  string            `json:"JMBG"`
	RadiU PolicijskaStanica `json:"RadiU"`
}

type PrekrsajniNalog struct {
	Id             int64     `json:"id"`
	Datum          time.Time `json:"datum"`
	Opis           string    `json:"opis", required`
	IzdatoOdStrane string    `json:"izdatoOdStrane",required max=60`
	JMBGSluzbenika string    `json:"JMBGSluzbenika", len=13`
	IzdatoZa       string    `json:"izdatoZa",max = 60,required`
	JMBGZapisanog  string    `json:"JMBGZapisanog", len=13`
	TipPrekrsaja   string    `json:"tipPrekrsaja", required`
	JedinicaMere   *string   `json:"jedinicaMere"`
	Vrednost       *float64  `json:"vrednost", min=0`
	KaznaIzvrsena  bool      `json:"kaznaIzvrsena"`
	Slike          []string  `json:"slike"`
}

type PrekrsajniNalogDTO struct {
	Id             int64     `json:"id"`
	Datum          time.Time `json:"datum"`
	Opis           string    `json:"opis"`
	IzdatoOdStrane string    `json:"izdatoOdStrane"`
	IzdatoZa       string    `json:"izdatoZa"`
	JMBGZapisanog  string    `json:"JMBGZapisanog"`
	TipPrekrsaja   string    `json:"tipPrekrsaja"`
	JedinicaMere   *string   `json:"jedinicaMere"`
	Vrednost       *float64  `json:"vrednost"`
	Slike          []string  `json:"slike"`
	KaznaIzvrsena  bool      `json:"kaznaIzvrsena"`
}

type SudskiNalog struct {
	Id             int64     `json:"id"`
	Datum          time.Time `json:"datum"`
	Naslov         string    `json:"naslov"`
	Opis           string    `json:"opis"`
	IzdatoOdStrane string    `json:"izdatoOdStrane"`
	JMBGSluzbenika string    `json:"JMBGSluzbenika"`
	Optuzeni       string    `json:"Optuzeni"`
	JMBGoptuzenog  string    `json:"JMBGoptuzenog"`
	StatusSlucaja  string    `json:"statusSlucaja"`
	Dokumenti      []string  `json:"dokumenti"`
}

type SudskiNalogDTO struct {
	Id            int64     `json:"id"`
	Datum         time.Time `json:"datum"`
	Naslov        string    `json:"naslov"`
	Opis          string    `json:"opis"`
	Optuzeni      string    `json:"optuzeni"`
	JMBGoptuzenog string    `json:"JMBGoptuzenog"`
	StatusSlucaja string    `json:"statusSlucaja"`
	Dokumenti     []string  `json:"dokumenti"`
}

type PrijavaKradjeVozila struct {
	Prijvio          string    `json:"prijvio", max=60`
	KontaktTelefon   string    `json:"kontaktTelefon",min=10,max=13`
	BrojRegistracije string    `json:"brojRegistracije"max=7,min=3`
	Datum            time.Time `json:"datum"`
	JMBGVlasnika     string    `json:"JMBGVlasnika", len=13`
}

type VozackaDozvola struct {
	BrojVozackeDozvole   string    `json:"brojVozackeDozvole"`
	KategorijeVozila     []string  `json:"kategorijeVozila"`
	DatumIzdavavanja     time.Time `json:"datumIzdavavanja"`
	DatumIsteka          time.Time `json:"datumIsteka"`
	BrojKaznenihPoena    int       `json:"brojKaznenihPoena"`
	StatusVozackeDozvole string    `json:"statusVozackeDozvole"`
}

type SaobracjanaDozvola struct {
	Marka              string     `json:"marka"`
	Model              string     `json:"model"`
	GodinaProizvodnje  int        `json:"godinaProizvodnje"`
	Boja               string     `json:"boja"`
	RegBroj            string     `json:"regBroj"`
	SnagaMotora        float64    `json:"snagaMotora"`
	MaksimalnaBrzina   float64    `json:"maksimalnaBrzina"`
	BrojSedista        int        `json:"brojSedista"`
	Tezina             float64    `json:"tezina"`
	TipVozila          string     `json:"tipVozila"`
	StatusRegistracije string     `json:"statusRegistracije"`
	PrijavljenaKradja  *time.Time `json:"prijavljenaKradja",omitempty`
}
