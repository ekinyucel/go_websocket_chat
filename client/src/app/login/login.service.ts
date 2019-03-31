import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from './user';

@Injectable({
    providedIn: 'root'
})
export class LoginService {
    constructor(private httpClient: HttpClient) { }

    public login(user: User) {
        const headers = {
            headers: new HttpHeaders({
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json',
                'Cache-Control': 'no-cache'
            })
        };

        return this.httpClient.post(`http://localhost:8080/login`, user, headers)
            .subscribe(res => {
                console.log('res', res);
            }, err => {
                console.log('Error occured', err);
            });
    }
}
