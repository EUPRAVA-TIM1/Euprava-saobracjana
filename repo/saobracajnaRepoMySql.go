package repo

import (
	"EuprvaSaobracajna/data"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const schema = "eupravaMilicija"

type SaobracjanaRepoSql struct {
	pass string
	host string
	port string
}

/*
NewGrdjaninRepoSql
generates new SaobracajnaRepo struct and accepts port , pass and Host strings in that order
*/
func NewSaobracjanaRepoSql(port, pass, host string) data.SaobracajnaRepo {
	return &SaobracjanaRepoSql{
		pass: pass,
		host: host,
		port: port,
	}
}

func (s SaobracjanaRepoSql) GetPolcajacPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error) {
	db, err := s.OpenConnection()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()
	query := "SELECT Id,Datum,Opis,IzdatoOdStrane,IzdatoZa,JMBGZapisanog,TipPrekrsaja,JedinicaMere,Vrednost,KaznaIzvrsena FROM PrekrsajniNalog where JMBGSluzbenika = ?;"
	rows, err := db.Query(query, JMBG)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with reading nalog from db")
	}
	nalozi := make([]data.PrekrsajniNalogDTO, 0)
	for rows.Next() {
		var nalog data.PrekrsajniNalogDTO
		var dateStr string
		err := rows.Scan(&nalog.Id, &dateStr, &nalog.Opis, &nalog.IzdatoOdStrane, &nalog.IzdatoZa, &nalog.JMBGZapisanog, &nalog.TipPrekrsaja, &nalog.JedinicaMere, &nalog.Vrednost, &nalog.KaznaIzvrsena)
		if err != nil {
			log.Fatal(err)
		}
		datum, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
		}
		nalog.Datum = datum
		var imgs = make([]string, 0)
		imgQuery := "select UrlSlike from SlikeNaloga where NalogId = ?;"
		imgRows, err := db.Query(imgQuery, nalog.Id)
		if err != nil {
			log.Fatal(err)
			return nil, errors.New("There has been problem with reading imgs from db")
		}
		for imgRows.Next() {
			var url string
			err := imgRows.Scan(&url)
			if err != nil {
				log.Fatal(err)
			}
			imgs = append(imgs, url)
		}
		nalog.Slike = imgs
		nalozi = append(nalozi, nalog)
	}
	return nalozi, nil
}

func (s SaobracjanaRepoSql) GetPolcajacNeIzvrsenePrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error) {
	db, err := s.OpenConnection()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()
	query := "SELECT Id,Datum,Opis,IzdatoOdStrane,IzdatoZa,JMBGZapisanog,TipPrekrsaja,JedinicaMere,Vrednost,KaznaIzvrsena FROM PrekrsajniNalog where JMBGSluzbenika = ? and KaznaIzvrsena = false;"
	rows, err := db.Query(query, JMBG)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with reading nalog from db")
	}
	nalozi := make([]data.PrekrsajniNalogDTO, 0)
	for rows.Next() {
		var nalog data.PrekrsajniNalogDTO
		var dateStr string
		err := rows.Scan(&nalog.Id, &dateStr, &nalog.Opis, &nalog.IzdatoOdStrane, &nalog.IzdatoZa, &nalog.JMBGZapisanog, &nalog.TipPrekrsaja, &nalog.JedinicaMere, &nalog.Vrednost, &nalog.KaznaIzvrsena)
		if err != nil {
			log.Fatal(err)
		}
		datum, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
		}
		nalog.Datum = datum
		var imgs = make([]string, 0)
		imgQuery := "select UrlSlike from SlikeNaloga where NalogId = ?;"
		imgRows, err := db.Query(imgQuery, nalog.Id)
		if err != nil {
			log.Fatal(err)
			return nil, errors.New("There has been problem with reading imgs from db")
		}
		for imgRows.Next() {
			var url string
			err := imgRows.Scan(&url)
			if err != nil {
				log.Fatal(err)
			}
			imgs = append(imgs, url)
		}
		nalog.Slike = imgs
		nalozi = append(nalozi, nalog)
	}
	return nalozi, nil
}

func (s SaobracjanaRepoSql) GetPolicajacSudskeNaloge(JMBG string) ([]data.SudskiNalogDTO, error) {
	db, err := s.OpenConnection()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()
	query := "SELECT Id,Datum,Naslov,Opis,Optuzeni,JMBGOptuzenog,StatusSlucaja FROM SudskiNalog where JMBGSluzbenika = ?;"
	rows, err := db.Query(query, JMBG)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with reading nalog from db")
	}
	nalozi := make([]data.SudskiNalogDTO, 0)
	for rows.Next() {
		var nalog data.SudskiNalogDTO
		var dateStr string
		err := rows.Scan(&nalog.Id, &dateStr, &nalog.Naslov, &nalog.Opis, &nalog.Optuzeni, &nalog.JMBGoptuzenog, &nalog.StatusSlucaja)
		if err != nil {
			log.Fatal(err)
		}
		datum, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
		}
		nalog.Datum = datum
		var files = make([]string, 0)
		fileQuery := "select UrlDokumenta from DokumentiSudskogNaloga where NalogId = ?;"
		fileRows, err := db.Query(fileQuery, nalog.Id)
		if err != nil {
			log.Fatal(err)
			return nil, errors.New("There has been problem with reading imgs from db")
		}
		for fileRows.Next() {
			var url string
			err := fileRows.Scan(&url)
			if err != nil {
				log.Fatal(err)
			}
			files = append(files, url)
		}
		nalog.Dokumenti = files
		nalozi = append(nalozi, nalog)
	}
	return nalozi, nil
}

