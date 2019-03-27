import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { map } from 'rxjs/operators';
import { WebSocketService } from './websocket.service';

export interface Message {
    user: string;
    message: string;
}

@Injectable({
    providedIn: 'root'
})
export class ChatService {
    public messages: Subject<Message>;

    constructor(_websocketService: WebSocketService) {
        this.messages = <Subject<Message>>_websocketService.connect('ws://0016597b.ngrok.io/ws').pipe(map(
            (response: MessageEvent): Message => {
                const data = JSON.parse(response.data);
                return {
                    user: data.user,
                    message: data.message
                };
            }
        ));
     }
}
