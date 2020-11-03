import { TestBed } from '@angular/core/testing';

import { ChatterService } from './chatter.service';

describe('ChatterService', () => {
  let service: ChatterService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ChatterService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
