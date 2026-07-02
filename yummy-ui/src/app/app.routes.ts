import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('./home/home').then((m) => m.Home),
  },
  {
    path: 'foods',
    loadComponent: () => import('./food-grid/food-grid').then((m) => m.FoodGrid),
  },
  {
    path: 'add-item',
    loadComponent: () => import('./food-form/food-form').then((m) => m.FoodForm),
  },
  {
    path: 'foods/:id/edit',
    loadComponent: () => import('./food-form/food-form').then((m) => m.FoodForm),
  },
];
