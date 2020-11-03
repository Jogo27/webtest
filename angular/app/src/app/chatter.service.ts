import { Injectable } from '@angular/core';

import { HttpClient } from '@angular/common/http';
import { HttpErrorResponse, HttpResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ChatterService {

  greetUrl = 'greet/';

  constructor(private http : HttpClient) { }

  getNewMessage() {
    return this.http.get(this.greetUrl, { observe: 'response', responseType: 'text' });
  }
}