func (s SaobracjanaRepoSql) GetPrekrsajniNalog(nalogId string) (*data.PrekrsajniNalog, error) {
	db, err := s.OpenConnection()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()
	query := "SELECT Id,Datum,Opis,IzdatoOdStrane,JMBGSluzbenika,IzdatoZa,JMBGZapisanog,TipPrekrsaja,JedinicaMere,Vrednost,KaznaIzvrsena FROM PrekrsajniNalog where id = ?;"
	rows, err := db.Query(query, nalogId)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with reading nalog from db")
	}
	var nalog data.PrekrsajniNalog
	for rows.Next() {
		var dateStr string
		err := rows.Scan(&nalog.Id, &dateStr, &nalog.Opis, &nalog.IzdatoOdStrane, &nalog.JMBGSluzbenika, &nalog.IzdatoZa, &nalog.JMBGZapisanog, &nalog.TipPrekrsaja, &nalog.JedinicaMere, &nalog.Vrednost, &nalog.KaznaIzvrsena)
		if err != nil {
			log.Fatal(err)
		}
		datum, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Fatal(err)
		}
		nalog.Datum = datum
		var imgs = make([]string, 0)
		imgQuery := "select UrlSlike from SlikeNaloga where NalogId = ?;"
		imgRows, err := db.Query(imgQuery, nalog.Id)
		if err != nil {
			log.Fatal(err)
			return nil, errors.New("There has been problem with reading imgs from db")
		}
		for imgRows.Next() {
			var url string
			err := imgRows.Scan(&url)
			if err != nil {
				log.Fatal(err)
			}
			imgs = append(imgs, url)
		}
		nalog.Slike = imgs
	}
	return &nalog, nil
}

func (s SaobracjanaRepoSql) GetGradjaninPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error) {
	db, err := s.OpenConnection()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()
	query := "SELECT Id,Datum,Opis,IzdatoOdStrane,IzdatoZa,JMBGZapisanog,TipPrekrsaja,JedinicaMere,Vrednost,KaznaIzvrsena FROM PrekrsajniNalog where JMBGZapisanog = ?;"
	rows, err := db.Query(query, JMBG)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with reading nalog from db")
	}
	nalozi := make([]data.PrekrsajniNalogDTO, 0)
	for rows.Next() {
		var nalog data.PrekrsajniNalogDTO
		var dateStr string
		err := rows.Scan(&nalog.Id, &dateStr, &nalog.Opis, &nalog.IzdatoOdStrane, &nalog.IzdatoZa, &nalog.JMBGZapisanog, &nalog.TipPrekrsaja, &nalog.JedinicaMere, &nalog.Vrednost, &nalog.KaznaIzvrsena)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
		}
		datum, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
		}
		nalog.Datum = datum
		var imgs = make([]string, 0)
		imgQuery := "select UrlSlike from SlikeNaloga where NalogId = ?;"
		imgRows, err := db.Query(imgQuery, nalog.Id)
		if err != nil {
			log.Fatal(err)
			return nil, errors.New("There has been problem with reading imgs from db")
		}
		for imgRows.Next() {
			var url string
			err := imgRows.Scan(&url)
			if err != nil {
				log.Fatal(err)
			}
			imgs = append(imgs, url)
		}
		nalog.Slike = imgs
		nalozi = append(nalozi, nalog)
	}
	return nalozi, nil
}

func (s SaobracjanaRepoSql) GetStanice() ([]data.PolicijskaStanica, error) {
	db, err := s.OpenConnection()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()

	query := "select Id,Adresa,BrojTelefona,Email,VremeOtvaranja,VremeZatvaranja,Naziv,PTT from PolicijskaStanica p, Opstina o where o.PTT = p.OpstinaID ;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with reading from db")
	}

	stanice := make([]data.PolicijskaStanica, 0)
	for rows.Next() {
		var stanica data.PolicijskaStanica
		err := rows.Scan(&stanica.Id, &stanica.Adresa, &stanica.BrojTelefona, &stanica.Email, &stanica.VremeOtvaranja, &stanica.VremeZatvaranja, &stanica.Opstina.Naziv, &stanica.Opstina.PTT)
		if err != nil {
			log.Fatal(err)
		}
		stanice = append(stanice, stanica)
	}
	return stanice, nil
}

