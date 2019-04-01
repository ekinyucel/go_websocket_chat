import { Injectable, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from './user';
import { Router } from '@angular/router';
import { map } from 'rxjs/internal/operators/map';

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

    // implement a proper authentication logic. this one is temporary
    get isAuthenticated() {
        return this.authenticated;
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
                    this.authenticated = true;
                    this.router.navigateByUrl('/');
                }
            }, err => {
                console.log('Error occured', err);
            });
    }
}
