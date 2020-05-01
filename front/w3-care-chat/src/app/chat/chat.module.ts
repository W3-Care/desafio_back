import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RoomComponent } from './room/room.component';
import { ChatRoutingModule } from './chat-routing.module';
import { FormsModule } from '@angular/forms';



@NgModule({
  declarations: [RoomComponent],
  imports: [
    CommonModule,
    ChatRoutingModule,
    FormsModule
  ]
})
export class ChatModule { }
