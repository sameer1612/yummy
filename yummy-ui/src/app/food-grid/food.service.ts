import { httpResource } from '@angular/common/http';
import { Service } from '@angular/core';

import { environment } from '../../environments/environment';
import { Food } from './food.model';

interface FoodsResponse {
  data: Food[];
}

@Service()
export class FoodService {
  readonly foods = httpResource<Food[]>(() => `${environment.apiBaseUrl}/foods`, {
    defaultValue: [],
    parse: (raw) => (raw as FoodsResponse).data
  });
}
