import { Component } from '@angular/core';
import { WebSocketService } from './websocket.service';
import { ChatService } from './chat.service';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { DatePipe } from '@angular/common';

export interface Message {
  user: string;
  data: string;
  room: string;
  date: string;
  type: string;
}

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.css'],
  providers: [ WebSocketService, ChatService, DatePipe ]
})
export class ChatComponent {
  formGroup: FormGroup;
  messages: Message[] = [];
  totalClients = 0;
  clients: string[] = [];

  constructor(private _chatService: ChatService,
              fb: FormBuilder,
              private datePipe: DatePipe) {
    _chatService.messages.subscribe(msg => {
      msg = this.detectMessageType(msg);

      this.messages.push(msg);
    });

    this.formGroup = fb.group({
      message: fb.control('message', Validators.required)
    });
  }

  public detectMessageType(msg: Message): Message {
    switch (msg.type) {
      case 'connect': {
        this.totalClients = +msg.data; // string to int conversion
        msg.data = 'A new client has connected';
        this.clients.push(msg.user);
        break;
      }
      case 'disconnect': {
        this.totalClients = +msg.data;
        msg.data = 'A client has disconnected';
        break;
      }
    }
    return msg;
  }

  public sendMessage(): void {
    const message = {
      user: this.formGroup.get('username').value,
      data: this.formGroup.get('message').value,
      room: 'general', // hardcoded for now
      date: this.formatDate(new Date()),
      type: 'message' // TODO make it constant
    };
    this._chatService.messages.next(message);
  }

  public formatDate(date: Date) {
    return this.datePipe.transform(date, 'yyyy/MM/dd hh:mm:ss');
  }

}
