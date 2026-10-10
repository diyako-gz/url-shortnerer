package repo

import (
	"context"
	"urlShortnerer/cmd/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type LinkRepository struct {
	mongoCollection *mongo.Collection
}

func NewLinkRepository(m *mongo.Collection) *LinkRepository {
	return &LinkRepository{
		mongoCollection: m,
	}
}

func (r *LinkRepository) Create(c context.Context, link model.Link) error {
	_, err := r.mongoCollection.InsertOne(c, link)
	return err
}
