package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/pergpau/go-graphql-jishono/graph/generated"
	"github.com/pergpau/go-graphql-jishono/graph/model"
	"github.com/pergpau/go-graphql-jishono/services/definisjonservice"
	"github.com/pergpau/go-graphql-jishono/services/oppslagservice"
)

func (r *queryResolver) Oppslag(ctx context.Context) ([]*model.Oppslag, error) {
	var resultOppslag []*model.Oppslag
	var dbOppslag []oppslagservice.Oppslag
	dbOppslag = oppslagservice.GetAll()
	for _, oppslag := range dbOppslag {
		var resultDefinisjoner []*model.Definisjon
		var definisjoner []definisjonservice.Definisjon
		definisjoner = definisjonservice.GetDefinisjonerByLemmaID(oppslag.ID)
		for _, definisjon := range definisjoner {
			resultDefinisjoner = append(resultDefinisjoner, &model.Definisjon{ID: definisjon.ID,
				Prioritet: definisjon.Prioritet, Definisjon: definisjon.Definisjon})
		}
		resultOppslag = append(resultOppslag, &model.Oppslag{ID: oppslag.ID,
			Oppslag: oppslag.Oppslag, BoyTabell: oppslag.BoyTabell, Definisjoner: resultDefinisjoner})
	}
	return resultOppslag, nil
}

func (r *queryResolver) EnkeltOppslag(ctx context.Context, id string) (*model.Oppslag, error) {
	var dbOppslag oppslagservice.Oppslag
	dbOppslag = oppslagservice.GetSingleByID(id)
	var resultDefinisjoner []*model.Definisjon
	var definisjoner []definisjonservice.Definisjon
	definisjoner = definisjonservice.GetDefinisjonerByLemmaID(id)
	for _, definisjon := range definisjoner {
		resultDefinisjoner = append(resultDefinisjoner, &model.Definisjon{ID: definisjon.ID,
			Prioritet: definisjon.Prioritet, Definisjon: definisjon.Definisjon, Opprettet: definisjon.Opprettet})
	}
	result := &model.Oppslag{ID: dbOppslag.ID,
		Oppslag: dbOppslag.Oppslag, BoyTabell: dbOppslag.BoyTabell, Definisjoner: resultDefinisjoner}
	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) DefinisjonerFraOppslag(ctx context.Context, id string) ([]*model.Definisjon, error) {
	panic(fmt.Errorf("not implemented"))
}
