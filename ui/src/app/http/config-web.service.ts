import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import {User} from '../model/user';
import {ENVIRONMENT} from '../../environments/environment';
import {HttpClient} from '@angular/common/http';

@Injectable()
export class ConfigWebService{
  readonly USER_URL = `${ENVIRONMENT.api}/users`
  constructor(public http: HttpClient) {}

  getUser(id: string): Observable<User> {
    return this.http.get<User>(this.USER_URL + '/id/:' + id)
  }

  saveUser(user: User): Observable<User> {
    return this.http.post<User>(this.USER_URL, user)
  }
}
