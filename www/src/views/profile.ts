import ξ from 'mithril';
import base from './base';
import {UserModel} from '../models/user';
import Network from '../services/network';
import Model from '../models/statistics';
import Chart from 'chart.js';

export default class Profile implements ξ.ClassComponent {
    uploadError: APIError;
    downloadError: APIError;

    upload(e: Any) {
      const avatar = e.target.files[0];
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
          ξ('.ui.main.text.container.two.column.stackable.grid[style=margin-top: 2em;]',
              ξ('.ui.column', 
                  ξ('input#upload[type=file][style=display: none;]',
                      {onchange: this.upload}),
                  ξ('img.ui.image.rounded.medium#picture[style=cursor: pointer;]',
                      {src: '/user/' + UserModel.getUsername() + '/avatar',
                          onclick: this.clickImg},
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
                              'chart-area')as HTMLCanvasElement)
                              .getContext('2d');
                          new Chart(ctx, {
                              type: 'doughnut',
                              data: Model.dataDonut,
                              options: Model.optionsDonut
                          })
                      }})
              )
          )
      );
    }
};
