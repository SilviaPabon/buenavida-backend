package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
