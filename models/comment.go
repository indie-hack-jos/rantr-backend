package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Content string             `bson:"content"`
	Author  primitive.ObjectID `bson:"author"` // Assuming a reference to a User
	Post    primitive.ObjectID `bson:"post"`   // Assuming a reference to a Post
}

// Add methods here for operations related to the Comment model.
