package repo

import (
	"EuprvaSaobracajna/data"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type SecretRepoSql struct {
	pass string
	host string
	port string
}

/*
NewGrdjaninRepoSql
generates new SaobracajnaRepo struct and accepts port , pass and Host strings in that order
*/
func NewSecretRepoSql(port, pass, host string) data.SecretRepo {
	return &SecretRepoSql{
		pass: pass,
		host: host,
		port: port,
	}
}

func (s SecretRepoSql) GetSecret() (*data.Secret, error) {
	db, err := s.OpenConnection()
	if err != nil {
		return nil, errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()

	query := "select SecretKey,ExpiresAt from Secrets where Id = 1 ;"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
		return nil, errors.New("There has been problem with reading from db")
	}

	var secret data.Secret
	var dateStr string

	for rows.Next() {
		err := rows.Scan(&secret.Secret, &dateStr)
		expiresAt, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			panic(err.Error())
		}
		secret.ExpiresAt = data.CustomTime{expiresAt}
		if err != nil {
			panic(err.Error())
		}
	}
	return &secret, nil
}

func (s SecretRepoSql) SaveSecret(value data.Secret) error {
	db, err := s.OpenConnection()
	if err != nil {
		return errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()

	query := "Insert into Secrets (SecretKey,ExpiresAt) values (?,?)"
	_, err = db.Exec(query, value.Secret, value.ExpiresAt.Time)
	if err != nil {
		return fmt.Errorf("failed to insert secret key: %v", err)
	}
	return nil
}

func (s SecretRepoSql) UpdateSecret(value data.Secret) error {
	db, err := s.OpenConnection()
	if err != nil {
		return errors.New("There has been problem with connectiong to db")
	}
	defer db.Close()

	query := "update Secrets set SecretKey = ?,ExpiresAt = ? where Id = 1"
	_, err = db.Exec(query, value.Secret, value.ExpiresAt.Time)
	if err != nil {
		return fmt.Errorf("failed to insert secret key: %v", err)
	}
	return nil
}

func (s SecretRepoSql) OpenConnection() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s:%s)/%s", s.pass, s.host, s.port, schema))
}
