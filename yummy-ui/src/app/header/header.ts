import { NgOptimizedImage } from '@angular/common';
import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';
import { ButtonModule } from 'primeng/button';

@Component({
  selector: 'app-header',
  imports: [NgOptimizedImage, RouterLink, ButtonModule],
  templateUrl: './header.html',
})
export class Header {}
