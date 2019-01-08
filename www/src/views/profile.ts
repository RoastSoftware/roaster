import ξ from 'mithril';
import base from './base';
import Network from '../services/network';
import {
  UserFeed,
  UserProfileHeader,
  RoastRatio,
  UserFolloweeList,
  UserFollowerList,
} from './user';
import {StatisticsFilter} from '../models/statistics';
import {UserModel} from '../models/user';

class FullnameField implements ξ.ClassComponent {
    updateError: Error;
    visible: boolean = true;
    
    updateFullName(){
        this.visible = !this.visible;
    };

    submitNewFullname(){
        Network.request<any>('PATCH', '/user/' + UserModel.getUsername(), {"fullname": UserModel.getFullname()})
            .then(() => {
                this.updateFullName();
            })
            .catch((err: Error) => { 
            this.updateError = err;
            console.error(this.updateError);
            });
    }

    view({attrs}) {
        return [
            this.visible ?
            ξ('p', {
                onclick: () => this.updateFullName(),
                style: 'cursor: pointer;',
            },
            ξ('i.user.icon'),
                attrs.fullname)
            :
            [
                ξ('.ui.icon.input.fluid',{
                        class: UserModel.validFullname() ? '' : 'error',
                    },
                    [
                    ξ('input', {
                        type: 'text',
                        value: UserModel.getFullname(),
                        oninput: (e: any) =>
                        UserModel.setFullname(e.currentTarget.value),
                        placeholder: 'YourNewFullname',
                    }),
                ξ('button.ui.button', {
                    onclick: () => this.submitNewFullname(),
                    disabled: !(UserModel.validFullname()),
                },'Save'),
                ]),
          this.updateError ?
          ξ('.ui.negative.message',
              ξ('.header',
                  this.updateError.message)): '',
            ξ.redraw(),
            ]
        ]
    }
}

class EmailField implements ξ.ClassComponent {
    updateError: Error;
    visible: boolean = true;
    
    updateEmail(){
        this.visible = !this.visible;
    };

    submitNewEmail(){
        Network.request<any>('PATCH', '/user/' + UserModel.getUsername(), {"email": UserModel.getEmail()})
            .then(() => {
                this.updateEmail();
            })
            .catch((err: Error) => { 
            this.updateError = err;
            console.error(this.updateError);
            });
    }

    view({attrs}) {
        return [
            this.visible ?
            ξ('p', {
                onclick: () => this.updateEmail(),
                style: 'cursor: pointer;',
            },
            ξ('i.mail.icon'),
                attrs.email)
            :
            [
                ξ('.ui.icon.input.fluid',{
                    class: UserModel.validEmail() ? '' : 'error',
                    },
                    [
                    ξ('input', {
                        type: 'text',
                        value: UserModel.getEmail(),
                        oninput: (e: any) =>
                        UserModel.setEmail(e.currentTarget.value),
                        placeholder: 'YourNewEmail',
                    }),
                ξ('button.ui.button', {
                    onclick: () => this.submitNewEmail(),
                    disabled: !(UserModel.validEmail()),
                },'Save'),
                ]),
          this.updateError ?
          ξ('.ui.negative.message',
              ξ('.header',
                  this.updateError.message)): '',
            ξ.redraw(),
            ]
        ]
    }
}

export default class Profile implements ξ.ClassComponent {
    uploadError: Error;

    tooLargeImageMessage() {
      return 'The image is too large, please choose a picture under 9MB';
    };

    profileImageURI: string = `/user/${UserModel.getUsername()}/avatar?` +
      new Date().getTime();
    username: string = UserModel.getUsername();
    fullname: string = UserModel.getFullname();
    email: string = UserModel.getEmail();
    score: number = 0;

