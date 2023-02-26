package controller

import (
	"app/models"
	"fmt"
	"sort"
	"time"
)

func (c *Controller) Filter(from_date, to_date string) ([]models.ShopCart, error) {
	data, err := c.store.ShopCart().Filter(from_date, to_date)
	if err != nil {
		return []models.ShopCart{}, err
	}
	return data, nil
}

func (c *Controller) UserHistory(req *models.UserPrimaryKey) (string, []models.History) {

	history := []models.History{}
	shopcart, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return "", []models.History{}
	}

	user, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: req.Id})
	if err != nil {
		return "", []models.History{}
	}

	name := user.Name

	for _, shopcart := range shopcart {
		if shopcart.Status == true && shopcart.UserId == req.Id {
			// fmt.Println("ok")

			product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{shopcart.ProductId})
			if err != nil {
				return "", []models.History{}
			}

			history = append(history, models.History{
				Name:  product.Name,
				Price: product.Price,
				Count: shopcart.Count,
				Total: int(product.Price) * shopcart.Count,
				Time:  shopcart.Time,
			})

		}

	}

	return name, history

}

func (c *Controller) Spend(req *models.UserPrimaryKey) (string, int) {
	name, val := c.UserHistory(req)

	spend := 0

	for _, v := range val {
		spend += v.Total
	}
	return name, spend
}

func (c *Controller) SoldProducts() (map[string]int, error) {
	soldproducts := map[string]int{}
	shopcart, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return soldproducts, err
	}

	for _, val := range shopcart {
		if val.Status == true {
			product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{val.ProductId})
			if err != nil {
				return soldproducts, err
			}
			soldproducts[product.Name] += val.Count
		}
	}

	return soldproducts, nil
}

func (c *Controller) Top10Products() (map[string]int, error) {

	result := map[string]int{}

	top10, err := c.SoldProducts()
	if err != nil {
		return result, err
	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range top10 {
		ss = append(ss, kv{k, v})
	}

	// Then sorting the slice by value, higher first.
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for i := 0; i < 10; i++ {

		fmt.Println(ss[i])

	}
	return result, nil
}

func (c *Controller) Last10Products() (map[string]int, error) {

	result := map[string]int{}

	top10, err := c.SoldProducts()
	if err != nil {
		return result, err
	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range top10 {
		ss = append(ss, kv{k, v})
	}

	// Then sorting the slice by value, higher first.
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	for i := 0; i < 10; i++ {

		fmt.Println(ss[i])

	}
	return result, nil
}

func (c *Controller) DayTopSold() error {

	top := map[string]int{}

	shopcart, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return err
	}

	for _, element := range shopcart {

		layout := "2006-01-02 15:04:05"
		date, err := time.Parse(layout, element.Time)
		if err != nil {
			fmt.Println(err)
			return err
		}
		formatted := date.Format("2006-10-12")

		top[formatted] = top[formatted] + element.Count
	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range top {
		ss = append(ss, kv{k, v})
	}

	// Then sorting the slice by value, higher first.
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for idx, val := range ss {
		fmt.Println(idx+1, val)
	}

	return nil
}

func (c *Controller) SoldCategory() (map[string]int, error) {

	soldcategory := map[string]int{}

	shopcart, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return soldcategory, err
	}

	for _, shopcart := range shopcart {
		if shopcart.Status == true {

			product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{shopcart.ProductId})
			if err != nil {
				return soldcategory, err
			}

			category, err := c.store.Category().GetAll(&models.GetListCategoryRequest{0, 10})
			if err != nil {
				return soldcategory, err
			}

			for _, val := range category.Categories {
				if product.CategoryID == val.Id {
					soldcategory[val.Name] += shopcart.Count
				}
			}
		}
	}
	return soldcategory, nil
}

func (c *Controller) ActiveClient() error {

	soldcategory := map[string]int{}

	shopCarts, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return err
	}

	for _, val := range shopCarts {
		if val.Status == true {
			soldcategory[val.UserId] = soldcategory[val.UserId] + val.Count
		}
	}

	max := 0
	userId := ""

	for key, val := range soldcategory {
		if max < val {
			max = val
			userId = key
		}
	}

	user, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: userId})
	if err != nil {
		return err
	}

	fmt.Println(user.Name, max)

	return nil
}

func (c *Controller) Sale(req *models.UserPrimaryKey) (string, []models.History) {
	sale := []models.History{}
	shopCarts, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return "", []models.History{}
	}

	user, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: req.Id})
	if err != nil {
		return "hty", []models.History{}
	}

	name := user.Name

	for _, val := range shopCarts {
		if val.Status == true && val.Count > 9 && val.UserId == req.Id{
			product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{val.ProductId})
			if err != nil {
				return "", []models.History{}
			}
			sale = append(sale, models.History{
				Name:  product.Name,
				Price: product.Price,
				Count: val.Count,
				Total: int(product.Price) * val.Count - int(product.Price),
				Time:  val.Time,
			})
		}
	}
	return name, sale
}
