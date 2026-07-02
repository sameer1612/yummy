import { NgOptimizedImage } from '@angular/common';
import { Component, computed, inject, input, linkedSignal, signal } from '@angular/core';
import { form, FormField, required, submit } from '@angular/forms/signals';
import { Router, RouterLink } from '@angular/router';
import { ConfirmationService, MessageService } from 'primeng/api';
import { ButtonModule } from 'primeng/button';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { InputTextModule } from 'primeng/inputtext';
import { MessageModule } from 'primeng/message';
import { SliderModule } from 'primeng/slider';
import { ToastModule } from 'primeng/toast';

import { FoodService } from '../food-grid/food.service';

@Component({
  selector: 'app-food-form',
  imports: [
    ButtonModule,
    ConfirmDialogModule,
    InputTextModule,
    MessageModule,
    SliderModule,
    ToastModule,
    FormField,
    NgOptimizedImage,
    RouterLink,
  ],
  providers: [ConfirmationService, MessageService],
  templateUrl: './food-form.html',
})
export class FoodForm {
  readonly id = input<string>();

  protected readonly foodService = inject(FoodService);
  private readonly router = inject(Router);
  private readonly messageService = inject(MessageService);
  private readonly confirmationService = inject(ConfirmationService);

  protected readonly isEditMode = computed(() => this.id() !== undefined);

  protected readonly existingFood = computed(() =>
    this.isEditMode()
      ? this.foodService.foods.value().find((food) => food.id === Number(this.id()))
      : undefined,
  );

  protected readonly notFound = computed(
    () => this.isEditMode() && !this.foodService.foods.isLoading() && !this.existingFood(),
  );

  protected readonly model = linkedSignal(() => {
    const food = this.existingFood();
    return {
      name: food?.name ?? '',
      caption: food?.caption ?? '',
      rating: food?.rating ?? 0,
      photo_path: food ? new URL(food.photo_path).pathname : '',
    };
  });

  protected readonly foodForm = form(this.model, (path) => {
    required(path.name, { message: 'Name is required.' });
    required(path.caption, { message: 'Caption is required.' });
  });

  protected readonly submitting = signal(false);

  protected onSubmit(event: Event) {
    event.preventDefault();
    submit(this.foodForm, async () => {
      this.submitting.set(true);
      try {
        const value = this.model();
        const payload = {
          name: value.name,
          caption: value.caption,
          rating: value.rating > 0 ? value.rating : null,
          photo_path: value.photo_path,
        };
        if (this.isEditMode()) {
          await this.foodService.update(Number(this.id()), payload);
        } else {
          await this.foodService.create(payload);
        }
        this.router.navigate(['/foods']);
      } catch {
        this.messageService.add({
          severity: 'error',
          summary: 'Error',
          detail: 'Something went wrong. Please try again.',
        });
      } finally {
        this.submitting.set(false);
      }
    });
  }

  protected onDelete(event: Event) {
    this.confirmationService.confirm({
      target: event.target as EventTarget,
      message: `Delete "${this.existingFood()?.name}"? This cannot be undone.`,
      header: 'Delete food item',
      icon: 'pi pi-exclamation-triangle',
      rejectButtonProps: { label: 'Cancel', severity: 'secondary', outlined: true },
      acceptButtonProps: { label: 'Delete', severity: 'danger' },
      accept: async () => {
        try {
          await this.foodService.delete(Number(this.id()));
          this.router.navigate(['/foods']);
        } catch {
          this.messageService.add({
            severity: 'error',
            summary: 'Error',
            detail: 'Could not delete this item.',
          });
        }
      },
    });
  }
}
