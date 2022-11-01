package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ### ### ###
// Business logic interfaces

type Article struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id, omitempty"`
	Serial      int                `bson:"serial" json:"serial"`
	Name        string             `bson:"name" json:"name"`
	Image       string             `bson:"image" json:"image"`
	Units       string             `bson:"units" json:"units"`
	Annotations string             `bson:"annotations" json:"annotations"`
	Discount    float64            `bson:"discount" json:"discount"`
	Price       float64            `bson:"price" json:"price"`
	Description string             `bson:"description" json:"description"`
}

type ArticleImage struct {
	ID     primitive.ObjectID `bson:"_id, omitempty"`
	Serial int                `bson:"serial" json:"serial"`
	Image  string             `bson:"image" json"image"`
}

type Cart_Items struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id, omitempty"`
	Quantity int8               `bson:"quantity" "json:"quantity"`
}

type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id, omitempty"`
	Order    []Cart_Items       `bson:"order" json:"order"`
	Discount float64            `bson:"discount" json:"discount"`
	Total    float64            `bson:"total" json:"total"`
}

type Users struct {
	ID        primitive.ObjectID `json:"id, omitempty"`
	Firstname string             `json:"firstname"`
	Lastname  string             `json:"lastname"`
	Password  string             `json:"password"`
	Email     string             `json:"email"`
}

// ### ### ###
// Helpers interfaces

type GenericResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
