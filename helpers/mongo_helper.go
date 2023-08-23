package helpers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func HexMongoID(input *mongo.InsertOneResult) string {
	if oid, ok := input.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex()
	}
	return ""
}
