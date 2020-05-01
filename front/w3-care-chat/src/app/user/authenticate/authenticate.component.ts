import { Component, OnInit } from '@angular/core';
import { UserService } from 'src/app/user.service';
import { first } from 'rxjs/operators';
import { Router } from '@angular/router';

@Component({
  selector: 'app-authenticate',
  templateUrl: './authenticate.component.html',
  styleUrls: ['./authenticate.component.scss']
})
export class AuthenticateComponent implements OnInit {
  loading= false;
  constructor(private service: UserService, private router: Router) { }
  email: string = 'bernardo@patient.com';
  password: string = '123456';
  error = '';

  ngOnInit(): void {
  }

  auth() {
    this.loading = true;
    this.service.auth(this.email, this.password)
    .subscribe(
        data => {
          this.loading=false;
          this.router.navigate(['/chat']);
        },
        error => {
            this.error = error;
            this.loading = false;
        });
  }

}
