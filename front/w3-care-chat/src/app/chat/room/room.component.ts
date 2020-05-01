import { Component, OnInit } from '@angular/core';
import { UserModel } from 'src/app/models/user.model';
import { UserService } from 'src/app/user.service';
import { QueueService } from 'src/app/queue.service';
import { MessageService } from 'src/app/services/message.service';
import { Message } from 'src/app/models/message.model';

@Component({
  selector: 'app-room',
  templateUrl: './room.component.html',
  styleUrls: ['./room.component.scss']
})
export class RoomComponent implements OnInit {

  currentUser: UserModel;
  status: String = `WAITING`;
  messageList: string[] = [];
  newMessage: string;
  roomId: number;
  constructor(private service: UserService, private queueService: QueueService, private messageService: MessageService) { }

  ngOnInit(): void {
    this.service.currentUserModel.subscribe(u=>{
      this.currentUser = u;
      if (this.currentUser.type === 'PATIENT') {
        this.registerAsPatient();
      }
    });
  }

  startNewRoom(): void {
    this.queueService.pullNewRoom()
    .subscribe(queue=>{
      this.messageService.createAndJoinRoom(queue.body.patient.id).subscribe(message=>{
        console.log(message)
        this.roomId = message;
        this.status = 'IN_EXECUTION';
        this.messageService
      .getMessages()
      .subscribe((message: string) => {
        console.log(message);
        this.messageList.push(message);
      });
      })
    })
  }

  registerAsPatient(): void {
    console.log(this.currentUser.id)
    this.queueService.registerAsPatient()
    .subscribe(()=>{
      this.messageService.waitAndJoinRoom(this.currentUser.id).subscribe(message=>{
        this.status = 'IN_EXECUTION';
        this.roomId = message;
        this.messageService
      .getMessages()
      .subscribe((message: string) => {
        console.log(message);
        this.messageList.push(message);
      });
      });
    });
;
  }

  sendMessage() {
    this.messageService.sendMessage(new Message(this.currentUser, this.newMessage, this.roomId));
    this.newMessage = '';
  }

}
