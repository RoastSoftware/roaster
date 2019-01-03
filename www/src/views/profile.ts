import ξ from 'mithril';
import base from './base';
import Network from '../services/network';
import {
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
            console.error(this.uploadError);
            ξ.redraw();
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
          this.uploadError ?
          ξ('.ui.text.container', {
            style: 'margin-bottom: 1em;',
          },
            ξ('.ui.negative.message',
            ξ('.header',
            this.uploadError.message))): '',
          ξ('.ui.main.text.container.two.column.stackable.grid',
              ξ('.ui.column',
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
              ξ('.ui.column[minheight=10em]',
                  ξ(RoastRatio, {
                    filter: StatisticsFilter.User,
                    username: this.username,
                  }),
              ),
          ),
      );
    }
};
