import ξ from 'mithril';
import base from './base';
import {UserModel} from '../models/user';
import Network from '../services/network';
import {RoastDoughnutStatisticsModel} from '../models/statistics';
import Chart from 'chart.js';

export default class Profile implements ξ.ClassComponent {
    uploadError: APIError;
    downloadError: APIError;
    profileImageURI: string = `/user/${UserModel.getUsername()}/avatar?` +
      new Date().getTime();

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
          .catch((err: APIError) => {
            this.uploadError = err;
            console.log(this.uploadError);
          });
    };

    clickImg() {
      document.getElementById('upload').click();
    };

    getUserStat() {
      console.log('getting user statistics');
      // TODO: Network.request statistics
    }

    view(vnode: ξ.CVnode) {
      return ξ(base,
          ξ('.ui.text.container', {
            style: 'margin-top: 1em;',
          },
          ξ('h1.ui.header',
              ξ('img.ui.circular.image', {
                src: this.profileImageURI,
              }),
              ξ('.content',
                  'MY PROFILE',
                  ξ('.sub.header', `Hello there, ${UserModel.getFullname()}!`)),
          ),
          ξ('.ui.divider')),
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
                    onclick: this.clickImg,
                    style: 'cursor: pointer;',
                  },
                  'User profile picture.'),
                  ξ('h2',
                      UserModel.getFullname()),
                  ξ('p',
                      ξ('i.user.icon'),
                      UserModel.getUsername()),
                  ξ('p',
                      ξ('i.mail.icon'),
                      UserModel.getEmail()),
              ),
              ξ('.ui.column[minheight=10em]',
                  ξ('canvas#chart-area', {
                    oncreate: ({dom}) => {
                      const ctx = (document.getElementById(
                          'chart-area') as HTMLCanvasElement)
                          .getContext('2d');
                      new Chart(ctx, {
                        type: 'doughnut',
                        data: RoastDoughnutStatisticsModel.dataDonut,
                        options: RoastDoughnutStatisticsModel.optionsDonut,
                      });
                    }})
              )
          )
      );
    }
};
