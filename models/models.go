package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// URLMapping represents the schema for a URL mapping stored in MongoDB.
type URLMapping struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	LongURL   string             `bson:"longUrl" json:"longUrl"`
	ShortURL  string             `bson:"shortUrl" json:"shortUrl"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}
