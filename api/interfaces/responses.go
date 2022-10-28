package interfaces

type GenericProductsArrayResponse struct {
	Error    bool      `json:"error"`
	Message  string    `json:"message"`
	Products []Article `json:"products"`
}

type GenericProductResponse struct {
	Error   bool    `json:"error"`
	Message string  `json:"message"`
	Product Article `json:"product"`
}

type ProductImageResponse struct{
  Error 	bool		`json:"error"`
  Message	string 		`json:"message"`
  Image		string		`json:"image"`
}
