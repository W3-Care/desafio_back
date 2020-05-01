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
  successMessage="";
  errorMessage="";
  constructor(private service: UserService) { }

  ngOnInit(): void {
  }

  saveUser() {
    this.loading = true;
    this.service.save(this.user).subscribe(ret=>{
      this.loading = false;
      this.successMessage="Usuário cadastrado com sucesso!"
      this.errorMessage = "";
      this.user = new UserModel();
    },
    (error)=> {
      this.successMessage=""
      this.errorMessage = "Ocorreu um Erro. Verifique os dados e tente novamente.";
    })
    ;
  }

}
