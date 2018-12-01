import ξ from 'mithril';
import base from './base';
import {UserModel, createUser} from '../models/user';


/**
 *  Onsubmit function sends form data
 *
 *  @param {User} user current user object to be registered
 */
function registerUser(user: UserModel) {
  console.log('On submit function called, but not implemented');
  if (user.validUsername()) {
    createUser(user);
  }
}

export default class Register implements ξ.ClassComponent {
  view(vnode: ξ.CVnode) {
    return ξ(base,
        ξ('.ui.main.text.container[style=margin-top: 2em;]',
            ξ('.ui.grid',
                ξ('.ui.container.six.wide.column.centered',
                    ξ('.ui.segments',
                        ξ('.ui.segment',
                            ξ('h2', 'REGISTER')),
                        ξ('.ui.segment',
                            ξ('form.ui.form', {onsubmit: () => {
                              registerUser(User);
                            }}, [
                              ξ('.field', {
                                class:
                                UserModel.validUsername() ? '' : 'error'}, [
                                ξ('label', 'Username'),
                                ξ('.ui.input',
                                    ξ('input', {
                                      type: 'text',
                                      value: UserModel.getUsername(),
                                      oninput: (e: any) =>
                                        UserModel.setUsername(
                                            e.currentTarget.value),
                                      placeholder: 'Thisisausername',
                                    }))]),
                              ξ('.field', {
                                class:
                                UserModel.validFullname() ? '' : 'error'}, [
                                ξ('label', 'Full name'),
                                ξ('.ui.input',
                                    ξ('input', {
                                      type: 'text',
                                      value: UserModel.getFullname(),
                                      oninput: (e: any) =>
                                        UserModel.setFullname(
                                            e.currentTarget.value),
                                      placeholder: 'Mynameis',
                                    }))]),
                              ξ('.field', {
                                class:
                                UserModel.validPassword() ? '' : 'error'}, [
                                ξ('label', 'Password'),
                                ξ('.ui.input',
                                    ξ('input', {
                                      type: 'password',
                                      value: UserModel.getPassword(),
                                      oninput: (e: any) =>
                                        UserModel.setPassword(
                                            e.currentTarget.value),
                                      placeholder: 's3cur3p#55w0rd',
                                    }))]),
                              ξ('.field', {
                                class:
                                UserModel.validEmail() ? '' : 'error'}, [
                                ξ('label', 'Email'),
                                ξ('.ui.input',
                                    ξ('input', {
                                      type: 'text',
                                      value: UserModel.getEmail(),
                                      oninput: (e: any) => UserModel.setEmail(
                                          e.currentTarget.value),
                                      placeholder: 'Mynameis@example.com',
                                    }))]),
                              ξ('button.ui.teal.basic.button', {
                                disabled: !(UserModel.validAll()),
                              },
                              'GET ROASTED!'),
                            ])),
                    ),
                ),
            )),
    );
  };
};
