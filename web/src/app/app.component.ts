import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { RouterOutlet } from '@angular/router';
import { SecretListComponent } from './components/secret-list/secret-list.component';

@Component({
  standalone: true,
  imports: [
    CommonModule,
    RouterOutlet,
  ],
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'asm-web';
}
