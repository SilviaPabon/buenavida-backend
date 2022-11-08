package interfaces

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

type CartItems struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id, omitempty"`
	Iduser   int8   `bson:"iduser" json:"iduser"`
	Quantity int8   `bson:"quantity" "json:"quantity"`
}

type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id, omitempty"`
	Order    []CartItems       `bson:"order" json:"order"`
	Discount float64            `bson:"discount" json:"discount"`
	Total    float64            `bson:"total" json:"total"`
}

type User struct {
	ID        int    `json:"id, omitempty"`
	Firstname string `json:"firstname" validate:"required,min=1,max=125"`
	Lastname  string `json:"lastname" validate:"required,min=1,max=125"`
	Password  string `json:"password" validate:"required,min=8,max=250,containsany=!@#?*,containsany=1234567890,containsany=ABCDEFGHIJKLMNÑOPQRSTUVWXYZ"`
	Email     string `json:"email" validate:"required,max=250,email"`
}

// ### ### ###
// Helpers interfaces
type PublicUser struct {
	ID        int    `json:"id, omitempty"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type GenericResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type JWTCustomClaims struct {
	jwt.RegisteredClaims           // Default claims
	ID                   int       `json:"id"`
	Email                string    `json:"email"`
	UUID                 uuid.UUID `json:"uuid"`
}

type Favorite struct {
	FavoriteId string
}

type CartVerbose struct {
	Name     string  `json:"name"`
	Units    string  `json:"units"`
	Quantity int8    `json:"quantity"`
	Price    float64 `json:"price"`
	Image    string  `json:"image"`
}

type OrderProduct struct {
  Product	string		`json:"product"`
  Amount	int		`json:"amount"`
}

type OrderResume struct {
  Order		int		`json:"order"`
  Products	[]OrderProduct	`json:"products"`
}