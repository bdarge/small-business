import { Injectable } from '@angular/core';
import { HttpClient} from '@angular/common/http';
import { ENVIRONMENT } from '../../environments/environment';
import {Observable} from 'rxjs';
import {User} from '../model/user';

@Injectable({
  providedIn: 'root'
})
export class AuthWebService {
  readonly AUTH_URL = `${ENVIRONMENT.api}/auth/`
  constructor(private http: HttpClient) {
  }

  login({email, password}) {
    const file = {username: email, password}
    const options = {headers: {'Content-Type': 'application/json'}}
    return this.http.post(this.AUTH_URL + 'login', JSON.stringify(file), options)
  }

  register(value: any): Observable<User> {
    return this.http.post<User>(this.AUTH_URL + 'register', JSON.stringify(value));
  }
}
