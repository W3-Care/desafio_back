import { Injectable } from '@angular/core';
import { UserService } from './user.service';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { map } from 'rxjs/operators';
import { MessageService } from './services/message.service';

@Injectable({
  providedIn: 'root'
})
export class QueueService {
  currentUser;
  constructor(
    private userService: UserService, 
    private http: HttpClient,
    private messageService: MessageService

    ) {}

  ngOnInit(): void {
  }
  
  pullNewRoom() {
    this.currentUser = this.userService.currentUserValue;
    return this.http.post<any>(`${environment.api}/api/queues/pull`, {doctorId: this.currentUser.id}, {observe: 'response' as 'body'});
  }
  
  registerAsPatient() {
    this.currentUser = this.userService.currentUserValue;
    return this.http.post<any>(`${environment.api}/api/queues/`, {patientId: this.currentUser.id}, {observe: 'response' as 'body'});
  }

  finishMedicalCare(queueId) {
    this.currentUser = this.userService.currentUserValue;
    return this.http.post<any>(`${environment.api}/api/queues/finish`, {id: queueId, patientId: this.currentUser.id}, {observe: 'response' as 'body'});
  }
}
