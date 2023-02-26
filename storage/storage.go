package storage

import (
	"app/models"
)

type StorageI interface {
	CloseDb()
	User() UserRepoI
	Product() ProductRepoI
	ShopCart() ShopCartRepoI
	Commission() CommissionRepoI
	Category() CategoryRepoI
	Order() OrderRepoI
	Branch() BranchRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (string, error)
	Delete(*models.UserPrimaryKey) error
	Update(*models.UpdateUser, string) error
	GetByID(*models.UserPrimaryKey) (models.User, error)
	GetAll(*models.GetListRequest) (models.GetListResponse, error)
}

type ProductRepoI interface {
	Create(*models.CreateProduct) (string, error)
	GetByID(*models.ProductPrimaryKey) (models.ProductWithCategory, error)
	GetAll(*models.GetListProductRequest) (models.GetListProduct, error)
	Update(*models.UpdateProduct, string) error
	Delete(*models.ProductPrimaryKey) error
}

type ShopCartRepoI interface {
	AddShopCart(*models.Add) (string, error)
	RemoveShopCart(*models.Remove) error
	GetUserShopCart(*models.UserPrimaryKey) ([]models.ShopCart, error)
	UpdateShopCart(string) error
	GetAllShopCarts() ([]models.ShopCart, error)
	Filter(from_date, to_date string) ([]models.ShopCart, error)
	// DTSProducts(*models.UserPrimaryKey) ([]models.ShopCart, error)
}

type CommissionRepoI interface {
	AddCommission(*models.Commission) error
}

type CategoryRepoI interface {
	Create(*models.CreateCategory) (string, error)
	GetByID(*models.CategoryPrimaryKey) (models.Category, error)
	GetAll(*models.GetListCategoryRequest) (models.GetListCategoryResponse, error)
	Update(*models.UpdateCategory, string) error
	Delete(*models.CategoryPrimaryKey) error
}

type OrderRepoI interface {
	Create(req *models.CreateOrder, total float64) (string, error)
	GetByOrderID(id string) (models.GetOrderById, error)
	Update(*models.UpdateOrder, string) error
	Delete(*models.OrderPrimaryKey) error
}

type BranchRepoI interface {
	Create(req *models.CreateBranch) (string, error)
	Delete(*models.BranchPrimaryKey) error
	Update(*models.UpdateBranch, string) error
	GetByID(*models.BranchPrimaryKey) (models.Branch, error)
	GetAll(*models.GetBranchListRequest) (models.GetBranchListResponse, error)
}
