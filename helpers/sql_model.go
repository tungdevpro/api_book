package helpers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SQLModel struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt *time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time         `bson:"updated_at" json:"updated_at"`
}
