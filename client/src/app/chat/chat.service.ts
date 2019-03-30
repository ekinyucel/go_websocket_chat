import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { map } from 'rxjs/operators';
import { WebSocketService } from './websocket.service';

export interface Message {
    user: string;
    data: string;
    date: string;
    type: string;
}

@Injectable({
    providedIn: 'root'
})
export class ChatService {
    public messages: Subject<Message>;

    constructor(_websocketService: WebSocketService) {
        this.messages = <Subject<Message>>_websocketService.connect('ws://localhost:8080/ws').pipe(map(
            (response: MessageEvent): Message => {
                const responseJSON = JSON.parse(response.data);
                return {
                    user: responseJSON.user,
                    data: responseJSON.data,
                    date: responseJSON.date,
                    type: responseJSON.type
                };
            }
        ));
     }
}
