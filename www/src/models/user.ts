import ξ from 'mithril';

export class User {
    username: string = '';
    private usernameError: string = '';
    fullname: string = '';
    private fullnameError: string = '';
    password: string = '';
    private passwordError: string = '';
    email: string = '';
    private emailError: string = '';

    setUsername(user: string) {
      this.username = user;
      if (this.username.length < 1) {
        this.usernameError = 'Username must be longer than 0 characters';
      } else if (this.username.length > 30) {
        this.usernameError = 'Username must be shorter than 30 characters';
      } else {
        this.usernameError = '';
      }
    };

    getUsername(): string {
      return this.username;
    };

    validUsername(): boolean {
      return this.usernameError == '';
    };

    errorUsername(): string {
      return this.usernameError;
    };

    setFullname(user: string) {
      this.fullname = user;
      if (!(this.fullname.length < 255)) {
        this.fullnameError = 'Name must be shorter than 255 characters';
      } else {
        this.fullnameError = '';
      }
    };

    getFullname(): string {
      return this.fullname;
    };

    validFullname(): boolean {
      return this.fullnameError == '';
    };

    errorFullname(): string {
      return this.fullnameError;
    };

    setPassword(user: string) {
      this.password = user;
      if (this.password.length >= 4096) {
        this.passwordError = `Password must be shorter than 4096 characters,\
          that is more than 1 googol^98 in entropy`;
      } else if (this.password.length < 8) {
        this.passwordError = 'Password must be atleast 8 characters';
      } else {
        this.passwordError = '';
      }
    };

    getPassword(): string {
      return this.password;
    };

    validPassword(): boolean {
      return this.passwordError == '';
    };

    errorPassword(): string {
      return this.passwordError;
    };

    setEmail(user: string) {
      /* eslint max-len: ["error", { "ignoreRegExpLiterals": true }] */
      const re = new RegExp(/^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i);

      this.email = user;
      if (!(this.email.length < 255)) {
        this.emailError = 'Email must be shorter than 255 characters';
      } else if (!(re.test(this.email))) {
        console.log('noo, es not valid');
        this.emailError = 'Must be an valid email';
      } else {
        this.emailError = '';
      }
    };

    getEmail(): string {
      return this.email;
    };

    validEmail(): boolean {
      return this.emailError == '';
    };

    errorEmail(): string {
      return this.emailError;
    };

    validAll(): boolean {
      return (this.validUsername() && this.validEmail() &&
            this.validFullname() && this.validPassword()) &&
            ((this.getUsername() !='') && (this.getEmail() != '') &&
                (this.getFullname() != '') && (this.getPassword() != ''));
    };

    validLogin(): boolean {
      return (this.validUsername() && this.validPassword()) &&
            ((this.getUsername() != '') && (this.getPassword() != ''));
    };
};
/* eslint-disable */
function retrieveUser(username: string): Promise<User> {
  return ξ.request<User>({
    method: 'GET',
    url: '/user/' + this.username,
  }).then((result) => {
    const user: User = result;
    return user;
  });
};
/* eslint-enable */
export function createUser(user: User) {
  return ξ.request({
    method: 'POST',
    url: '/user/',
    data: user,
  });
};
/* eslint-disable */
function saveUser() {
  return ξ.request({
    method: 'PATCH',
    url: '/user/' + this.username,
    data: this,
  });
};
/* eslint-enable */
export function authenticateUser(user: User) {
  return ξ.request({
    method: 'POST',
    url: '/session/',
    data: user,
  });
};
/* eslint-disable */
function deAuthenticateUser() {
  return ξ.request({
    method: 'DELETE',
    url: '/session/',
  });
};
/* eslint-enable */
