import { UserModel } from './user.model';

export class Message {
    user: UserModel;
    body: string;
    roomId: number;
    constructor(user,body, roomId) {
        this.user = user;
        this.body = body;
        this.roomId = roomId;
    }
}