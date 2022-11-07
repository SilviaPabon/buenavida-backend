package models

import (
	//"fmt"
	"context"
	"time"

	"github.com/SilviaPabon/buenavida-backend/configs"
	"github.com/SilviaPabon/buenavida-backend/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongodb collection
var productsCollection = configs.GetCollection("products")
var imagesCollection = configs.GetCollection("images")

// GetAllProducts get entire products collection
func GetAllProducts() (p []interfaces.Article, e error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Make query
	var products []interfaces.Article
	cursor, err := productsCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return products, nil
	}

	// Parse to struct annd return
	if err = cursor.All(ctx, &products); err != nil {
		return products, err
	}

	return products, nil
}

// GetProductsByPage get products by given page
func GetProductsByPage(page int) (p []interfaces.Article, e error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Prepare query
	options := options.Find().SetSkip(int64(page-1) * 12).SetLimit(12)

	// Make query
	var products []interfaces.Article
	cursor, err := productsCollection.Find(ctx, bson.D{{}}, options)

	if err != nil {
		return products, err
	}

	// Parse from bson to struct
	if err = cursor.All(ctx, &products); err != nil {
		return products, err
	}

	return products, nil
}

func GetProductsFiltrated(criteria string, min float32, max float32) (p []interfaces.Article, e error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{
		"$or", bson.A{
			bson.D{{"name", primitive.Regex{
				Pattern: criteria,
				Options: "gi",
			}}},
			bson.D{{"description", primitive.Regex{
				Pattern: criteria,
				Options: "gi",
			}}},
		}},
		{"$and",
			bson.A{
				bson.D{{"price", bson.D{{"$gte", min}}}},
				bson.D{{"price", bson.D{{"$lte", max}}}},
			},
		},
	}

	cursor, err := productsCollection.Find(ctx, filter)

	var results []interfaces.Article
	if err != nil {
		return results, nil
	}

	// Parse to struct annd return
	if err = cursor.All(ctx, &results); err != nil {
		return results, err
	}

	return results, nil
}

// SearchByText filter product on database by provided text
func SearchByText(criteria string) (p []interfaces.Article, e error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Prepare query

	filter := bson.D{{
		"$or", bson.A{
			bson.D{{"name", primitive.Regex{
				Pattern: criteria,
				Options: "gi",
			}}},
			bson.D{{"description", primitive.Regex{
				Pattern: criteria,
				Options: "gi",
			}}},
		},
	}}

	var products []interfaces.Article

	// Make query
	cursor, err := productsCollection.Find(ctx, filter)

	if err != nil {
		return products, err
	}

	// Parse from bson to struct
	if err = cursor.All(ctx, &products); err != nil {
		return products, err
	}

	return products, nil
}

func GetDetailsFromID(id string) (p interfaces.Article, e error) {
	//Search in database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return interfaces.Article{}, err
	}

	var product interfaces.Article
	err = productsCollection.FindOne(ctx, bson.D{{"_id", mongoId}}).Decode(&product)

	if err != nil {
		return interfaces.Article{}, err
	}

	return product, err
}

// GetProductImageFromSerial Obtain product image
func GetProductImageFromSerial(serial int) (res interfaces.ArticleImage, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var image interfaces.ArticleImage
	err = imagesCollection.FindOne(ctx, bson.D{{"serial", serial}}).Decode(&image)

	if err != nil {
		return image, err
	}

	return image, nil
}
