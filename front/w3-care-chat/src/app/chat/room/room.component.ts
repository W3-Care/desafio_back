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
  queue: any;
  infoMessage: String;

  constructor(private service: UserService, private queueService: QueueService, private messageService: MessageService) { }

  ngOnInit(): void {
    this.currentUser = this.service.currentUserValue;
    this.registerAsPatient();
  }

  startNewRoom(): void {
    this.infoMessage = "";
    this.queueService.pullNewRoom()
    .subscribe(queue=>{
      this.queue = queue.body;
      this.messageService.createAndJoinRoom(queue.body.patient.id).subscribe(message=>{
        this.roomId = message;
        this.status = 'IN_EXECUTION';
        this.messageService
      .getMessages()
      .subscribe((message: string) => {
        this.messageList.push(message);
      });
      })
      this.messageService.waitForCloseRoom(queue.body.patient.id).subscribe(data=>{
        this.status = 'DONE';
      });
    },
    error=>{
      if(error.status === 404) {
        this.infoMessage = "Não há nenhum paciente na fila"
      }
    })
  }

  registerAsPatient(): void {
    this.queueService.registerAsPatient()
    .subscribe(()=>{
      this.messageService.waitAndJoinRoom(this.currentUser.id).subscribe(message=>{
        this.status = 'IN_EXECUTION';
        this.roomId = message;
        this.messageService
      .getMessages()
      .subscribe((message: string) => {
        this.messageList.push(message);
      });
      });

      this.messageService.waitForCloseRoom(this.currentUser.id).subscribe(data=>{
        this.status = 'DONE';
      });
    });
;
  }

  sendMessage() {
    this.messageService.sendMessage(new Message(this.currentUser, this.newMessage, this.roomId));
    this.newMessage = '';
    this.status = 'DONE';
  }

  finish() {
    this.queueService.finishMedicalCare(this.queue.id).subscribe((data)=>{
      this.messageService.closeRoom(this.roomId);
    })
  }

}
