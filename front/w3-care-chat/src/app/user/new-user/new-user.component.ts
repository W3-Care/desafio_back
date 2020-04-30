import { Component, OnInit } from '@angular/core';
import { UserModel } from 'src/app/models/user.model';
import { UserService } from 'src/app/user.service';

@Component({
  selector: 'app-new-user',
  templateUrl: './new-user.component.html',
  styleUrls: ['./new-user.component.scss']
})
export class NewUserComponent implements OnInit {

  user: UserModel = new UserModel();
  loading = false;
  constructor(private service: UserService) { }

  ngOnInit(): void {
  }

  saveUser() {
    this.loading = true;
    this.service.save(this.user).subscribe(ret=>{
      this.loading = false;
    });
  }

}
