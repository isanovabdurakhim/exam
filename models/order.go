package models

type OrderPrimaryKey struct {
	Id string `json:"id"`
}

type Order struct{
  Id string `json:"id"`
  User_id string `json:"user_id"`
  Cutomer_name string `json:"customer_name"`
  Customer_address string `json:"customer_address"`
  Customer_phone string `json:"customer_phone"`
  Total float64 `json:"total"`
}

type CreateOrder struct{
  User_id string `json:"user_id"`
  Cutomer_name string `json:"cutomer_name"`
  Customer_address string `json:"customer_address"`
  Customer_phone string `json:"customer_phone"`
  OrderItems []CreateOrderItems `json:"orderitems"`
  
}
type CreateOrderItems struct{
  Product_id string `json:"product_id"`
  Count int `json:"count"`
}

type OrderItems struct{
  Product_id string `json:"product_id"`
  Count int `json:"count"`
  Order_id string `json:"order_id"`
}

type GetOrderById struct{
	Id string `json:"id"`
	User_id string `json:"user_id"`
	Cutomer_name string `json:"customer_name"`
	Customer_address string `json:"customer_address"`
	Customer_phone string `json:"customer_phone"`
	Total float64 `json:"total"`
	OrderItems []OrderItems	
}

type UpdateOrder struct {
	Cutomer_name string `json:"customer_name"`
	Customer_address string `json:"customer_address"`
	Customer_phone string `json:"customer_phone"`
}