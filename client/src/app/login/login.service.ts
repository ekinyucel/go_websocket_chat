import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from './user';

export interface Response {
    response: string;
}

@Injectable({
    providedIn: 'root'
})
export class LoginService {
    constructor(private httpClient: HttpClient) { }

    public login(user: User) {
        console.log('user ', user);
        const headers = {
            headers: new HttpHeaders({
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json',
                'Cache-Control': 'no-cache'
            })
        };

        return this.httpClient.post<User>(`http://localhost:8080/login`, user, headers)
            .subscribe(res => {
                console.log('res', res);
            }, err => {
                console.log('Error occured', err);
            });
    }
}