func (s SaobracjanaRepoSql) IsAWorker(jmbg string) (bool, error) {
	db, err := s.OpenConnection()
	if err != nil {
		log.Fatal(err)
		return false, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()

	query := "select count(*) from Zaposleni where JMBG = ? ;"
	var count int
	err = db.QueryRow(query, jmbg).Scan(&count)
	if err != nil {
		log.Fatal(err)
		return false, errors.New("There has been problem with reading from db")
	}
	return count > 0, nil
}

func (s SaobracjanaRepoSql) GetZaposleni(jmbg string) (*data.Zaposleni, error) {
	db, err := s.OpenConnection()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()

	query := "select JMBG,ps.Id,Adresa,BrojTelefona,Email,VremeOtvaranja,VremeZatvaranja,Ptt,Naziv  from Zaposleni z ,Opstina o,PolicijskaStanica ps where z.RadiU = ps.Id and ps.OpstinaID = o.PTT and z.jmbg = ?;"
	rows, err := db.Query(query, jmbg)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with reading zaposleni from db")
	}
	var zaposleni data.Zaposleni
	for rows.Next() {
		err := rows.Scan(&zaposleni.JMBG, &zaposleni.RadiU.Id, &zaposleni.RadiU.Adresa, &zaposleni.RadiU.BrojTelefona, &zaposleni.RadiU.Email, &zaposleni.RadiU.VremeOtvaranja,
			&zaposleni.RadiU.VremeZatvaranja, &zaposleni.RadiU.Opstina.PTT, &zaposleni.RadiU.Opstina.Naziv)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &zaposleni, nil
}

func (s SaobracjanaRepoSql) SaveNalog(nalog data.PrekrsajniNalog) (*data.PrekrsajniNalog, error) {
	db, err := s.OpenConnection()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()

	query := "INSERT INTO PrekrsajniNalog ( Datum, Opis, IzdatoOdStrane, JMBGSluzbenika, IzdatoZa, JMBGZapisanog, TipPrekrsaja, JedinicaMere, Vrednost,KaznaIzvrsena) VALUES (?,?,?,?,?,?,?,?,?,?)"
	res, err := db.Exec(query, nalog.Datum, nalog.Opis, nalog.IzdatoOdStrane, nalog.JMBGSluzbenika, nalog.IzdatoZa, nalog.JMBGZapisanog, nalog.TipPrekrsaja, nalog.JedinicaMere, nalog.Vrednost, nalog.KaznaIzvrsena)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("failed to insert secret key: %v", err)
	}
	id, _ := res.LastInsertId()
	if len(nalog.Slike) > 0 {
		for i := 0; i < len(nalog.Slike); i++ {
			query := "INSERT INTO SlikeNaloga(NalogId,UrlSlike) VALUES (?,?)"
			_, err := db.Exec(query, id, nalog.Slike[i])
			if err != nil {
				log.Fatal(err)
				query := "DELETE FROM PrekrsajniNalog where Id = ?"
				db.Exec(query, id)
				return nil, fmt.Errorf("failed to insert secret key: %v", err)
			}
		}
	}

	nalog.Id = id
	return &nalog, nil
}

func (s SaobracjanaRepoSql) SaveSudskiNalog(nalog data.SudskiNalog) (*data.SudskiNalog, error) {
	db, err := s.OpenConnection()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()

	query := "INSERT INTO SudskiNalog ( Datum, Naslov, Opis, IzdatoOdStrane, JMBGSluzbenika, Optuzeni, JMBGOptuzenog, StatusSlucaja) VALUES (?,?,?,?,?,?,?,?)"
	res, err := db.Exec(query, nalog.Datum, nalog.Naslov, nalog.Opis, nalog.IzdatoOdStrane, nalog.JMBGSluzbenika, nalog.Optuzeni, nalog.JMBGoptuzenog, nalog.StatusSlucaja)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("failed to insert secret key: %v", err)
	}
	id, _ := res.LastInsertId()
	if len(nalog.Dokumenti) > 0 {
		for i := 0; i < len(nalog.Dokumenti); i++ {
			query := "INSERT INTO DokumentiSudskogNaloga(NalogId,UrlDokumenta) VALUES (?,?)"
			_, err := db.Exec(query, id, nalog.Dokumenti[i])
			if err != nil {
				log.Fatal(err)
				query := "DELETE FROM SudskiNalog where Id = ?"
				db.Exec(query, id)
				return nil, fmt.Errorf("failed to insert secret key: %v", err)
			}
		}
	}

	nalog.Id = id
	return &nalog, nil
}

func (s SaobracjanaRepoSql) UpdateSudNalogStatus(id, status string) error {
	db, err := s.OpenConnection()
	if err != nil {
		log.Fatal(err)
		return errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()

	query := "UPDATE SudskiNalog SET StatusSlucaja = ? WHERE Id = ?"
	_, err = db.Exec(query, status, id)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("failed to insert secret key: %v", err)
	}
	return nil
}

func (s SaobracjanaRepoSql) OpenConnection() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s:%s)/%s", s.pass, s.host, s.port, schema))
}
