package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// Email    string             `bson:"email" binding:"required"`
	Username string `bson:"username" binding:"required"`
	Password string `bson:"password" binding:"required"`
}
