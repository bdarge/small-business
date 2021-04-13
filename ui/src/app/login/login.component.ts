import { Component, OnInit } from '@angular/core';
import { TranslateService} from '@ngx-translate/core';
import { FormBuilder, Validators} from '@angular/forms';
import jwt_decode from 'jwt-decode';
import {
  LocalStorageService,
  NotificationService,
  ROUTE_ANIMATIONS_ELEMENTS,
  routeAnimations
} from '../core/core.module';
import {Router} from '@angular/router';
import {AuthWebService} from '../http/auth-web.service';

@Component({
  selector: 'app-login',
  animations: [routeAnimations],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  routeAnimationsElements = ROUTE_ANIMATIONS_ELEMENTS;

  form = this.fb.group({
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required]],
  });

  constructor(
    private router: Router,
    private fb: FormBuilder,
    private translate: TranslateService,
    private notificationService: NotificationService,
    private authService: AuthWebService,
    private localStorageSvc: LocalStorageService
  ) { }

  ngOnInit() {}

  async submit() {
    if (this.form.valid) {
      this.authService.login(this.form.value)
        .subscribe((result: any) => {
          const decoded = jwt_decode(result.token);
          this.localStorageSvc.setItem('USER', decoded);
          this.localStorageSvc.setItem('TOKEN', result.token);
          this.router.navigate(['business']);
        })
    }
  }
}
