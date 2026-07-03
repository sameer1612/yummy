package repository

import (
	"gorm.io/gorm"
)

type FoodItem struct {
	ID        int32 `gorm:"primaryKey"`
	Name      string
	Caption   string
	Rating    *float64
	PhotoPath string
}

type FoodRepository interface {
	List() ([]FoodItem, error)
	GetByID(id int32) (FoodItem, error)
	Create(name, caption, photoPath string, rating *float64) (FoodItem, error)
	Update(id int32, name, caption, photoPath string, rating *float64) (FoodItem, error)
	Delete(id int32) (FoodItem, error)
}

type repository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) FoodRepository {
	return &repository{db: db}
}

func (r *repository) List() ([]FoodItem, error) {
	var items []FoodItem
	result := r.db.Order("created_at").Find(&items)
	return items, result.Error
}

func (r *repository) GetByID(id int32) (FoodItem, error) {
	var item FoodItem
	result := r.db.First(&item, id)
	return item, result.Error
}

func (r *repository) Create(name, caption, photoPath string, rating *float64) (FoodItem, error) {
	item := FoodItem{
		Name:      name,
		Caption:   caption,
		PhotoPath: photoPath,
		Rating:    rating,
	}
	result := r.db.Create(&item)
	return item, result.Error
}

func (r *repository) Update(id int32, name, caption, photoPath string, rating *float64) (FoodItem, error) {
	var item FoodItem
	if result := r.db.First(&item, id); result.Error != nil {
		return item, result.Error
	}
	item.Name = name
	item.Caption = caption
	item.PhotoPath = photoPath
	item.Rating = rating
	result := r.db.Save(&item)
	return item, result.Error
}

func (r *repository) Delete(id int32) (FoodItem, error) {
	var item FoodItem
	if result := r.db.First(&item, id); result.Error != nil {
		return item, result.Error
	}
	result := r.db.Delete(&item)
	return item, result.Error
}
