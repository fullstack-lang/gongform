import { TestBed } from '@angular/core/testing';

import { GongformdatamodelService } from './gongformdatamodel.service';

describe('GongformdatamodelService', () => {
  let service: GongformdatamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongformdatamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
