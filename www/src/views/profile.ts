import ξ from 'mithril';
import base from './base';
import {UserModel} from '../models/user';
import Network from '../services/network';
import Model from '../models/statistics';
import Chart from 'chart.js';

export default class Profile implements ξ.ClassComponent {
    uploadError: Error;

    upload(e: Any) {
        this.uploadError = null;
      const avatar = e.target.files[0];
      if (avatar == undefined || avatar.length == 0) {
        return;
      }
      console.log('SUCCESS');
      const data = new FormData();
      data.append('file', avatar);
      Network.request<FormData>('PUT', '/user/' +
          UserModel.getUsername() + '/avatar', data)
          .then((_: any) => {
            const img = document
                .getElementById('picture') as HTMLImageElement;
            img.src = '/user/' + UserModel.getUsername() +
                    '/avatar?' + new Date().getTime();
          })
          .catch((err: Error) => {
            this.uploadError = err;
            console.error(this.uploadError);
            ξ.redraw();
          });
    };

    clickImg() {
      document.getElementById('upload').click();
    };

    view(vnode: ξ.CVnode) {
      return ξ(base,
          ξ('.ui.text.container', {
            style: 'margin-top: 1em;',
          },
          ξ('h1.ui.header',
              ξ('img.ui.circular.image', {
                src: '/user/' + UserModel.getUsername() + '/avatar',
              }),
              ξ('.content',
                  'MY PROFILE',
                  ξ('.sub.header', `Hello there, ${UserModel.getFullname()}!`)),
          ),
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
                  ξ('input#upload',
                      {onchange: (e) => {this.upload(e)},
                        type: 'File',
                        style: 'display: none;',
                        accept: '.png, .jpg, .jpeg;',
                      }),
                  ξ('img.ui.image.rounded.medium#picture',
                      {src: '/user/' + UserModel.getUsername() + '/avatar',
                        onclick: this.clickImg,
                        style: 'cursor: pointer;'},
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
                        data: Model.dataDonut,
                        options: Model.optionsDonut,
                      });
                    }})
              )
          )
      );
    }
};
