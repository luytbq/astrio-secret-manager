import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { catchError, map, of } from 'rxjs';
import { Observable } from 'rxjs/internal/Observable';
import { environment } from 'src/environments/environment';

@Injectable({ providedIn: 'root' })
export class AuthService {
  asmServiceUrl = environment.asm.service_url;

  constructor(
    private http: HttpClient,
    private cookieService: CookieService,
  ) {
  }

  getInfos(): Observable<User | null> {
    return this.http.get(this.asmServiceUrl+"/users/infos", {observe: 'response'}).pipe(
      map(res => res.body as User),
      catchError(err => of(null))
    );
  }

  getHeader(): HttpHeaders {
    const authCookie = this.cookieService.get("auth");
    const header = new HttpHeaders();
    header.append("Authorization", "Bearer " + authCookie);
    return header;
  }
}

export interface User {
  user_id: string,
  email: string
}