import { Component, OnInit } from '@angular/core';
import { SocketService } from './socket.services';
import { EventEnum, Action } from './shared/model/chat-enums';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.css']
})
export class ChatComponent implements OnInit {
  action = Action;
  events: Event[] = [];
  connection: any;

  constructor(private _socketService: SocketService) { }

  ngOnInit(): void {
    this.initConnection();
  }

  private initConnection(): void {
    this._socketService.initSocket();

    this.connection = this._socketService.onMessage()
      .subscribe((event: Event) => {
        console.log('data ', event);
      });

    /*this.connection = this._socketService.onMessage()
      .subscribe((event: Event) => {
        this.events.push(event);
      });

    this._socketService.onEvent(EventEnum.CONNECT)
      .subscribe(() => {
        console.log('connected');
      });

    this._socketService.onEvent(EventEnum.DISCONNECT)
      .subscribe(() => {
        console.log('disconnected');
      });*/
  }

  public sendMessage(message: string): void {
    console.log('sendMessage');
    if (!message) {
      return;
    }
    console.log('sendMessage');
  }

}
