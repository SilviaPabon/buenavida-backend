package interfaces

type FilterProductsByText struct{
  Criteria	string		`json:"search-criteria"`
}

type LoginPayload struct{
  Mail		string 		`json:"email"`
  Password	string		`json:"password"`
}
