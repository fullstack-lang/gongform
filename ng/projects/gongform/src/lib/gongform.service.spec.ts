import { TestBed } from '@angular/core/testing';

import { GongformService } from './gongform.service';

describe('GongformService', () => {
  let service: GongformService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongformService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
