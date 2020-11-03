import { Component, OnInit } from '@angular/core';

import { HttpErrorResponse, HttpResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';

import { ChatterService } from '../chatter.service';

@Component({
  selector: 'app-chat-log',
  templateUrl: './chat-log.component.html',
  styleUrls: ['./chat-log.component.sass']
})
export class ChatLogComponent implements OnInit {

  logs = ["Initial message"];

  constructor(private chatService : ChatterService) { }

  ngOnInit(): void {
  }

  add(): void {
    this.chatService.getNewMessage().subscribe(
      response => this.logs.push(response.body),
      error => {
        if (error.error instanceof ErrorEvent) {
          console.error("Argh: " + error.error.message);
        }
        else {
          this.logs.push("Error " + error.status + ': ' + error.error)
        }
      }
   )
  }

}
