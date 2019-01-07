import ξ from 'mithril';
import base from './base';
import Network, {encodeURL} from '../services/network';
import {
  UserFeed,
  UserProfileHeader,
  RoastRatio,
  UserFolloweeList,
  UserFollowerList,
} from './user';
import {StatisticsFilter} from '../models/statistics';
import {UserModel} from '../models/user';

export default class Profile implements ξ.ClassComponent {
    uploadError: Error;

    profileImageURI: string = '';
    username: string = UserModel.getUsername();
    fullname: string = UserModel.getFullname();
    email: string = UserModel.getEmail();
    score: number = 0;

    tooLargeImageMessage() {
      return 'The image is too large, please choose a picture under 9MB';
    };

    updateProfileImageURI() {
      this.profileImageURI = encodeURL(
          'user', UserModel.getUsername(), 'avatar?' + new Date().getTime());
    }

    upload({target}) {
      this.uploadError = null;
      const avatar = target.files[0];
      if (avatar == undefined || avatar.length == 0) {
        return;
      }

      const data = new FormData();
      data.append('file', avatar);

      Network.request<FormData>('PUT',
          ['user', UserModel.getUsername(), 'avatar'], data)
          .then(() => {
            this.updateProfileImageURI();
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

    oninit() {
      this.updateProfileImageURI();
    };

    oncreate() {
      Network.request<Object>('GET', ['user', this.username, 'score'])
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
