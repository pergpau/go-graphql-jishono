package oppslagservice

import (
	database "github.com/pergpau/go-graphql-jishono/db"
	"github.com/pergpau/go-graphql-jishono/services/definisjonservice"
	"log"
)

type Oppslag struct {
	ID string 
	Oppslag string 
	BoyTabell *string
	Definisjoner *definisjonservice.Definisjon
  }

func GetAll() []Oppslag {
	stmt, err := database.Db.Prepare("SELECT lemma_id, oppslag, boy_tabell FROM oppslag")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var oppslagListe []Oppslag
	for rows.Next() {
		var oppslag Oppslag
		err := rows.Scan(&oppslag.ID, &oppslag.Oppslag, &oppslag.BoyTabell)
		if err != nil{
			log.Fatal(err)
		}
		oppslagListe = append(oppslagListe, oppslag)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return oppslagListe

}

func GetSingleByID(id string) Oppslag {

	var oppslag Oppslag
	err := database.Db.QueryRow("SELECT lemma_id, oppslag, boy_tabell FROM oppslag WHERE lemma_id = ?", id).Scan(&oppslag.ID, &oppslag.Oppslag, &oppslag.BoyTabell)
	if err != nil {
		log.Fatal(err)
	}
	
	return oppslag

}
