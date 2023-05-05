package repo

import (
	"EuprvaSaobracajna/data"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

func (s SaobracjanaRepoSql) GetGradjaninPrekrsajneNaloge(JMBG string) ([]data.PrekrsajniNalogDTO, error) {
	db, err := s.OpenConnection()
	if err != nil {
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()
	query := "SELECT Id,Datum,Opis,IzdatoOdStrane,IzdatoZa,JMBGZapisanog,TipPrekrsaja,JedinicaMere,Vrednost FROM PrekrsajniNalog where JMBGZapisanog = ?;"
	rows, err := db.Query(query, JMBG)
	if err != nil {
		panic(err)
		return nil, errors.New("There has been problem with reading nalog from db")
	}
	nalozi := make([]data.PrekrsajniNalogDTO, 0)
	for rows.Next() {
		var nalog data.PrekrsajniNalogDTO
		var dateStr string
		err := rows.Scan(&nalog.Id, &dateStr, &nalog.Opis, &nalog.IzdatoOdStrane, &nalog.IzdatoZa, &nalog.JMBGZapisanog, &nalog.TipPrekrsaja, &nalog.JedinicaMere, &nalog.Vrednost)
		if err != nil {
			panic(err.Error())
		}
		datum, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			panic(err.Error())
		}
		nalog.Datum = datum
		var imgs = make([]string, 0)
		imgQuery := "select UrlSlike from SlikeNaloga where NalogId = ?;"
		imgRows, err := db.Query(imgQuery, nalog.Id)
		if err != nil {
			panic(err)
			return nil, errors.New("There has been problem with reading imgs from db")
		}
		for imgRows.Next() {
			var url string
			err := imgRows.Scan(&url)
			if err != nil {
				panic(err.Error())
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
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()

	query := "select Id,Adresa,BrojTelefona,Email,VremeOtvaranja,VremeZatvaranja,Naziv,PTT from PolicijskaStanica p, Opstina o where o.PTT = p.OpstinaID ;"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
		return nil, errors.New("There has been problem with reading from db")
	}

	stanice := make([]data.PolicijskaStanica, 0)
	for rows.Next() {
		var stanica data.PolicijskaStanica
		err := rows.Scan(&stanica.Id, &stanica.Adresa, &stanica.BrojTelefona, &stanica.Email, &stanica.VremeOtvaranja, &stanica.VremeZatvaranja, &stanica.Opstina.Naziv, &stanica.Opstina.PTT)
		if err != nil {
			panic(err.Error())
		}
		stanice = append(stanice, stanica)
	}
	return stanice, nil
}

func (s SaobracjanaRepoSql) OpenConnection() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s:%s)/%s", s.pass, s.host, s.port, schema))
}
