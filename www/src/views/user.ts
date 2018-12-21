
import ξ from 'mithril';
import base from './base';
import Network from '../services/network';
import Model from '../models/statistics';
import Chart from 'chart.js';
import {UserModel, addFriend, unFriend, getFriends} from '../models/user';

class UserProfile implements ξ.ClassComponent {
    isFriend: boolean = false;
    friendCheck(username: string) {
        console.log("friendcheck");
        console.log(UserModel.friends);
        for (let friend of UserModel.friends){
            console.log("iterating");
            if (friend.friend == username){
                this.isFriend = true;
                console.log("Friend:" + friend);
                console.log("is friend?:" + this.isFriend);
                break;
            } else {
                this.isFriend = false;
            }
        }
    };
  view(vnode: ξ.CVnode) {
    const username = vnode.attrs.username;
    const fullname = vnode.attrs.fullname;
    const email = vnode.attrs.email;
      
    return ξ(base,
        ξ('.ui.main.text.container.two.column.stackable.grid', {
          style: 'margin-top: 2em;',
        },
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
                    unFriend(username);
                    getFriends();
                    this.friendCheck(username);
                    ξ.redraw();
                    // TODO: on success update friends and this.isFriend=false
                },
            }, 'UNFOLLOW')
            :
            ξ('button.ui.basic.teal.button', {
                onclick: () => {
                    addFriend(username);
                    getFriends();
                    this.friendCheck(username);
                    ξ.redraw();
                    // TODO: on success update friends and this.isFriend=true
                },
            },
                'FOLLOW!'),
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
              }})
        )
        )
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
