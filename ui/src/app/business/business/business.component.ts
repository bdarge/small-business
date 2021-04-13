import {ChangeDetectionStrategy, Component, OnInit} from '@angular/core';
import {from, Observable} from 'rxjs';
import {select, Store} from '@ngrx/store';
import {
  routeAnimations,
  selectSettingsLanguage,
  selectTheme,

  LocalStorageService
} from '../../core/core.module';
import {actionSettingsChangeLanguage} from '../../core/settings/settings.actions';
import {State} from '../../core/settings/settings.model';
import {OverlayContainer} from '@angular/cdk/overlay';
import {User} from '../../model/user';
import {Router} from '@angular/router';

@Component({
  selector: 'app-business',
  animations: [routeAnimations],
  templateUrl: './business.component.html',
  styleUrls: ['./business.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class BusinessComponent implements OnInit {
  navigation = [
    { link: 'home', label: 'business.menu.home' },
    { link: 'profile', label: 'business.menu.profile' },
    { link: 'setting', label: 'business.menu.setting' }
  ];
  navigationSideMenu = [
    ...this.navigation
  ];

  theme$: Observable<string>;
  isAuthenticated$: Observable<boolean>;
  user$: Observable<string>;

  constructor(
    private router: Router,
    private overlayContainer: OverlayContainer,
    private store: Store<State>,
    private localStorageSvc: LocalStorageService,
  ) {
    this.theme$ = this.store.pipe(select(selectTheme));
    this.isAuthenticated$ = from(new Promise<boolean>((resolve) =>{
      resolve(this.localStorageSvc.getItem('E-ID') ||
        this.localStorageSvc.getItem('TOKEN'))
    }));
    this.user$ = from(new Promise<string>((resolve) => {
      const user = this.localStorageSvc.getItem('USER') as User;
      if (user) {
        resolve(user.ownerName)
      }
    }));
  }

  ngOnInit() {
  }

  onLanguageSelect({ value: language }) {
    this.store.dispatch(actionSettingsChangeLanguage({ language }));
  }

  onLogoutClick() {
    this.localStorageSvc.removeItem('E-ID');
    this.localStorageSvc.removeItem('USER');
    // TODO remove token for web
    this.router.navigate(['login']);
  }
}

