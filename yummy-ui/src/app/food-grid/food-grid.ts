import { NgOptimizedImage } from '@angular/common';
import { Component, inject } from '@angular/core';
import { RouterLink } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { MessageModule } from 'primeng/message';
import { ProgressSpinnerModule } from 'primeng/progressspinner';

import { FoodService } from './food.service';

@Component({
  selector: 'app-food-grid',
  imports: [
    ButtonModule,
    CardModule,
    MessageModule,
    ProgressSpinnerModule,
    NgOptimizedImage,
    RouterLink,
  ],
  templateUrl: './food-grid.html',
})
export class FoodGrid {
  protected readonly foodService = inject(FoodService);
}
