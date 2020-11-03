import { Injectable } from '@angular/core';

import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpErrorResponse, HttpResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';

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
    this.http.get(this.loginUrl, {
      params: new HttpParams().set('login', user),
      observe: "body",
      responseType: "text"
    }).subscribe(
      result => this.sessionPath = result
    );
    return this.sessionPath == oldPath;
  }

  getNewMessage() {
    return this.http.get(this.greetUrl + this.sessionPath, {
      observe: 'response', responseType: 'text' });
  }
}
