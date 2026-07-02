export interface Food {
  id: number;
  name: string;
  caption: string;
  rating: number | null;
  photo_url: string;
}

export interface FoodInput {
  name: string;
  caption: string;
  rating: number | null;
  photo_path: string;
}
