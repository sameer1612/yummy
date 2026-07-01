create table food_items (
  id serial primary key,
  name text not null,
  caption text not null,
  rating float8 check (rating >= 0 and rating <= 5),
  photo_path text not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);
