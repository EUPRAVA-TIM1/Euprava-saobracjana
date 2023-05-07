package data

import "time"

/*Secret is struct that contains the secret key used for signing JWT tokens and ExpiresAt that represents until when JWT would be used and valid
 */
type Secret struct {
	Secret    string     `json:"secret"`
	ExpiresAt CustomTime `json:"expiresAt"`
}

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	// Custom parsing logic for your date format
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
	Id             int       `json:"id"`
	Datum          time.Time `json:"datum"`
	Opis           string    `json:"opis"`
	IzdatoOdStrane string    `json:"izdatoOdStrane"`
	JMBGSluzbenika string    `json:"JMBGSluzbenika"`
	IzdatoZa       string    `json:"izdatoZa"`
	JMBGZapisanog  string    `json:"JMBGZapisanog"`
	TipPrekrsaja   string    `json:"tipPrekrsaja"`
	JedinicaMere   *string   `json:"jedinicaMere"`
	Vrednost       *int      `json:"vrednost"`
	Slike          []string  `json:"slike"`
}

type PrekrsajniNalogDTO struct {
	Id             int       `json:"id"`
	Datum          time.Time `json:"datum"`
	Opis           string    `json:"opis"`
	IzdatoOdStrane string    `json:"izdatoOdStrane"`
	IzdatoZa       string    `json:"izdatoZa"`
	JMBGZapisanog  string    `json:"JMBGZapisanog"`
	TipPrekrsaja   string    `json:"tipPrekrsaja"`
	JedinicaMere   *string   `json:"jedinicaMere"`
	Vrednost       *int      `json:"vrednost"`
	Slike          []string  `json:"slike"`
}

type SudskiNalog struct {
	Id             int       `json:"id"`
	Datum          time.Time `json:"datum"`
	Naslov         string    `json:"naslov"`
	Opis           string    `json:"opis"`
	IzdatoOdStrane string    `json:"izdatoOdStrane"`
	JMBGSluzbenika string    `json:"JMBGSluzbenika"`
	Optuzeni       string    `json:"Optuzeni"`
	JMBGoptuzenog  string    `json:"JMBGoptuzenog"`
	StatusSlucaja  *string   `json:"statusSlucaja"`
	Dokumenti      []string  `json:"dokumenti"`
}

type PrijavaKradjeVozila struct {
	Prijvio          string    `json:"prijvio"`
	KontaktTelefon   string    `json:"kontaktTelefon"`
	BrojRegistracije string    `json:"brojRegistracije"`
	Datum            time.Time `json:"datum"`
	JMBGVlasnika     string    `json:"JMBGVlasnika"`
}
