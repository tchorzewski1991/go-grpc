package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Blog represents our domain model. It should contain only critical
// business logic. According to clean architecture it refers to the most
// inner layer. It is the abstraction build independently on top of the
// any kind of database.

type Blog struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}
