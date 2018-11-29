import ξ from 'mithril';

export class User {
    static loggedIn: boolean = false;

    static username: string = '';
    private static usernameError: string = '';

    static fullname: string = '';
    private static fullnameError: string = '';

    static email: string = '';
    private static emailError: string = '';

    static isLoggedIn(): boolean {
      return User.loggedIn;
    };

    static setLoggedIn(state: boolean) {
      User.loggedIn = state;
    }

    static setUsername(user: string) {
      User.username = user;
      if (User.username.length < 1) {
        User.usernameError = 'Username must be longer than 0 characters';
      } else if (User.username.length > 30) {
        User.usernameError = 'Username must be shorter than 30 characters';
      } else {
        User.usernameError = '';
      }
    };

    static getUsername(): string {
      return User.username;
    };

    static validUsername(): boolean {
      return User.usernameError == '';
    };

    static errorUsername(): string {
      return User.usernameError;
    };

    static setFullname(user: string) {
      User.fullname = user;
      if (!(User.fullname.length < 255)) {
        User.fullnameError = 'Name must be shorter than 255 characters';
      } else {
        User.fullnameError = '';
      }
    };

    static getFullname(): string {
      return User.fullname;
    };

    static validFullname(): boolean {
      return User.fullnameError == '';
    };

    static errorFullname(): string {
      return User.fullnameError;
    };

    static setPassword(user: string) {
      User.password = user;
      if (User.password.length >= 4096) {
        User.passwordError = `Password must be shorter than 4096 characters,\
          that is more than 1 googol^98 in entropy`;
      } else if (User.password.length < 8) {
        User.passwordError = 'Password must be atleast 8 characters';
      } else {
        User.passwordError = '';
      }
    };

    static getPassword(): string {
      return User.password;
    };

    static setEmail(user: string) {
      /* eslint max-len: ["error", { "ignoreRegExpLiterals": true }] */
      const re = new RegExp(/^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i);

      User.email = user;
      if (!(User.email.length < 255)) {
        User.emailError = 'Email must be shorter than 255 characters';
      } else if (!(re.test(User.email))) {
        console.log('noo, es not valid');
        User.emailError = 'Must be an valid email';
      } else {
        User.emailError = '';
      }
    };

    static getEmail(): string {
      return User.email;
    };

    static validEmail(): boolean {
      return User.emailError == '';
    };

    static errorEmail(): string {
      return User.emailError;
    };

    static validAll(): boolean {
      return (User.validUsername() && User.validEmail() &&
            User.validFullname() && User.validPassword()) &&
            ((User.getUsername() !='') && (User.getEmail() != '') &&
                (User.getFullname() != '') && (User.getPassword() != ''));
    };

    static validLogin(): boolean {
      return (User.validUsername() && User.validPassword()) &&
            ((User.getUsername() != '') && (User.getPassword() != ''));
    };

    // TODO: Should we move this so it isn't inside the user model?
    static password: string = '';
    private static passwordError: string = '';

    static errorPassword(): string {
      return User.passwordError;
    };

    static validPassword(): boolean {
      return User.passwordError == '';
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
    url: '/user',
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
    url: '/session',
    data: user,
  });
};
/* eslint-disable */
function deAuthenticateUser() {
  return ξ.request({
    method: 'DELETE',
    url: '/session',
  });
};
/* eslint-enable */
