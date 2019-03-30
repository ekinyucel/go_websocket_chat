import { Component } from '@angular/core';
import { WebSocketService } from './websocket.service';
import { ChatService } from './chat.service';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { DatePipe } from '@angular/common';

export interface Message {
  user: string;
  data: string;
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
  clients = 0;

  constructor(private _chatService: ChatService,
              fb: FormBuilder,
              private datePipe: DatePipe) {
    _chatService.messages.subscribe(msg => {
      switch (msg.type) {
        case 'connect': {
          this.clients = +msg.data; // string to int conversion
          msg.data = 'A new client has connected';
          break;
        }
        case 'disconnect': {
          this.clients = +msg.data;
          msg.data = 'A client has disconnected';
          break;
        }
      }
      this.messages.push(msg);

      console.log('response from websocket: ' + msg.data);
    });

    this.formGroup = fb.group({
      username: fb.control('username', Validators.required),
      message: fb.control('message', Validators.required)
    });
  }

  public sendMessage(): void {
    const message = {
      user: this.formGroup.get('username').value,
      data: this.formGroup.get('message').value,
      date: this.formatDate(new Date()),
      type: 'message' // TODO make it constant
    };
    console.log('new message from client to websocket: ', message);
    this._chatService.messages.next(message);
  }

  public formatDate(date: Date) {
    return this.datePipe.transform(date, 'yyyy/MM/dd hh:mm:ss');
  }

}
