package services

import (
	"errors"
	"myapi/db"
	"myapi/models"
	"myapi/repositories"

	"gorm.io/gorm"
)

type OrderService struct {
	repo *repositories.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{
		repo: &repositories.OrderRepository{},
	}
}

func (s *OrderService) GetAll() ([]models.Order, error) {
	return s.repo.FindAll()
}

func (s *OrderService) GetByID(id int) (*models.Order, error) {
	return s.repo.FindByID(id)
}

func (s *OrderService) Create(req models.OrderPayload, userID uint) (*models.Order, error) {
	var order models.Order

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		var totalPrice float64
		var items []models.OrderItem

		for _, item := range req.Items {
			var product models.Product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				return errors.New("product not found")
			}

			if product.Quantity < item.Quantity {
				return errors.New("Insufficient stock for product: " + product.Name)
			}

			subTotal := product.Price * float64(item.Quantity)
			totalPrice += subTotal

			items = append(items, models.OrderItem{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     product.Price,
			})
		}

		order = models.Order{
			UserID:     userID,
			TotalPrice: totalPrice,
			Status:     "pending",
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		for i := range items {
			items[i].OrderID = order.ID
			if err := tx.Create(&items[i]).Error; err != nil {
				return err
			}

			s.repo.DecreaseStock(tx, items[i].ProductID, items[i].Quantity)
		}

		tx.Preload("Items.Product.Categories").Preload("Items.Product").First(&order, order.ID)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *OrderService) Pay(id int) (*models.Order, error) {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if order.Status != "pending" {
		return nil, errors.New("transaction status should be pending")
	}

	order.Status = "paid"

	return s.repo.Save(order)
}

func (s *OrderService) Cancel(id int) (*models.Order, error) {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("order not found")
	}

	if order.Status != "pending" {
		return nil, errors.New("transaction status should be pending")
	}

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		for _, item := range order.Items {
			s.repo.IncreaseStock(tx, item.ProductID, item.Quantity)
		}

		order.Status = "cancelled"

		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		return nil
	})

	return order, err
}

func (s *OrderService) Delete(id int) (*models.Order, error) {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("order not found")
	}

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		return s.repo.DeleteOrder(tx, order.ID)
	})

	return order, err
}
