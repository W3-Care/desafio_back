export class Message {
    id: string;
    body: string;
    roomId = 1;
    constructor(id,body) {
        this.id = id;
        this.body = body;
    }
}