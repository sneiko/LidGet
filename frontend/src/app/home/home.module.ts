import { MatProgressBarModule} from '@angular/material/progress-bar';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home.component';
import {MatTableModule} from "@angular/material/table";


@NgModule({
  declarations: [
    HomeComponent
  ],
  exports: [
    HomeComponent
  ],
  imports: [
    CommonModule,
    MatProgressBarModule,
    MatTableModule,
  ]
})
export class HomeModule { }
