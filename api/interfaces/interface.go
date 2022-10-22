package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Image       string             `bson:"image,omitempty"`
	Units       string             `bson:"units,omitempty"`
	Annotations string             `bson:"annotations,omitempty"`
	Discount    float32            `bson:"discountCurrent,omitempty"`
	Price       float32            `bson:"price,omitempty"`
	Description string             `bson:"description,omitempty"`
}

type Cart_Items struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Quantity float32            `bson:"quantity,omitempty"`
}

type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Order    []Cart_Items       `bson:"order,omitempty"`
	Discount float32            `bson:"discount,omitempty"`
	Total    float32            `bson:"total,omitempty"`
}

type Users struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	user_name      string               `bson:"user_name,omitempty"`
	user_last_name string               `bson:"user_last_name,omitempty"`
	password       string               `bson:"password,omitempty"`
	Email          string               `bson:"email,omitempty"`
	Favorites      []primitive.ObjectID `bson:"favorites,omitempty"`
	Cart           []Cart_Items         `bson:"cart,omitempty"`
	Purchases      []Order              `bson:"ordersBuy,omitempty"`
}
