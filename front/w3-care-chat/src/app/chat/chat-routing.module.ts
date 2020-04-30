import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RoomComponent } from './room/room.component';

const RegisterRoutes: Routes = [
  {
    path: '',
    component: RoomComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(RegisterRoutes)],
  exports: [RouterModule],
})
export class ChatRoutingModule { }