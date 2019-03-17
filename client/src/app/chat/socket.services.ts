import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';
import { EventEnum } from './shared/model/chat-enums';

const SERVER_URL = 'http://6d06fd5b.ngrok.io';

@Injectable({
    providedIn: 'root'
})
export class SocketService {
    private ws;
    private socket;
    private eventListener: any;

    constructor() {}

    public initSocket(): void {
        this.ws = new this.WsClient('ws://6d06fd5b.ngrok.io/ws');
    }

    WsClient(url) {
        this.ws = new WebSocket(url);
        this.eventListener = [];

        this.on = (event, cb) => {
            this.on(event, cb);
        };

        this.emit = (name, data) => {
            this.emit(name, data);
        };
    }

    on(event, cb) {
        return this.eventListener[event] = cb;
    }

    emit(name, data) {
        const event = {
            username: 'test',
            event: name, // 'message'
            data: data,
            date: Date.now(),
        };
        const input = JSON.stringify(event); // creating JSON object.
        this.ws.send(input);
    }

    public onMessage(): Observable<Event> {
        return new Observable<Event>(observer => {
            this.ws.on('event', (data: Event) => { console.log('data ', data); observer.next(data); });
        });
    }

/*    public send(event: Event): void {
        this.socket.emit('event', event);
    }

    public onMessage(): Observable<Event> {
        return new Observable<Event>(observer => {
            this.socket.on('event', (data: Event) => observer.next(data));
        });
    }

    public onEvent(eventenum: EventEnum): Observable<any> {
        return new Observable<EventEnum>(observer => {
            this.socket.on(eventenum, () => observer.next());
        });
    }*/
}
