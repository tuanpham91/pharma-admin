package recordType 

type Record struct {
	Id 			   int
	Name           string
	Quantity       int
	Price          float32
	ExpirationDate string
	DateOfRecord   string
}

type ItemInventory struct {
	Id 				int
	Name 			string 
	Quantity 		int 
	ExpirationDate 	string
}