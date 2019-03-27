import { Component, OnInit } from '@angular/core';
import { WebSocketService } from './websocket.service';
import { EventEnum, Action } from './shared/model/chat-enums';
import { ChatService } from './chat.service';
import { FormGroup, FormControl, FormBuilder, Validators } from '@angular/forms';
import { Subject } from 'rxjs';

export interface Message {
  user: string;
  message: string;
}

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.css'],
  providers: [ WebSocketService, ChatService ]
})
export class ChatComponent {
  formGroup: FormGroup;
  messages: Message[] = [];

  constructor(private _chatService: ChatService, fb: FormBuilder) {
    console.log(this.messages.length);
    _chatService.messages.subscribe(msg => {
      this.messages.push(msg);
      console.log('response from websocket: ' + msg.message);
    });

    this.formGroup = fb.group({
      username: fb.control('username', Validators.required),
      message: fb.control('message', Validators.required)
    });
  }

  public sendMessage(): void {
    const message = {
      user: this.formGroup.get('username').value,
      message: this.formGroup.get('message').value
    };
    console.log('new message from client to websocket: ', message);
    this._chatService.messages.next(message);
  }

}
