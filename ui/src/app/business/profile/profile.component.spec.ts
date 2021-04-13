import { waitForAsync, ComponentFixture, TestBed } from '@angular/core/testing';
import { ProfileComponent } from './profile.component';
import {SharedModule} from '../../shared/shared.module';
import {FontAwesomeIconsModule} from '../../shared/font.awesome.icons.module';
import {NoopAnimationsModule} from '@angular/platform-browser/animations';
import {HttpClientTestingModule} from '@angular/common/http/testing';
import {TranslateModule} from '@ngx-translate/core';
import {NotificationService} from '../../core/notifications/notification.service';
import {MatDialog} from '@angular/material/dialog';
import {LocalStorageService} from '../../core/local-storage/local-storage.service';
import {of} from 'rxjs';
import {User} from '../../model/user';
import {ConfigWebService} from '../../http/config-web.service';

describe('ProfileComponent', () => {
  let component: ProfileComponent;
  let fixture: ComponentFixture<ProfileComponent>;
  let getUserSpy;

  beforeEach(waitForAsync(() => {
    const user  = {
      id: 'id',
      email: 'test@email.com'
    } as User
    const configService = jasmine.createSpyObj('ConfigWebService', ['getUser']);
    getUserSpy = configService.getUser.and.returnValue(of(user));

    TestBed.configureTestingModule({
      imports: [
        SharedModule,
        FontAwesomeIconsModule,
        NoopAnimationsModule,
        HttpClientTestingModule,
        TranslateModule.forRoot()
      ],
      declarations: [ProfileComponent],
      providers: [
        {
          provide: NotificationService
        },
        {
          provide: LocalStorageService
        },
        {
          provide: ConfigWebService
        },
        {
          provide: MatDialog
        }]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
