import ξ from 'mithril';
import base from './base';
import {User} from '../models/user';
import Auth from '../services/auth';

export default class Login implements ξ.ClassComponent {
  loginError: APIError;

  authenticate(): Promise {
    if (User.validLogin()) {
      return Auth.login(User.getUsername(), User.getPassword())
          .then(() => {
            User.setLoggedIn(true);
            ξ.route.set('/');
          })
          .catch((err: APIError) => {
            this.loginError = err;
          });
    }
  };

  view(controller) {
    return ξ(base, ξ('.ui.main.text.container[style=margin-top: 2em;]',
        ξ('.ui.grid',
            ξ('.ui.container.six.wide.column.centered',
                ξ('.ui.segments',
                    ξ('.ui.segment', ξ('h2', 'LOGIN')),

                    (this.errorMessage == null ? '':
                      ξ('.ui.segment',
                          ξ('.ui.negative.message',
                              ξ('i.close.icon'),
                              ξ('.header',
                                  'Oh noeh!'),
                              ξ('p', this.errorMessage)))),

                    ξ('.ui.segment', ξ('form.ui.form', {
                      onsubmit: this.authenticate},
                    [
                      ξ('.field', {
                        class: User.validUsername() ? '' : 'error'}, [
                        ξ('label', 'Username'),
                        ξ('.ui.input',
                            ξ('input', {
                              type: 'text',
                              value: User.getUsername(),
                              oninput: (e: any) =>
                                User.setUsername(e.currentTarget.value),
                              placeholder: 'Thisisausername',
                            }))]),

                      ξ('.field', {
                        class: User.validPassword() ? '' : 'error'}, [
                        ξ('label', 'Password'),
                        ξ('.ui.input',
                            ξ('input', {
                              type: 'password',
                              value: User.getPassword(),
                              oninput: (e: any) =>
                                User.setPassword(e.currentTarget.value),
                              placeholder: 's3cur3p#55w0rd',
                            }))]),

                      ξ('button.ui.teal.basic.button', {
                        disabled: !(User.validLogin()),
                      }, 'LOGIN!'),
                    ]))
                ),
            ),
        )));
  };
};
