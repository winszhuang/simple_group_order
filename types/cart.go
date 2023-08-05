package types

type CartItemBaseData struct {
	CartItemID int
	Quantity   int
	ProductID  int
}

type CartData struct {
	UserID int
	CartID int
	Items  []CartItemBaseData
}
