package controller

import "app/models"

func (c *Controller) CreateOrder(req *models.CreateOrder) (string, error) {

	var total float64

	for _, v := range req.OrderItems{
		
		product,err:=c.store.Product().GetByID(&models.ProductPrimaryKey{Id: v.Product_id})
		if err!= nil{
			return "", err
		}
		 total+= product.Price*float64(v.Count)
	}

	id, err := c.store.Order().Create(req, total)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) GetByOrderID(id string) (models.GetOrderById, error) {
	order, err := c.store.Order().GetByOrderID(id)
	if err != nil {
		return models.GetOrderById{}, err
	}
	return order, nil
}

func (c *Controller) UpdateOrder(req *models.UpdateOrder, order_id string) error {
	err := c.store.Order().Update(req, order_id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) DeleteOrder(req *models.OrderPrimaryKey) error {
	err := c.store.Order().Delete(req)
	if err != nil {
		return err
	}
	return nil
}