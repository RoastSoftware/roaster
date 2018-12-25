import ξ from 'mithril';
import base from './base';
import Network from '../services/network';
import {RoastDoughnutStatisticsModel} from '../models/statistics';
import Chart from 'chart.js';

export class UserProfileHeader implements ξ.ClassComponent {
  view({attrs}) {
    return ξ('.ui.vertical.basic.segment.clearing', {
      style: 'padding: 0; margin: 0;',
    },
    ξ('h1.ui.header.left.floated[style=margin: 0;]',
        ξ('img.ui.circular.image', {
          src: `/user/${attrs.username}/avatar`,
        }),
        attrs.loggedIn ?
        ξ('.content', 'MY PROFILE',
            ξ('.sub.header', `Hello there, ${attrs.fullname}!`))
        :
        ξ('.content', `${attrs.username}\'S PROFILE`.toUpperCase(),
            ξ('.sub.header', `You are viewing ${attrs.fullname}'s profile.`)),
    ),
    ξ('.ui.right.floated.statistic',
        ξ('.value',
            ξ('i.trophy.icon[style=color: gold;]'),
            attrs.score,
        ),
        ξ('.label',
            'ROAST® SCORE™',
        ),
    ));
  };
};

class UserProfile implements ξ.ClassComponent {
  view({attrs}) {
    const username = attrs.username;
    const fullname = attrs.fullname;
    const email = attrs.email;
    const score = attrs.score;

    return ξ(base,
        ξ('.ui.text.container', {
          style: 'margin-top: 1em;',
        },
        ξ(UserProfileHeader, {
          username: username, fullname: fullname, score: score, loggedIn: false,
        }),
        ξ('.ui.divider')),
        ξ('.ui.main.text.container.two.column.stackable.grid',
            ξ('.ui.column',
                ξ('img.ui.image.rounded.medium#picture', {
                  src: `/user/${username}/avatar`,
                },
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
                      data: RoastDoughnutStatisticsModel.dataDonut,
                      options: RoastDoughnutStatisticsModel.optionsDonut,
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
    score: number = 0;
    ready: boolean;

    oncreate({attrs}) {
      Network.request<User>('GET', '/user/' + attrs.username)
          .then((user: User) => {
            this.user = user;
          });

      Network.request<Object>('GET', '/user/' + attrs.username + '/score')
          .then(({score}) => {
            this.score = score;
          });
    }

    view(vnode: ξ.CVnode) {
      return this.user ?
            ξ(UserProfile, {
              username: this.user.username,
              fullname: this.user.fullname,
              email: this.user.email,
              score: this.score,
            }): '';
    }
};
