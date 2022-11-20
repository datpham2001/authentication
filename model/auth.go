package model

import (
	"authentication/model/enum"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID              *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime     *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	LastUpdatedTime *time.Time          `json:"lastUpdatedTime,omitempty" bson:"last_updated_time,omitempty"`

	AccountID string                    `json:"accountId,omitempty" bson:"account_id,omitempty"`
	Username  string                    `json:"username,omitempty" bson:"username,omitempty"`
	Email     string                    `json:"email,omitempty" bson:"email,omitempty"`
	Password  string                    `json:"password,omitempty" bson:"password,omitempty"`
	Status    enum.AccountStatusValue   `json:"status,omitempty" bson:"status,omitempty"`
	Role      enum.AccountRoleEnumValue `json:"role,omitempty" bson:"role,omitempty"`
}
