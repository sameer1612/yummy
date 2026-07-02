import { Component, inject } from '@angular/core';
import { CardModule } from 'primeng/card';
import { MessageModule } from 'primeng/message';
import { ProgressSpinnerModule } from 'primeng/progressspinner';

import { FoodService } from './food.service';

@Component({
  selector: 'app-food-grid',
  imports: [CardModule, MessageModule, ProgressSpinnerModule],
  templateUrl: './food-grid.html'
})
export class FoodGrid {
  protected readonly foodService = inject(FoodService);
}
