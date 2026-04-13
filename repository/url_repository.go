package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-short/config"
	"github.com/go-short/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get the collection
func getCol() string {
	return "urls"
}

func Save(url models.URL) (models.URL, error) {
	col := config.GetCollection(getCol())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existing models.URL
	err := col.FindOne(ctx, bson.M{"url": url.URL}).Decode(&existing)

	if err == nil {
		return existing, nil
	}

	if err != mongo.ErrNoDocuments {
		return models.URL{}, err
	}

	_, err = col.InsertOne(ctx, url)

	return url, err
}

// GET the url model with count
func Get(code string) (models.URL, error) {
	col := config.GetCollection(getCol())

	var url models.URL

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := col.FindOne(ctx, bson.M{"short_code": code}).Decode(&url)
	return url, err
}

// Update the same url with the updated new url
func Update(code string, newURL string) error {
	col := config.GetCollection(getCol())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"url":       newURL,
			"updatedAt": time.Now().Unix(),
		},
	}

	res, err := col.UpdateOne(
		ctx,
		bson.M{"short_code": code},
		update,
	)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("NotFound")
	}

	return nil
}

// Delete the code from db when we need.
func Delete(code string) error {
	col := config.GetCollection(getCol())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := col.DeleteOne(ctx, bson.M{"short_code": code})
	if res.DeletedCount == 0 {
		return fmt.Errorf("NotFound")
	}

	return err
}

// This is use for when we hit then we can increment the click count
func IncrementClicks(code string) error {
	col := config.GetCollection(getCol())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := col.UpdateOne(
		ctx,
		bson.M{"short_code": code},
		bson.M{"$inc": bson.M{"clicks": 1}},
	)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no document found with code: %s", code)
	}

	return nil
}
