
import ξ from 'mithril';
import base from './base';
import {UserModel, User} from '../models/user';
import Network from '../services/network';
import Model from '../models/statistics';
import Chart from 'chart.js';

export default class UserView implements ξ.ClassComponent {
    downloadError: APIError;

    getUserStat() {
        console.log('getting user statistics');
        // TODO: Network.request statistics

    }

    view(vnode: ξ.CVnode) {
        Network.request<User>("GET", "/user/" + vnode.attrs.username).then((user: User) => {
            console.log("/user/" + vnode.attrs)
            // TODO: fill in temporary class wide user model.
            // see username to request will be available in vnode.attrs.username
            // see: https://mithril.js.org/route.html
        })
      return ξ(base,
          ξ('.ui.main.text.container.two.column.stackable.grid[style=margin-top: 2em;]',
              ξ('.ui.column', 
                  ξ('img.ui.image.rounded.medium#picture[style=cursor: pointer;]',
                      {src: '/user/' + UserModel.getUsername() + '/avatar'},
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
