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

type ProductImageResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Image   string `json:"image"`
}

type LoginResponse struct {
	Error   bool       `json:"error"`
	Message string     `json:"message"`
	User    PublicUser `json:"user"`
}

type FavoritesResponse struct {
	Error     bool     `json:"error"`
	Message   string   `json:"message"`
	Favorites []string `json:"favorites"`
}

type GenericCartProductsResponse struct {
	Error    bool          `json:"error"`
	Message  string        `json:"message"`
	Products []CartVerbose `json:"products"`
}
