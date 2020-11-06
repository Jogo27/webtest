import { Injectable } from '@angular/core';

import { HttpClient, HttpParams } from '@angular/common/http';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ChatterService {

  loginUrl = 'login';
  greetUrl = 'greet/';

  sessionPath = '';

  constructor(private http : HttpClient) { }

  logAs(user : string) {
    var oldPath = this.sessionPath;
    return this.http.get(this.loginUrl, {
      params: new HttpParams().set('login', user),
      observe: "body",
      responseType: "text"
    }).pipe(map(
      result => {
        this.sessionPath = result;
        return this.sessionPath != oldPath;
      })
    );
  }

  getNewMessage() {
    return this.http.get(this.greetUrl + this.sessionPath, {
      observe: 'response', responseType: 'text' });
  }

  isLogged() : boolean {
    return this.sessionPath != '';
  }
}
