import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterComponent } from './register.component';
import { LocalStorageService} from '../core/local-storage/local-storage.service';
import { NotificationService} from '../core/notifications/notification.service';
import { AuthWebService} from '../http/auth-web.service';
import {of} from 'rxjs';
import {Store} from '@ngrx/store';
import {SharedModule} from '../shared/shared.module';
import {FontAwesomeIconsModule} from '../shared/font.awesome.icons.module';
import {NoopAnimationsModule} from '@angular/platform-browser/animations';
import {HttpClientTestingModule} from '@angular/common/http/testing';
import {TranslateModule} from '@ngx-translate/core';
import {RouterTestingModule} from '@angular/router/testing';

describe('RegisterComponent', () => {
  let component: RegisterComponent;
  let fixture: ComponentFixture<RegisterComponent>;
  let abstractServiceStub: Partial<AuthWebService>;
  const testStore = jasmine.createSpyObj('Store', ['pipe']);

  beforeEach(async () => {
    testStore.pipe.and.returnValue(of(''))
    await TestBed.configureTestingModule({
      imports: [
        SharedModule,
        FontAwesomeIconsModule,
        NoopAnimationsModule,
        HttpClientTestingModule,
        TranslateModule.forRoot(),
        RouterTestingModule
      ],
      declarations: [
        RegisterComponent
      ],
      providers: [
        {
          provide: Store, useValue: testStore
        },
        {
          provide: LocalStorageService
        },
        {
          provide: NotificationService
        },
        {
          provide: AuthWebService, useValue: abstractServiceStub
        }]
    })
      .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RegisterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
