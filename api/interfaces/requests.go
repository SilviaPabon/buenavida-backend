package interfaces

import(
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type FilterProductsByText struct{
  Criteria	string		`json:"search-criteria"`
}

type LoginPayload struct{
  Mail		string 		`json:"email"`
  Password	string		`json:"password"`
}

type AddToCartPayload struct{
  Id		primitive.ObjectID	`bson:"_id, omitempty" json:"id"`			
}

type UpdateCartPayload struct{
  Id		primitive.ObjectID	`bson:"_id, omitempty" json:"id"`
  Amount	int			`json:"amount"`
}
