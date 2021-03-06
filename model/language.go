package model

import (
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

// Language 言語情報
type Language struct {
	ID   int    `db:"language_id" json:"id"`
	Name string `db:"language_name" json:"name"`
}

// GetLanguages 言語情報を取得します
func GetLanguages(tx *gorp.Transaction) ([]Language, error) {

	languages, err := selectToLanguages(tx)
	if err != nil {
		return languages, err
	}

	return languages, nil
}

// selectToLanguages 言語情報を検索します
func selectToLanguages(tx *gorp.Transaction) ([]Language, error) {
	var languages []Language
	_, err := tx.Select(&languages, `
		select
		  id language_id,
		  language_name
		from
		  languages
		order by
		  id
		`)
	if err != nil {
		return languages, err
	}

	return languages, nil
}
