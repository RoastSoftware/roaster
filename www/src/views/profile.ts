import ξ from 'mithril';
import base from './base';
import Network from '../services/network';
import {
  UserFeed,
  UserProfileHeader,
  RoastRatio,
} from './user';
import {StatisticsFilter} from '../models/statistics';
import {UserModel} from '../models/user';

export default class Profile implements ξ.ClassComponent {
    uploadError: Error;
    downloadError: Error;
    profileImageURI: string = `/user/${UserModel.getUsername()}/avatar?` +
      new Date().getTime();
    username: string = UserModel.getUsername();
    fullname: string = UserModel.getFullname();
    email: string = UserModel.getEmail();
    score: number = 0;

    upload({target}) {
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
            console.error(this.uploadError);
          });
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
            score: this.score,
            loggedIn: true,
          }),
          ξ('.ui.divider')),
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
                        style: 'cursor: pointer;',
                      },
                      'User profile picture.'),
                      ξ('h2',
                          this.fullname),
                      ξ('p',
                          ξ('i.user.icon'),
                          this.username),
                      ξ('p',
                          ξ('i.mail.icon'),
                          this.email),
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
          ),
      );
    }
};
