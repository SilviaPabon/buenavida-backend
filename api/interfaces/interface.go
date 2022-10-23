package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ### ### ###
// Business logic interfaces

type Article struct {
	ID          primitive.ObjectID `bson:"_id,omitempty", json:"id, omitempty"`
	Name        string             `bson:"name", json:"name"`
	Image       string             `bson:"image", json:"image"`
	Units       string             `bson:"units", json:"units"`
	Annotations string             `bson:"annotations", json:"annotations"`
	Discount    float64            `bson:"discount", json:"discount"`
	Price       float64            `bson:"price", json:"price"`
	Description string             `bson:"description", json:"description"`
}

type Cart_Items struct {
	ID       primitive.ObjectID `bson:"_id,omitempty", json:"id, omitempty"`
	Quantity int8            `bson:"quantity,omitempty", json : "quantity"`
}

type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty", json:"id, omitempty"`
	Order    []Cart_Items       `bson:"order,omitempty", json : "order"`
	Discount float64            `bson:"discount,omitempty", json : "discount"`
	Total    float64            `bson:"total,omitempty", json : "total"`
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

// ### ### ###
// Helpers interfaces

type GenericResponse struct{
  Error 	bool		`json:"error"`
  Message	string		`json:"message"`
}
