package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type URL struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShortCode string             `bson:"short_code" json:"code"`
	URL       string             `bson:"url" json:"original_url"`
	Clicks    int                `bson:"clicks" json:"clicks"`
	CreatedAt int64              `bson:"createdAt" json:"created_at"`
	UpdatedAt int64              `bason:"updatedAt" json:"updated_at"`
}
