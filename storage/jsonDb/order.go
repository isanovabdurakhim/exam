package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type orderRepo struct {
	fileName string
}

func NewOrderRepo(fileName string) *orderRepo {
	return &orderRepo{
		fileName: fileName,
	}
}

func (o *orderRepo) Read() ([]models.Order, error) {
	data, err := ioutil.ReadFile(o.fileName)
	if err != nil {
		return []models.Order{}, err
	}

	var orders []models.Order
	err = json.Unmarshal(data, &orders)
	if err != nil {
		return []models.Order{}, err
	}
	return orders, nil
}

func (o *orderRepo) Create(req *models.CreateOrder, total float64) (string, error) {
	orders, err := o.Read()
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	orders = append(orders, models.Order{
		Id:               uuid,
		User_id:          req.User_id,
		Cutomer_name:     req.Cutomer_name,
		Customer_address: req.Customer_address,
		Customer_phone:   req.Customer_phone,
		Total:            total,
	})
	data, err := ioutil.ReadFile("./data/order_items.json")
	if err != nil {
		return "", err
	}

	orderItems := []models.OrderItems{}
	err = json.Unmarshal(data, &orderItems)
	if err != nil {
		return "", err
	}

	for _, v := range req.OrderItems {
		orderItems = append(orderItems, models.OrderItems{
			Product_id: v.Product_id,
			Count:      v.Count,
			Order_id:   uuid,
		})
	}

	bodyOrderItems, err := json.MarshalIndent(orderItems, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile("./data/order_items.json", bodyOrderItems, os.ModePerm)
	if err != nil {
		return "", err
	}

	body, err := json.MarshalIndent(orders, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(o.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (o *orderRepo) GetByOrderID(id string) (models.GetOrderById, error) {
	orders, err := o.Read()
	if err != nil {
		return models.GetOrderById{}, err
	}
	var order models.GetOrderById

	flag := false
	for _, v := range orders {
		if v.Id == id {
			order.User_id = v.User_id
			order.Customer_address = v.Customer_address
			order.Customer_address = v.Customer_address
			order.Customer_phone = v.Customer_phone
			order.Total = v.Total

			flag = true
		}
	}
	if !flag {
		return models.GetOrderById{}, errors.New("order not found")
	}

	data, err := ioutil.ReadFile("./data/order_items.json")
	if err != nil {
		return models.GetOrderById{}, err
	}

	orderItems := []models.OrderItems{}
	err = json.Unmarshal(data, &orderItems)
	if err != nil {
		return models.GetOrderById{}, err
	}

	for _, v := range orderItems {
		if v.Order_id==id{
			orderItems = append(orderItems, models.OrderItems{
				Product_id: v.Product_id,
				Count:      v.Count,
				Order_id:   v.Order_id,
			})
		}
	}

	order.OrderItems = orderItems

	return order, nil
}

func (o *orderRepo) Update(req *models.UpdateOrder, userId string) error {
	orders, err := o.Read()
	if err != nil {
		return err
	}

	flag := true
	for i, v := range orders {
		if v.Id == userId {
			orders[i].Cutomer_name = req.Cutomer_name
			orders[i].Customer_address = req.Customer_address
			orders[i].Customer_phone = req.Customer_phone
			flag = false
		}
	}

	if flag {
		return errors.New("There is no order with this id")
	}

	body, err := json.MarshalIndent(orders, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(o.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (o *orderRepo) Delete(req *models.OrderPrimaryKey) error {
	orders, err := o.Read()
	if err != nil {
		return err
	}
	flag := true
	for i, v := range orders {
		if v.Id == req.Id {
			orders = append(orders[:i], orders[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		return errors.New("There is no product with this id")
	}

	body, err := json.MarshalIndent(orders, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(o.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}


