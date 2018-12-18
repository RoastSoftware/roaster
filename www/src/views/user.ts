import ξ from 'mithril';
import base from './base';
import Network from '../services/network';
import Model from '../models/statistics';
import Chart from 'chart.js';

class UserProfile implements ξ.ClassComponent {
  view(vnode: ξ.CVnode) {
    const username = vnode.attrs.username;
    const fullname = vnode.attrs.fullname;
    const email = vnode.attrs.email;

    return ξ(base,
        ξ('.ui.text.container', {
          style: 'margin-top: 1em;',
        },
        ξ('h1', `${username}\'S PROFILE`.toUpperCase()),
        ξ('.ui.divider')),
        ξ('.ui.main.text.container.two.column.stackable.grid',
            ξ('.ui.column',
                ξ('img.ui.image.rounded.medium#picture',
                    {src: '/user/' + username + '/avatar'},
                    'User profile picture.'),
                ξ('h2',
                    fullname),
                ξ('p',
                    ξ('i.user.icon'),
                    username),
                ξ('p',
                    ξ('i.mail.icon'),
                    email),
                ξ('.ui.placeholder',
                    ξ('ui.image')),

            ),
            ξ('.ui.column[min-height = 10em]',
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
                  }}),
            ),
        ),
    );
  }
};

export default class UserView implements ξ.ClassComponent {
    downloadError: APIError;
    user: User;
    ready: boolean;

    getUserStat() {
      console.log('getting user statistics');
      // TODO: Network.request statistics
    }

    oncreate(vnode: ξ.CVnodeDOM) {
      Network.request<User>('GET', '/user/' + vnode.attrs.username)
          .then((user: User) => {
            this.user = user;
            this.ready = true;
            ξ.redraw();
          });
    }

    view(vnode: ξ.CVnode) {
      // let idUser = ξ.route.param(username);
      // TODO: fill in temporary class wide user model.
      // see username to request will be available in vnode.attrs.username
      // see: https://mithril.js.org/route.html
      // })
      return this.ready ?
            ξ(UserProfile, {
              username: this.user.username,
              fullname: this.user.fullname,
              email: this.user.email,
            })
            : '';
    }
};
