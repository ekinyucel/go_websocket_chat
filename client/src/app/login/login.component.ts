import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { LoginService } from './login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  formGroup: FormGroup;

  constructor(private _loginService: LoginService,
              fb: FormBuilder) {
    this.formGroup = fb.group({
      username: fb.control('username', Validators.required),
      password: fb.control('password', Validators.required)
    });
  }

  ngOnInit() {
  }

  public login(): void {
    const user = {
      username: this.formGroup.get('username').value,
      password: this.formGroup.get('password').value,
    };
    console.log('user ', user);
    this._loginService.login(user);
  }

}
