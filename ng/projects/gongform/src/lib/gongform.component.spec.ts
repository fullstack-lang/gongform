import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongformComponent } from './gongform.component';

describe('GongformComponent', () => {
  let component: GongformComponent;
  let fixture: ComponentFixture<GongformComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [GongformComponent]
    });
    fixture = TestBed.createComponent(GongformComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
