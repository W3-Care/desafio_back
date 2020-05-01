import { Component } from '@angular/core';
import { MessageService } from './services/message.service';
import { Message } from './models/message.model';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  newMessage: string;
  messageList:  string[] = [];
  id=  Math.random();
  
  constructor(private chatService: MessageService) {
  }
  
}
