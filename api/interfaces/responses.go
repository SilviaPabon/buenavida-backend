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

type GenericFavoriteResponse struct {
	Error    bool    `json:"error"`
	Message  string  `json:"message"`
	Favorite Article `json:"product"`
}

type FavoritesResponse struct {
	Error     bool     `json:"error"`
	Message   string   `json:"message"`
	Favorites []string `json:"favorites"`
}
