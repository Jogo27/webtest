import { HttpClient, HttpResponse } from '@angular/common/http';
import { of } from 'rxjs';

import { ChatterService } from './chatter.service';

describe('ChatterService', () => {
  let httpSpy : {get: jasmine.Spy};
  let service: ChatterService;

  beforeEach(() => {
    httpSpy = jasmine.createSpyObj('HttpClient', ['get']);
    service = new ChatterService(httpSpy as any);
  });

  it('is created unlogged', () => {
    expect(service).toBeTruthy();
    expect(service.isLogged()).toBeFalse();
  });

  it('does not log with getNewMessage', () => {
    httpSpy.get.and.returnValue(of(new HttpResponse({body: "Hi"})));
    service.getNewMessage().subscribe(result => 0);
    expect(service.isLogged()).toBeFalse();
  });

  it('logs', () => {
    httpSpy.get.and.returnValue(of('AbCd'));
    service.logAs("Toto").subscribe(
      result => expect(result).toBeTrue()
    );
  });
});
