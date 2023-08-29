package models

import (
	"api_book/helpers"
)

type User struct {
	// ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	*helpers.SQLModel `bson:",inline" json:",inline"`
	Name              string `bson:"name" json:"name"`
	Email             string `bson:"email" json:"email"`
	Password          string `bson:"password" json:"password"`
	Phone             string `bson:"phone" json:"phone"`
	Address           string `bson:"address" json:"address"`
	Avatar            string `bson:"avatar" json:"avatar"`
	Role              string `bson:"role" json:"role"`
	AccessToken       string `bson:"access_token" json:"access_token"`
}
