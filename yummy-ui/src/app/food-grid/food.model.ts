export interface Food {
  id: number;
  name: string;
  caption: string;
  rating: number | null;
  photo_path: string;
}

export interface CreateFoodInput {
  name: string;
  caption: string;
  rating: number | null;
  photo: File;
}

export interface UpdateFoodInput {
  name: string;
  caption: string;
  rating: number | null;
  photo?: File;
}
