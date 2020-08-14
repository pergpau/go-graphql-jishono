package definisjonservice

import (
	database "github.com/pergpau/go-graphql-jishono/db"
	"log"
	"fmt"
)

type Definisjon struct {
	ID string 
	Definisjon string
	Prioritet int
	Opprettet string
  }


func GetDefinisjonerByLemmaID(id string) []Definisjon{

	var definisjoner []Definisjon
	rows, err := database.Db.Query("SELECT def_id, prioritet, definisjon, opprettet FROM definisjon WHERE lemma_id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var definisjon Definisjon
		err := rows.Scan(&definisjon.ID, &definisjon.Prioritet, &definisjon.Definisjon, &definisjon.Opprettet)
		if err != nil{
			log.Fatal(err)
		}
		definisjoner = append(definisjoner, definisjon)
	}
	fmt.Println(definisjoner)
	return definisjoner
}

