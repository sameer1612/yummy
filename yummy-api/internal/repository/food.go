package repository

import (
	"database/sql"
	"yummy/internal/db/jet/yummy/public/model"
	"yummy/internal/db/jet/yummy/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

var tbl = table.FoodItems

type FoodRepository interface {
	List() ([]model.FoodItems, error)
	GetByID(id int32) (model.FoodItems, error)
	Create(name, caption, photoPath string, rating *float64) (model.FoodItems, error)
	Update(id int32, name, caption, photoPath string, rating *float64) (model.FoodItems, error)
	Delete(id int32) (model.FoodItems, error)
}

type repository struct {
	db *sql.DB
}

func NewFoodRepository(db *sql.DB) FoodRepository {
	return &repository{db: db}
}

func (r *repository) List() ([]model.FoodItems, error) {
	var results []model.FoodItems
	err := SELECT(tbl.AllColumns).FROM(tbl).Query(r.db, &results)
	return results, err
}

func (r *repository) GetByID(id int32) (model.FoodItems, error) {
	var result model.FoodItems
	err := SELECT(tbl.AllColumns).FROM(tbl).WHERE(tbl.ID.EQ(Int32(id))).Query(r.db, &result)
	return result, err
}

func (r *repository) Create(name, caption, photoPath string, rating *float64) (model.FoodItems, error) {
	var result model.FoodItems
	err := tbl.INSERT(tbl.Name, tbl.Caption, tbl.Rating, tbl.PhotoPath).
		MODEL(model.FoodItems{Name: name, Caption: caption, Rating: rating, PhotoPath: photoPath}).
		RETURNING(tbl.AllColumns).
		Query(r.db, &result)
	return result, err
}

func (r *repository) Update(id int32, name, caption, photoPath string, rating *float64) (model.FoodItems, error) {
	var result model.FoodItems
	err := tbl.UPDATE(tbl.Name, tbl.Caption, tbl.Rating, tbl.PhotoPath).
		MODEL(model.FoodItems{Name: name, Caption: caption, Rating: rating, PhotoPath: photoPath}).
		WHERE(tbl.ID.EQ(Int32(id))).
		RETURNING(tbl.AllColumns).
		Query(r.db, &result)
	return result, err
}

func (r *repository) Delete(id int32) (model.FoodItems, error) {
	var result model.FoodItems
	err := tbl.DELETE().WHERE(tbl.ID.EQ(Int32(id))).RETURNING(tbl.AllColumns).Query(r.db, &result)
	return result, err
}
