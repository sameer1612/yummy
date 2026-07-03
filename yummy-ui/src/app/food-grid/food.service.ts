import { HttpClient, httpResource } from '@angular/common/http';
import { inject, Service } from '@angular/core';
import { firstValueFrom } from 'rxjs';

import { environment } from '../../environments/environment';
import { CreateFoodInput, Food, UpdateFoodInput } from './food.model';

interface FoodsResponse {
  data: Food[];
}

@Service()
export class FoodService {
  private readonly http = inject(HttpClient);
  private readonly baseUrl = `${environment.apiBaseUrl}/foods`;

  readonly foods = httpResource<Food[]>(() => this.baseUrl, {
    defaultValue: [],
    parse: (raw) => (raw as FoodsResponse).data,
  });

  async create(input: CreateFoodInput): Promise<void> {
    const formData = new FormData();
    formData.append('name', input.name);
    formData.append('caption', input.caption);
    if (input.rating !== null) {
      formData.append('rating', String(input.rating));
    }
    formData.append('photo', input.photo);
    await firstValueFrom(this.http.post<void>(this.baseUrl, formData));
    this.foods.reload();
  }

  async update(id: number, input: UpdateFoodInput): Promise<void> {
    const formData = new FormData();
    formData.append('name', input.name);
    formData.append('caption', input.caption);
    if (input.rating !== null) {
      formData.append('rating', String(input.rating));
    }
    if (input.photo) {
      formData.append('photo', input.photo);
    }
    await firstValueFrom(this.http.put<void>(`${this.baseUrl}/${id}`, formData));
    this.foods.reload();
  }

  async delete(id: number): Promise<void> {
    await firstValueFrom(this.http.delete<void>(`${this.baseUrl}/${id}`));
    this.foods.reload();
  }
}
