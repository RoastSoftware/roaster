import ξ from 'mithril';
import base from './base';
import Network from '../services/network';
import Model from '../models/statistics';
import Chart from 'chart.js';
import {UserModel} from '../models/user';

class UserProfile implements ξ.ClassComponent {
    isFriend: boolean = false;

    hasFriend(username: string) {
        this.isFriend = false;
      for (const friend of UserModel.friends) {
        if (friend.friend == username) {
          this.isFriend = true;
          break;
        }  
      }
    };

    async friendCheck(username: string) {
      return Network.request<Array<Friend>>('GET', '/user/' +
            UserModel.getUsername() + '/friend')
          .then((result) => {
            UserModel.friends = result;
            this.hasFriend(username);
            ξ.redraw();
          });
    };

    async unFriend(username: string) {
      Network.request('DELETE', '/user/' + username + '/friend')
          .then(() => {
            this.friendCheck(username);
          });
    };

    async addFriend(username: string) {
      Network.request('POST', '/user/' + username + '/friend', {
        'friend': username,
      })
          .then(() => {
            this.friendCheck(username);
          });
    };

    oncreate(vnode: ξ.CVnodeDOM) {
      this.friendCheck(vnode.attrs.username);
    };

    view({attrs}) {
      const username = attrs.username;
      const fullname = attrs.fullname;
      const email = attrs.email;

      return ξ(base,
          ξ('.ui.text.container', {
            style: 'margin-top: 1em;',
          },
          ξ('h1.ui.header',
              ξ('img.ui.circular.image', {
                src: `/user/${username}/avatar`,
              }),
              ξ('.content', `${username}\'S PROFILE`.toUpperCase(),
                  ξ('.sub.header', `You are viewing ${fullname}'s profile.`)),
          ),
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
                this.isFriend ?
                ξ('button.ui.basic.red.button', {
                  onclick: () => {
                    this.unFriend(username);
                  },
                }, 'UNFOLLOW')
                :
                ξ('button.ui.basic.teal.button', {
                  onclick: () => {
                    this.addFriend(username);
                  },
                },
                'FOLLOW!'),
              ),
              ξ('.ui.column[style=min-height: 10em;]',
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
    downloadError: Error;
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
