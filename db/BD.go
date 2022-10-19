package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Article struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"Name,omitempty"`
	Units           int32              `bson:"Units,omitempty"`
	Brand           string             `bson:"Brand,omitempty"`
	Unit_Value      int32              `bson:"Unit_Value,omitempty"`
	DiscountCurrent int32              `bson:"DiscountCurrent,omitempty"`
	Price           int32              `bson:"Price,omitempty"`
	Description     string             `bson:"Description,omitempty"`
}

type Cart_Items struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Quantity int32              `bson:"Quantity,omitempty"`
}

type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Order    []string           `bson:"Order,omitempty"`
	Discount int32              `bson:"Discount,omitempty"`
	Total    int32              `bson:"Total,omitempty"`
}

type Users struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	user_name      string             `bson:"user_name,omitempty"`
	user_last_name string             `bson:"user_last_name,omitempty"`
	password       string             `bson:"password,omitempty"`
	Email          string             `bson:"Email,omitempty"`
	Favorites      []string           `bson:"Favorites,omitempty"`
	Cart           []string           `bson:"Cart,omitempty"`
	OrdersBuy      []string           `bson:"OrdersBuy,omitempty"`
}

func main() {
	opt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(opt)
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)

	ParcialWEBDB := client.Database("ParcialWEB")
	ArticleCollection := ParcialWEBDB.Collection("Article")
	Cart_ItemsCollection := ParcialWEBDB.Collection("Cart_Items")
	Order := ParcialWEBDB.Collection("Order")
	Users := ParcialWEBDB.Collection("Users")

	defer ArticleCollection.Drop(ctx)
	defer Cart_ItemsCollection.Drop(ctx)
	defer Order.Drop(ctx)
	defer Users.Drop(ctx)

}
