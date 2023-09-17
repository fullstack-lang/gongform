import { TestBed } from '@angular/core/testing';

import { GongformspecificService } from './gongformspecific.service';

describe('GongformspecificService', () => {
  let service: GongformspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongformspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
