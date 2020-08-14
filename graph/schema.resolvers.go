package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/pergpau/go-graphql-jishono/graph/generated"
	"github.com/pergpau/go-graphql-jishono/graph/model"
	"github.com/pergpau/go-graphql-jishono/services/definisjonservice"
	"github.com/pergpau/go-graphql-jishono/services/oppslagservice"
)

func (r *queryResolver) AlleOppslag(ctx context.Context) ([]*model.Oppslag, error) {
	var resultOppslag []*model.Oppslag
	var dbOppslag []oppslagservice.Oppslag
	dbOppslag = oppslagservice.GetAll()
	for _, oppslag := range dbOppslag {
		resultOppslag = append(resultOppslag, &model.Oppslag{ID: oppslag.ID,
			Oppslag: oppslag.Oppslag, BoyTabell: oppslag.BoyTabell})
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
