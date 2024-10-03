import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { CookieService } from "ngx-cookie-service";
import { Observable } from "rxjs";

@Injectable({ providedIn: 'root' })
export class AuthInterceptor implements HttpInterceptor {
    constructor(
        private cookieService: CookieService,
    ) {
    }

    intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        const authCookie = this.cookieService.get("auth");

        const clonedReq = authCookie ? req.clone({
            headers: req.headers.set("Authorization", "Bearer " + authCookie)
        }) : req;


        return next.handle(clonedReq);
    }
}
