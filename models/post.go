package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Content string             `bson:"content"`
	User    primitive.ObjectID `bson:"user"` // Assuming a reference to a User
}

// Add methods here for operations related to the Post model.
