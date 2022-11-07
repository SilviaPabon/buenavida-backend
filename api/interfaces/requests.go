package interfaces

type FilterProductsByText struct {
	Criteria string `json:"search-criteria"`
}

type LoginPayload struct {
	Mail     string `json:"email"`
	Password string `json:"password"`
}

type FilterProducts struct {
	From     float32 `json:"from"`
	To       float32 `json:"to"`
	Criteria string  `json:"search_criteria"`
}
