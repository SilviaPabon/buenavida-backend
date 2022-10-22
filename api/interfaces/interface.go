package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID          primitive.ObjectID `bson:"_id,omitempty", json:"id, omitempty"`
	Name        string             `bson:"name,omitempty", json : "name"`
	Image       string             `bson:"image,omitempty", json : "image"`
	Units       string             `bson:"units,omitempty", json : "units"`
	Annotations string             `bson:"annotations,omitempty", json : "annotations"`
	Discount    float32            `bson:"discount,omitempty", json : "discount"`
	Price       float32            `bson:"price,omitempty", json : "price"`
	Description string             `bson:"description,omitempty", json : "description"`
}

type Cart_Items struct {
	ID       primitive.ObjectID `bson:"_id,omitempty", json:"id, omitempty"`
	Quantity float32            `bson:"quantity,omitempty", json : "quantity"`
}

type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty", json:"id, omitempty"`
	Order    []Cart_Items       `bson:"order,omitempty", json : "order"`
	Discount float32            `bson:"discount,omitempty", json : "discount"`
	Total    float32            `bson:"total,omitempty", json : "total"`
}

type Users struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty", json:"id, omitempty"`
	user_name      string               `bson:"user_name,omitempty", json : "user_name"`
	user_last_name string               `bson:"user_last_name,omitempty", json : "user_last_name"`
	password       string               `bson:"password,omitempty", json : "password"`
	Email          string               `bson:"email,omitempty", json : "email"`
	Favorites      []primitive.ObjectID `bson:"favorites,omitempty", json : "favorites"`
	Cart           []Cart_Items         `bson:"cart,omitempty", json : "cart"`
	Purchases      []Order              `bson:"ordersBuy,omitempty", json : "ordersBuy"`
}
