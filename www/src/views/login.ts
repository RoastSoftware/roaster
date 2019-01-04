import ξ from 'mithril';
import base from './base';
import {UserModel} from '../models/user';
import Auth from '../services/auth';

export default class Login implements ξ.ClassComponent {
  loginError: Error;

  authenticate(): Promise<User> {
    if (UserModel.validLogin()) {
      return Auth.login<User>(UserModel.getUsername(), UserModel.getPassword())
          .then((user) => {
            UserModel.setLoggedIn(true);
            Object.assign(UserModel, user);
            ξ.route.set('/');

            return user;
          })
          .catch((err: Error) => {
            this.loginError = err;
            console.log(this);
            ξ.redraw();
          });
    }
  };

  view() {
    return ξ(base, ξ('.ui.main.text.container[style=margin-top: 2em;]',
        ξ('.ui.grid',
            ξ('.ui.container.eight.wide.column.centered',
                ξ('.ui.segments',
                    ξ('.ui.segment',
                        ξ('h2',
                            ξ('i.sign.in.icon'),
                            ' LOGIN')),

                    (this.loginError == null ? '':
                      ξ('.ui.segment',
                          ξ('.ui.negative.message',
                              ξ('.header',
                                  'Oh noeh!'),
                              ξ('p', this.loginError.message)))),

                    ξ('.ui.segment.clearing', ξ('form.ui.form', {
                      onsubmit: () => {
                        this.authenticate();
                      }},
                    [
                      ξ('.field', {
                        class: UserModel.validUsername() ? '' : 'error'}, [
                        ξ('label', 'Username'),
                        ξ('.ui.left.icon.input',
                            ξ('input', {
                              type: 'text',
                              value: UserModel.getUsername(),
                              oninput: (e: any) =>
                                UserModel.setUsername(e.currentTarget.value),
                              placeholder: 'Thisisausername',
                            }),
                            ξ('i.user.icon'),
                        )]),

                      ξ('.field', {
                        class: UserModel.validPassword() ? '' : 'error'}, [
                        ξ('label', 'Password'),
                        ξ('.ui.left.icon.input',
                            ξ('input', {
                              type: 'password',
                              value: UserModel.getPassword(),
                              oninput: (e: any) =>
                                UserModel.setPassword(e.currentTarget.value),
                              placeholder: 's3cur3p#55w0rd',
                            }),
                            ξ('i.lock.icon'),
                        )]),

                      ξ('button.ui.teal.basic.button.right.floated', {
                        disabled: !(UserModel.validLogin()),
                      },
                      ξ('i.sign.in.icon'), 'LOGIN!'),
                    ]))
                ),
            ),
        )));
  };
};
