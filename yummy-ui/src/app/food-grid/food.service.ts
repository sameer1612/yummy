import { HttpClient, httpResource } from '@angular/common/http';
import { inject, Service } from '@angular/core';
import { firstValueFrom } from 'rxjs';

import { environment } from '../../environments/environment';
import { Food, FoodInput } from './food.model';

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

  async create(input: FoodInput): Promise<void> {
    await firstValueFrom(this.http.post<void>(this.baseUrl, input));
    this.foods.reload();
  }

  async update(id: number, input: FoodInput): Promise<void> {
    await firstValueFrom(this.http.put<void>(`${this.baseUrl}/${id}`, input));
    this.foods.reload();
  }

  async delete(id: number): Promise<void> {
    await firstValueFrom(this.http.delete<void>(`${this.baseUrl}/${id}`));
    this.foods.reload();
  }
}
