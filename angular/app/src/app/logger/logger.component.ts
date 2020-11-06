import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { ChatterService } from '../chatter.service';

@Component({
  selector: 'app-logger',
  templateUrl: './logger.component.html',
  styleUrls: ['./logger.component.sass']
})
export class LoggerComponent implements OnInit {

  login = new FormControl('');

  constructor(private service : ChatterService) { }

  logIn(): void {
    this.service.logAs(this.login.value).subscribe(() => 0)
  }

  ngOnInit(): void {
  }

}
