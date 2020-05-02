import { Injectable } from '@angular/core';
import { Socket } from 'ngx-socket-io';
import { Observable } from 'rxjs';
import { Message } from '../models/message.model';
import { UserService } from '../user.service';

@Injectable({
  providedIn: 'root'
})
export class MessageService {
  constructor(private socket: Socket, private userService: UserService) { }

  public sendMessage(message: Message) {
    console.log(2);
    this.socket.emit('new-message', message);
}

public getMessages(): Observable<any> {
  return Observable.create((observer) => {
          this.socket.on('new-message', (message) => {
              observer.next(message);
          });
  });
}

public createAndJoinRoom(room) {
  this.socket.emit('new-room', room);
  return Observable.create((observer) => {
    this.socket.on('new-room', (message) => {
      this.socket.emit('join-room', room);
        observer.next(message);
    });
});
}

public waitAndJoinRoom(room) {
  this.socket.emit('wait-room', room);
  
  return Observable.create((observer) => {
    this.socket.on('new-room', (message) => {
      if(message === this.userService.currentUserValue.id) {
        this.socket.emit('join-room', room);
        observer.next(message);
      }
    });
});
}


public waitForCloseRoom(room) {
  return Observable.create((observer) => {
    this.socket.on('close-room', (message) => {
      if(message === room) {
        observer.next(message);
      }
    });
});
}

public closeRoom(roomId) {
  this.socket.emit('close-room', roomId);
  return Observable.create((observer) => {
//     this.socket.on('new-room', (message) => {
//       console.log(message)
//       if(message === this.userService.currentUserValue.id) {
//         this.socket.emit('join-room', room);
//         observer.next(message);
//       }
//     });
 });
}
}