import { Injectable } from '@angular/core';
import { Socket } from 'ngx-socket-io';
import { Observable } from 'rxjs';
import { Message } from '../models/message.model';

@Injectable({
  providedIn: 'root'
})
export class MessageService {
  constructor(private socket: Socket) { }

  public sendMessage(message: Message) {
    this.socket.emit('new-message', message);
}

public getMessages = () => {
  return Observable.create((observer) => {
          this.socket.on('new-message', (message) => {
              observer.next(message);
          });
  });
}

}