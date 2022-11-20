package dao

import (
	"authentication/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	AUTHENTICATION_COLLECTION_NAME = "authentication"
)

func NewAuthDAO(db *mongo.Database) AuthDAO {
	m := &Instance{
		db:             db,
		coll:           db.Collection(AUTHENTICATION_COLLECTION_NAME),
		DBName:         db.Name(),
		TemplateObject: &model.Account{},
	}

	return m
}

func (r *Instance) Register(input *model.Account) (*model.Account, error) {
	res, err := r.Create(input)
	if err != nil {
		return nil, err
	}

	return res.([]*model.Account)[0], nil
}

func (r *Instance) Login(input *model.Account) (*model.Account, error) {

}
