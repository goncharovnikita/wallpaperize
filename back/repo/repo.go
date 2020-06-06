package repo

import "go.mongodb.org/mongo-driver/mongo"

// Repo type
type Repo struct {
	imagesCol *mongo.Collection
}

// New creates new Repo
func New(db *mongo.Database) *Repo {
	return &Repo{
		imagesCol: db.Collection("images"),
	}
}
