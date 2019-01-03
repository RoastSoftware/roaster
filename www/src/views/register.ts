import ξ from 'mithril';
import base from './base';
import {UserModel} from '../models/user';
import Auth from '../services/auth';


export default class Register implements ξ.ClassComponent {
  registerError: Error;

  registerUser() {
    if (UserModel.validUsername()) {
      Auth.register<User>(
          Object.assign({}, UserModel), // TODO: Implement encode/decode funcs.
          UserModel.getPassword()) // TODO: Password isn't separated yet.
          .then((user) => {
            console.log(user);
            if (user) {
              Object.assign(UserModel, user);
              UserModel.setLoggedIn(true);
              ξ.route.set('/');
            }
          })
          .catch((err: Error) => {
            this.registerError = err;
            ξ.redraw();
          });
    }
  };

  view(vnode: ξ.CVnode) {
    return ξ(base,
        ξ('.ui.main.text.container[style=margin-top: 2em;]',
            ξ('.ui.grid',
                ξ('.ui.container.eight.wide.column.centered',
                    ξ('.ui.segments',
                        ξ('.ui.segment',
                            ξ('h2',
                                ξ('i.user.plus.icon'),
                                ' REGISTER')),

                        this.registerError ?
                        ξ('.ui.segment',
                            ξ('.ui.negative.message',
                                ξ('i.close.icon'),
                                ξ('.header',
                                    'Oh noeh!'),
                                ξ('p', this.registerError.message))): '',

                        ξ('.ui.segment.clearing',
                            ξ('form.ui.form', {onsubmit: () => {
                              this.registerUser();
                            }}, [
                              ξ('.field', {
                                class:
                                UserModel.validUsername() ? '' : 'error'}, [
                                ξ('label', 'Username'),
                                ξ('.ui.left.icon.input',
                                    ξ('input', {
                                      type: 'text',
                                      value: UserModel.getUsername(),
                                      oninput: (e: any) =>
                                        UserModel.setUsername(
                                            e.currentTarget.value),
                                      placeholder: 'Thisisausername',
                                    }),
                                    ξ('i.user.icon'),
                                )]),

                              ξ('.field', {
                                class:
                                UserModel.validFullname() ? '' : 'error'}, [
                                ξ('label', 'Full name'),
                                ξ('.ui.left.icon.input',
                                    ξ('input', {
                                      type: 'text',
                                      value: UserModel.getFullname(),
                                      oninput: (e: any) =>
                                        UserModel.setFullname(
                                            e.currentTarget.value),
                                      placeholder: 'Mynameis',
                                    }),
                                    ξ('i.info.icon'),
                                )]),
                              ξ('.field', {
                                class:
                                UserModel.validPassword() ? '' : 'error'}, [
                                ξ('label', 'Password'),
                                ξ('.ui.left.icon.input',
                                    ξ('input', {
                                      type: 'password',
                                      value: UserModel.getPassword(),
                                      oninput: (e: any) =>
                                        UserModel.setPassword(
                                            e.currentTarget.value),
                                      placeholder: 's3cur3p#55w0rd',
                                    }),
                                    ξ('i.lock.icon'),
                                )]),
                              ξ('.field', {
                                class:
                                UserModel.validEmail() ? '' : 'error'}, [
                                ξ('label', 'Email'),
                                ξ('.ui.left.icon.input',
                                    ξ('input', {
                                      type: 'text',
                                      value: UserModel.getEmail(),
                                      oninput: (e: any) => UserModel.setEmail(
                                          e.currentTarget.value),
                                      placeholder: 'Mynameis@example.com',
                                    }),
                                    ξ('i.envelope.icon'),
                                )]),
                              ξ('button.ui.teal.basic.button.right.floated', {
                                disabled: !(UserModel.validAll()),
                              },
                              ξ('i.user.plus.icon'), 'GET ROASTED!'),
                            ])),
                    ),
                ),
            )),
    );
  };
};
