package interfaces

type ProductsPage struct{
  Error		bool		`json:"error"`
  Message	string		`json:"message"`
  Products	[]Article	`json:"products"`
}
