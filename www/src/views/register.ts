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
    user: User = new User();

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
                                registerUser(this.user);
                              }}, [
                                ξ('.field', {
                                  class:
                                this.user.validUsername() ? '' : 'error'}, [
                                  ξ('label', 'Username'),
                                  ξ('.ui.input',
                                      ξ('input', {
                                        type: 'text',
                                        value: this.user.getUsername(),
                                        oninput: (e: any) =>
                                          this.user.setUsername(
                                              e.currentTarget.value),
                                        placeholder: 'Thisisausername',
                                      }))]),
                                ξ('.field', {
                                  class:
                                this.user.validFullname() ? '' : 'error'}, [
                                  ξ('label', 'Full name'),
                                  ξ('.ui.input',
                                      ξ('input', {
                                        type: 'text',
                                        value: this.user.getFullname(),
                                        oninput: (e: any) =>
                                          this.user.setFullname(
                                              e.currentTarget.value),
                                        placeholder: 'Mynameis',
                                      }))]),
                                ξ('.field', {
                                  class:
                                this.user.validPassword() ? '' : 'error'}, [
                                  ξ('label', 'Password'),
                                  ξ('.ui.input',
                                      ξ('input', {
                                        type: 'password',
                                        value: this.user.getPassword(),
                                        oninput: (e: any) =>
                                          this.user.setPassword(
                                              e.currentTarget.value),
                                        placeholder: 's3cur3p#55w0rd',
                                      }))]),
                                ξ('.field', {
                                  class:
                                this.user.validEmail() ? '' : 'error'}, [
                                  ξ('label', 'Email'),
                                  ξ('.ui.input',
                                      ξ('input', {
                                        type: 'text',
                                        value: this.user.getEmail(),
                                        oninput: (e: any) => this.user.setEmail(
                                            e.currentTarget.value),
                                        placeholder: 'Mynameis@example.com',
                                      }))]),
                                ξ('button.ui.teal.basic.button', {
                                  disabled: !(this.user.validAll()),
                                },
                                'GET ROASTED!'),
                              ])),
                      ),
                  ),
              )),
      );
    };
};
