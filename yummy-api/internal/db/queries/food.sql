-- name: ListFoods :many
select * from food_items;

-- name: GetFoodItem :one
select * from food_items where id = $1;

-- name: CreateFoodItem :one
INSERT INTO food_items (name, caption, rating, photo_path) 
VALUES ($1, $2, $3, $4) 
RETURNING *;