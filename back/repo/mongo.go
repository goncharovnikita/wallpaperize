package repo

import "go.mongodb.org/mongo-driver/mongo"

// Mongo type
type Mongo struct {
	imagesCol *mongo.Collection
}

// NewMongo creates new Repo
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		imagesCol: db.Collection("images"),
	}
}
