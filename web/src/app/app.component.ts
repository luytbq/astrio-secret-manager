import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { AuthService } from './services/auth.service';
import { environment } from 'src/environments/environment';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';


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
      if (!!user) {
      } else {
        window.location.href = this.aasWebUrl + "/login?back=" + encodeURIComponent(window.location.href);
      }
    })
  }
}
