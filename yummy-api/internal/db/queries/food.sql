-- name: ListFoods :many
SELECT * FROM food_items ORDER BY created_at;

-- name: GetFoodItem :one
SELECT * FROM food_items where id = $1;

-- name: CreateFoodItem :one
INSERT INTO food_items (name, caption, rating, photo_path) 
VALUES ($1, $2, $3, $4) 
RETURNING *;

-- name: UpdateFoodItem :one
UPDATE food_items 
SET name = $1, caption = $2, rating = $3, photo_path = $4
WHERE id = $5
RETURNING *;

-- name: DeleteFoodItem :one
DELETE FROM food_items WHERE id = $1 RETURNING *;
