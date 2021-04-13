import { Component, OnInit } from '@angular/core';
import { TranslateService} from '@ngx-translate/core';
import { FormBuilder, Validators} from '@angular/forms';
import {
  LocalStorageService,
  NotificationService,
  ROUTE_ANIMATIONS_ELEMENTS,
  routeAnimations
} from '../core/core.module';
import { Store} from '@ngrx/store';
import { Router} from '@angular/router';
import { AuthWebService} from '../http/auth-web.service';

@Component({
  selector: 'app-register',
  animations: [routeAnimations],
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {
  routeAnimationsElements = ROUTE_ANIMATIONS_ELEMENTS

  form = this.fb.group({
    email: ['', [Validators.required, Validators.email]],
    name: ['', [Validators.required]],
    password: ['', [Validators.required]],
  });

  constructor(
    private router: Router,
    private fb: FormBuilder,
    private store: Store,
    private translate: TranslateService,
    private authService: AuthWebService,
    private localStorageSvc: LocalStorageService,
    private notificationsService: NotificationService
  ) { }

  ngOnInit(): void {
  }

  async submit() {
    if (this.form.valid) {
      this.authService.register(this.form.value)
        .subscribe((result: any) => {
          this.notificationsService.info('Account created. Logging in now.')
          this.localStorageSvc.setItem('TOKEN', result.token)
          this.router.navigate(['business'])
        })
    }
  }
}