    upload({target}) {
      this.uploadError = null;
      const avatar = target.files[0];
      if (avatar == undefined || avatar.length == 0) {
        return;
      }

      const data = new FormData();
      data.append('file', avatar);

      Network.request<FormData>('PUT', '/user/' +
          UserModel.getUsername() + '/avatar', data)
          .then(() => {
            this.profileImageURI = `/user/${UserModel.getUsername()}/avatar?`
              + new Date().getTime();
          })
          .catch((err: Error) => {
            this.uploadError = err;
            if ('code' in err && err.code == 413) {
              this.uploadError.message = this.tooLargeImageMessage();
            }
            console.error(this.uploadError);
            ξ.redraw();
          });
    };

    onupdate(){
        this.fullname = UserModel.getFullname();
        this.email = UserModel.getEmail();
    };

    oncreate() {
      Network.request<Object>('GET', '/user/' + this.username + '/score')
          .then(({score}) => {
            this.score = score;
          });
    };

    clickImage() {
      document.getElementById('upload').click();
    };

    view({attrs}) {
      return ξ(base,
          ξ('.ui.text.container', {
            style: 'margin-top: 1em;',
          },
          ξ(UserProfileHeader, {
            username: this.username,
            fullname: this.fullname,
            avatar: this.profileImageURI,
            score: this.score,
            loggedIn: true,
          }),
          ξ('.ui.divider')),
          this.uploadError ?
          ξ('.ui.text.container', {
            style: 'margin-bottom: 1em;',
          },
          ξ('.ui.negative.message',
              ξ('.header',
                  this.uploadError.message))): '',
          ξ('.ui.main.text.stackable.two.column.grid.container',
              ξ('.ui.row',
                  ξ('.column',
                      ξ('input#upload', {
                        onchange: (e) => {
                          this.upload(e);
                        },
                        type: 'File',
                        style: 'display: none;',
                        accept: '.png, .jpg, .jpeg;',
                      }),
                      ξ('img.ui.image.rounded.medium#picture', {
                        src: this.profileImageURI,
                        onclick: this.clickImage,
                        style: 'cursor: pointer; width: 100%;',
                      },
                      'User profile picture.'),
                      ξ('h2',
                          this.username),
                      ξ(FullnameField, {
                          fullname: this.fullname,
                      }),
                      ξ(EmailField, {
                            email: this.email, 
                    }),
                  ),
                  ξ('.column[minheight=10em]',
                      ξ(RoastRatio, {
                        filter: StatisticsFilter.User,
                        username: this.username,
                      }),
                  ),
              ),
              ξ('.ui.one.column.row',
                  ξ('.column',
                      ξ('.ui.basic.segment',
                          ξ('h2.ui.dividing.header',
                              ξ('i.feed.icon'),
                              ξ('.content', 'YOUR FEED',
                                  ξ('.sub.header',
                                      `What have you been up to lately?`)),
                          ),
                          ξ(UserFeed, {
                            username: this.username,
                          }),
                      ),
                      ξ('.ui.basic.center.aligned.segment',
                          'Check out the ', ξ('a', {
                            href: '/feed',
                            oncreate: ξ.route.link,
                          }, 'feed page'), ' for more!',
                      ),
                  ),
              ),
              ξ('.ui.two.column.row',
                  ξ('.column',
                      ξ('.ui.basic.segment',
                          ξ('h2.ui.dividing.header',
                              ξ('i.users.icon'),
                              ξ('.content', 'FOLLOWING',
                                  ξ('.sub.header', `${this.username} 
                                  finds these people very intriguing.`)),
                          ),
                          ξ('.ui.feed',
                              ξ(UserFolloweeList, {
                                username: this.username,
                              }),
                          )
                      ),
                  ),
                  ξ('.column',
                      ξ('.ui.basic.segment',
                          ξ('h2.ui.dividing.header',
                              ξ('i.users.icon'),
                              ξ('.content', 'FOLLOWERS',
                                  ξ('.sub.header', `${this.username} 
                                  is followed by EVERYONE! No, but
                                      by these people.`)),
                          ),
                          ξ('.ui.feed',
                              ξ(UserFollowerList, {
                                username: this.username,
                              }),
                          )
                      ),
                  ),
              ),
          ),
      );
    }
};
