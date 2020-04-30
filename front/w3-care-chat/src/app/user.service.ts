import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { UserModel } from 'src/app/models/user.model';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { map, first } from 'rxjs/operators';
import { BehaviorSubject, Observable } from 'rxjs';
import * as jwt_decode from 'jwt-decode';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private currentUserModelSubject: BehaviorSubject<UserModel>;
  public currentUserModel: Observable<UserModel>;

  isLoggedIn = false;
  save(user: UserModel) {
    return this.http.post(`${environment.api}/api/usuarios/`, user, {observe: 'response'});
  }

  auth(email: string, password: string) {
    return this.http.post<any>(`${environment.api}/api/autentica/`, {email: email, password: password}, {observe: 'response' as 'body'})
    .pipe(
      map(jwt => {
      // store user details and jwt token in local storage to keep user logged in between page refreshes
      console.log(jwt)
      const decodedJwt = jwt_decode(jwt.headers.get('authorization'));
      console.log(decodedJwt);
      const user: UserModel = JSON.parse(decodedJwt.details);
      user.token = jwt.headers.get('authorization')
      console.log(user)
      localStorage.setItem('currentUser', JSON.stringify(user));
      this.currentUserModelSubject.next(user);
      return user;
    }));
  }

  constructor(private http: HttpClient) {
    this.currentUserModelSubject = new BehaviorSubject<UserModel>(JSON.parse(localStorage.getItem('currentUser')));
    this.currentUserModel = this.currentUserModelSubject.asObservable();
}
public get currentUserValue(): UserModel {
  return this.currentUserModelSubject.value;
}
logout() {
  // remove user from local storage to log user out
  localStorage.removeItem('currentUser');
  this.currentUserModelSubject.next(null);
}
}
