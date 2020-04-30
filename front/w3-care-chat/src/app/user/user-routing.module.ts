import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { NewUserComponent } from './new-user/new-user.component';
import { AuthenticateComponent } from './authenticate/authenticate.component';

const RegisterRoutes: Routes = [
  {
    path: 'new',
    component: NewUserComponent
  },
  {
    path: 'auth',
    component: AuthenticateComponent
  },
];

@NgModule({
  imports: [RouterModule.forChild(RegisterRoutes)],
  exports: [RouterModule],
})
export class UserRoutingModule { }