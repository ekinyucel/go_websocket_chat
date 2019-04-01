import { CanActivate, Router } from '@angular/router';
import { Injectable } from '@angular/core';
import { LoginService } from '../login/login.service';

@Injectable({
    providedIn: 'root'
})
export class AuthGuard implements CanActivate {
    constructor(private router: Router,
                private loginService: LoginService) { }

    // TODO implement a proper authentication mechanism
    canActivate() {
        if (!this.loginService.isAuthenticated) {
            this.router.navigate(['/login']);
            return false;
        }
        return true;
    }
}
