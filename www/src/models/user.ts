import m from 'mithril';

export interface User {
    email: string;
    userName: string;
    password: string;
    fullName: string;
}

const UserModel = {
  current: {} as User,
  load(userName: string) {
    return m.request<User>({
      method: 'GET',
      url: '/user/' + userName,

    })
        .then((result) => {
          UserModel.current = result;
        });
  },

  save() {
    return m.request({
      method: 'PATCH',
      url: '/user/' + userName,
      data: UserModel.current,
    });
  },
  // TODO: make sure request data is sent as JSON, applies for all commands.
  create() {
    return m.request({
method: 'POST',
      url: '/user/',
      data: UserModel.current,
    });
  },

  authenticateSession() {
    return m.request({
      method: 'POST',
      url: '/session/',
      data: UserModel.current,
    });
  },

  removeSession() {
    return m.request({
      method: 'DELETE',
      url: '/session/',
    });
  },

    /*
  userNameField: {
    value: stream(''),
    error: '',
    validate() {
      UserModel.userNameField.error =
                    UserModel.userNameField.value().length > 30 ?
                    'expected no more than 30 characters' : '';
    },
  },
  emailField: {
    value: stream(''),
    error: '',
    valildate() {
      UserModel.emailField.error =
                    UserModel.emailField.value().length > 255 ?
                    'expected a email address shorter than 255 characters' : '';
    },
  },
  passwordField: {
    value: stream(''),
    error: '',
    validate() {
      UserModel.passwordField.error =
                    UserModel.passwordField.value().length > 500 ?
                    `do you really need more than 1.3*10^(1204)\
                possible entropy? We're blocking this length of \
                passwords due to DOS attacks.` : '';
      UserModel.passwordField.error =
                    UserModel.passwordField.value().length < 8 ?
                    'expected a password of atleast 8 characters' : '';
    },
  },
  fullnameField: {
    value: stream(''),
    error: '',
    validate() {
      UserModel.fullnameField.error =
                    UserModel.fullnameField.value().length > 255 ?
                    'expected a name, no more than 255 characters' : '';
    },
  },
  validateFields() {
    Object.keys(UserModel).forEach((field) =>
      UserModel[field].valildate());
  },
     */
};


type UserModel = typeof UserModel;

export default UserModel;
