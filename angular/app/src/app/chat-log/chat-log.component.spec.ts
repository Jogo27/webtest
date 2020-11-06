import { ComponentFixture, TestBed } from '@angular/core/testing';
import { of } from 'rxjs';

import { ChatLogComponent } from './chat-log.component';
import { ChatterService } from '../chatter.service';

describe('ChatLogComponent', () => {
  let chatterService : { getNewMessage: jasmine.Spy};
  let component: ChatLogComponent;
  let fixture: ComponentFixture<ChatLogComponent>;

  beforeEach(() => {
    chatterService = jasmine.createSpyObj('ChatterService', ['getNewMessage']);
    TestBed.configureTestingModule({
      declarations: [ ChatLogComponent ],
      providers: [{provide: ChatterService, useValue: chatterService}]
    });
    fixture = TestBed.createComponent(ChatLogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('add calls getNewMessage', () => {
    const expectedMsg = 'Youhou';
    chatterService.getNewMessage.and.returnValue(of({body: expectedMsg}));
    component.add();
    expect(chatterService.getNewMessage.calls.any()).toBeTrue();
    expect(component.logs.some((value, index, array) => value.includes('Youhou')));
  });
});
