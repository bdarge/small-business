import { Injectable } from '@angular/core';
import {ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot} from '@angular/router';
import {from, Observable} from 'rxjs';
import {HttpClient} from '@angular/common/http';
import {LocalStorageService} from '../local-storage/local-storage.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuardService implements CanActivate {
  constructor(private router: Router,
              private http: HttpClient,
              private localStorageSvc: LocalStorageService) {}

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean> {
    return from(new Promise<boolean>((resolve, reject) => {
      const token = this.localStorageSvc.getItem('TOKEN')
      if (!token) {
        return this.router.navigate(['/login'])
      } else {
        const isExpired = this.tokenExpired(token);
        if (isExpired) {
          return this.router.navigate(['/login'])
        } else {
          resolve(true)
        }
      }
    }))
  }

  private tokenExpired(token: string) {
    const expiry = (JSON.parse(atob(token.split('.')[1]))).exp;
    return (Math.floor((new Date).getTime() / 1000)) >= expiry;
  }
}
