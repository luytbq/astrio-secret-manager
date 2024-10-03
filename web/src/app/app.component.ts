import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { provideRouter, RouterOutlet, withDebugTracing, withPreloading } from '@angular/router';
import { AuthService } from './services/auth.service';
import { environment } from 'src/environments/environment';
import { APP_ROUTES } from './app.route';
import { HTTP_INTERCEPTORS, HttpClient } from '@angular/common/http';
import { CookieService } from 'ngx-cookie-service';
import { AuthInterceptor } from './services/http-interceptor';

@Component({
  standalone: true,
  imports: [
    CommonModule,
    RouterOutlet,
  ],
  providers: [
  ],
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  aasWebUrl = environment.aas.web_url;

  constructor( 
    private auth: AuthService,
  ) {
  }

  ngOnInit(): void {
    this.auth.getInfos().subscribe(user => {
      console.log('user', user);

      if (!!user) {
        console.log('logged in');
      } else {
        window.location.href = this.aasWebUrl + "/login?back=" + encodeURIComponent(window.location.href);
      }
    })
  }
}
