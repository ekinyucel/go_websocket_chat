import { Injectable, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from './user';
import { Router } from '@angular/router';

export interface Response {
    message: string;
    statuscode: number;
}

@Injectable({
    providedIn: 'root'
})
export class LoginService implements OnInit {
    private authenticated = false;

    constructor(private httpClient: HttpClient,
        private router: Router) { }

    ngOnInit() {
        this.authenticated = false;
    }

    public login(user: User) {
        const headers = {
            headers: new HttpHeaders({
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json',
                'Cache-Control': 'no-cache'
            })
        };

        return this.httpClient.post<Response>(`http://localhost:8080/login`, user, headers)
            .subscribe(res => {
                if (res.statuscode === 200) {
                    this.router.navigateByUrl('/');
                    this.authenticated = true;
                }
            }, err => {
                console.log('Error occured', err);
            });
    }
}
