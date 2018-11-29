import ξ from 'mithril';
import base from './base';
import {User, createUser} from '../models/user';


/**
 *  Onsubmit function sends form data
 *
 *  @param {User} user current user object to be registered
 */
function registerUser(user: User) {
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
                                User.validUsername() ? '' : 'error'}, [
                                ξ('label', 'Username'),
                                ξ('.ui.input',
                                    ξ('input', {
                                      type: 'text',
                                      value: User.getUsername(),
                                      oninput: (e: any) =>
                                        User.setUsername(
                                            e.currentTarget.value),
                                      placeholder: 'Thisisausername',
                                    }))]),
                              ξ('.field', {
                                class:
                                User.validFullname() ? '' : 'error'}, [
                                ξ('label', 'Full name'),
                                ξ('.ui.input',
                                    ξ('input', {
                                      type: 'text',
                                      value: User.getFullname(),
                                      oninput: (e: any) =>
                                        User.setFullname(
                                            e.currentTarget.value),
                                      placeholder: 'Mynameis',
                                    }))]),
                              ξ('.field', {
                                class:
                                User.validPassword() ? '' : 'error'}, [
                                ξ('label', 'Password'),
                                ξ('.ui.input',
                                    ξ('input', {
                                      type: 'password',
                                      value: User.getPassword(),
                                      oninput: (e: any) =>
                                        User.setPassword(
                                            e.currentTarget.value),
                                      placeholder: 's3cur3p#55w0rd',
                                    }))]),
                              ξ('.field', {
                                class:
                                User.validEmail() ? '' : 'error'}, [
                                ξ('label', 'Email'),
                                ξ('.ui.input',
                                    ξ('input', {
                                      type: 'text',
                                      value: User.getEmail(),
                                      oninput: (e: any) => User.setEmail(
                                          e.currentTarget.value),
                                      placeholder: 'Mynameis@example.com',
                                    }))]),
                              ξ('button.ui.teal.basic.button', {
                                disabled: !(User.validAll()),
                              },
                              'GET ROASTED!'),
                            ])),
                    ),
                ),
            )),
    );
  };
};
