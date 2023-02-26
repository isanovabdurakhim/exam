package main

import (
	"app/config"
	"app/controller"
	"app/storage/jsonDb"
	"fmt"
	"log"
)

func main() {
	cfg := config.Load()

	jsonDb, err := jsonDb.NewFileJson(&cfg)
	if err != nil {
		log.Fatal("error while connecting to database")
	}
	defer jsonDb.CloseDb()

	c := controller.NewController(&cfg, jsonDb)

	// Product(c)
	// Order(c)
	shopCart(c)
	// branch(c)

	// res, err := c.History(&models.UserPrimaryKey{Id: "e6ded598-675b-4de2-a1e9-00a876b8e719"})

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(res)

}

func shopCart(c *controller.Controller) {

	// res, err := c.Filter("2022-03-08", "2022-03-09")

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(res)

	// #######################################################  History  ######################################################################################

	// name, history := c.UserHistory(&models.UserPrimaryKey{Id: "e6ded598-675b-4de2-a1e9-00a876b8e719"})

	// fmt.Println(name,history)

	// ######################################################  Spend Money  ##################################################################################

	// name, history := c.Spend(&models.UserPrimaryKey{Id: "e6ded598-675b-4de2-a1e9-00a876b8e719"})

	// fmt.Println(name,history)

	// ###################################################### Sold Products  #################################################################################

	// name, count := c.SoldProducts()
	// fmt.Println(name, count)

	// ########################################## Top 10 Products &&  Last 10 Products ########################################################################

	// name, count := c.Top10Products()
	// fmt.Println(name, count)

	// name, count := c.Last10Products()
	// fmt.Println(name, count)

	// count := c.DayTopSold()
	// fmt.Println(count)

	// ###################################################### Day Top Sold Products  ########################################################################

	// name, history := c.DayTopSold("514faab1-009e-4352-868e-ef5041fc2adf", "2022-03-08")

	// fmt.Println(name,history)

	// #########################################################  Sold Category   ##########################################################################

	// name, history := c.SoldCategory()

	// fmt.Println(name,history)

	// #########################################################  Active Client   ##########################################################################

	history := c.ActiveClient()

	fmt.Println(history)

	// name, history := c.Sale(&models.UserPrimaryKey{Id: "c463393b-4690-4dfe-b5e0-7f4a7fa1b21e"})

	// fmt.Println(name, history)
}

// func branch(c *controller.Controller) {
// 	c.CreateBranch(&models.CreateBranch{
// 		Name: "Abdurakhim",
// 	})
// }

// func Product(c *controller.Controller) {

// 	// c.CreateProduct(&models.CreateProduct{
// 	// 	Name:       "Smartfon vivo V25 8/256 GB",
// 	// 	Price:      4_860_000,
// 	// 	CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
// 	// })

// 	// product, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: "38292285-4c27-497b-bc5f-dfe418a9f959"})

// 	// if err != nil {
// 	// 	log.Println(err)
// 	// 	return
// 	// }

// 	products, err := c.GetAllProduct(
// 		&models.GetListProductRequest{
// 			Offset:     0,
// 			Limit:      1,
// 			CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
// 		},
// 	)

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	for in, product := range products.Products {
// 		fmt.Println(in+1, product)
// 	}
// }

// func Category(c *controller.Controller) {
// 	// c.CreateCategory(&models.CreateCategory{
// 	// 	Name:     "Smartfonlar va telefonlar",
// 	// 	ParentID: "eed2e676-1f17-429f-b75c-899eda296e65",
// 	// })

// 	// category, err := c.GetByIdCategory(&models.CategoryPrimaryKey{Id: "eed2e676-1f17-429f-b75c-899eda296e65"})
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// 	return
// 	// }

// 	fmt.Println(category)

// }

// func User(c *controller.Controller) {

// 	sender := "bbda487b-1c0f-4c93-b17f-47b8570adfa6"
// 	receiver := "657a41b6-1bdc-47cc-bdad-1f85eb8fb98c"
// 	err := c.MoneyTransfer(sender, receiver, 500_000)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func Order(c *controller.Controller) {
// 	// id, err := c.CreateOrder(&models.CreateOrder{
// 	// 	User_id:          "e4421e6d-cf37-4dd7-a87f-97c91feffaef",
// 	// 	Cutomer_name:     "Abdurakhim",
// 	// 	Customer_address: "Samarqand",
// 	// 	Customer_phone:   "+998900947778",
// 	// 	OrderItems: []models.CreateOrderItems{
// 	// 		{
// 	// 			Product_id: "31216468-60bd-4694-b5a8-6da80febfdf6",
// 	// 			Count:      3,
// 	// 		},
// 	// 		{
// 	// 			Product_id: "31216468-60bd-4694-b5a8-6da80febfdf6",
// 	// 			Count:      4,
// 	// 		},
// 	// 	},
// 	// })
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// fmt.Println(id)

// 	// order, err := c.GetByOrderID("21ba04ff-a54c-4aad-bbae-66437806d84a")
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// 	return
// 	// }

// 	// fmt.Println(order)

// 	err := c.UpdateOrder(&models.UpdateOrder{
// 		Cutomer_name:     "Davlat",
// 		Customer_address: "Katta",
// 		Customer_phone:   "+998990828483",
// 	}, "8c8565c7-1707-4666-a19f-dba87f63f124")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// err := c.DeleteOrder(&models.OrderPrimaryKey{Id: "f7d3f2d3-6065-4281-b8ee-ee947a96610b"})
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// 	return
// 	// }
// }
