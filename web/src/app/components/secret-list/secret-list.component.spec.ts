import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SecretListComponent } from './secret-list.component';

describe('SecretListComponent', () => {
  let component: SecretListComponent;
  let fixture: ComponentFixture<SecretListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SecretListComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SecretListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
