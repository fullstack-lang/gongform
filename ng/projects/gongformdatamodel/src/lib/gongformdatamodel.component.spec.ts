import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongformdatamodelComponent } from './gongformdatamodel.component';

describe('GongformdatamodelComponent', () => {
  let component: GongformdatamodelComponent;
  let fixture: ComponentFixture<GongformdatamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongformdatamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongformdatamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
