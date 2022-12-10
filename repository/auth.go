package repository

import (
	"authentication/dto"
	"authentication/model"
	"authentication/model/enum"
	"authentication/pkg"
	"authentication/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	AUTHENTICATION_COLLECTION_NAME = "authentication"
)

func NewAuthRepository(db *mongo.Database) AuthRepository {
	r := &Instance{
		db:             db,
		coll:           db.Collection(AUTHENTICATION_COLLECTION_NAME),
		DBName:         db.Name(),
		TemplateObject: &model.Account{},
	}

	return r
}

func (r *Instance) SignUp(input *dto.AccountRegisterRequest) (*model.Account, dto.DatabaseError) {
	account := &model.Account{
		Email:    input.Account.Email,
		Username: input.Account.Username,
		Password: pkg.HashPassword(input.Account.Password),
	}

	// check account is existed in database
	_, err := r.QueryOne(model.Account{
		ComplexQuery: []*bson.M{
			{
				"$or": []*bson.M{
					{
						"username": account.Username,
					},
					{
						"email": account.Email,
					},
				},
			},
		},
	})

	if err == nil {
		return nil, dto.DatabaseError{Code: enum.ErrorCodeExisted.Account}
	}

	account.AccountID = utils.GenID(utils.ACCOUNT)
	account.Status = enum.AccountStatus.Active

	res, err := r.Create(account)
	if err != nil {
		return nil, dto.DatabaseError{
			Code: enum.ErrorDatabaseCode.Create,
		}
	}

	return res.([]*model.Account)[0], dto.DatabaseError{Code: enum.ErrorCodeValue("SUCCESS")}
}

func (r *Instance) Login(input *dto.AccountLoginRequest) (*model.Account, dto.DatabaseError) {

}

func (r *Instance) Logout(input *dto.AccountLoginRequest) (*model.Account, dto.DatabaseError) {

}
