import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongformspecificComponent } from './gongformspecific.component';

describe('GongformspecificComponent', () => {
  let component: GongformspecificComponent;
  let fixture: ComponentFixture<GongformspecificComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [GongformspecificComponent]
    });
    fixture = TestBed.createComponent(GongformspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
