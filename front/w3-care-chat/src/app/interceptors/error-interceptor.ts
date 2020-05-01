import { Injectable } from '@angular/core';
import { HttpRequest, HttpHandler, HttpEvent, HttpInterceptor } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { UserService } from '../user.service';

@Injectable()
export class ErrorInterceptor implements HttpInterceptor {
    constructor(private service: UserService) { }

    intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        return next.handle(request).pipe(catchError(err => {
            if (err.status === 403) {
                // auto logout if 401 response returned from api
                this.service.logout();
                location.reload(true);
            }
            return throwError(err);
        }))
    }
}