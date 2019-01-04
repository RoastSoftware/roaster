import ξ from 'mithril';
import Chart from 'chart.js';
import base from './base';
import Network from '../services/network';
import {UserModel} from '../models/user';
import {FeedList} from './feed';
import {
  RoastDoughnutStatisticsModel,
  StatisticsFilter,
} from '../models/statistics';

export class RoastRatio implements ξ.ClassComponent {
  chart: Chart;
  ctx: CanvasRenderingContext2D;
  ratio: RoastDoughnutStatisticsModel;

  oncreate({dom, attrs}) {
    this.ratio = new RoastDoughnutStatisticsModel(attrs.filter, attrs.username);

    this.ratio.update().then(() => {
      this.ctx = (dom as HTMLCanvasElement).getContext('2d');
      this.chart = new Chart(this.ctx,
          this.ratio.getConfig());
    });
  };

  onupdate({attrs}) {
    if (this.ratio.filter != attrs.filter) {
      this.ratio.filter = attrs.filter;
      this.ratio.update().then(() => {
        this.chart.data = this.ratio.getData();
        this.chart.update();
      });
    }
  };

  view() {
    return ξ('canvas');
  };
};

export class RoastScore implements ξ.ClassComponent {
  view({attrs}) {
    return ξ('.ui.right.floated.statistic',
        ξ('.value',
            ξ('i.trophy.icon[style=color: gold;]'),
            ` ${attrs.score}`,
        ),
        ξ('.label',
            'ROAST® SCORE™',
        ),
    );
  };
}

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
    ξ(RoastScore, {score: attrs.score}));
  };
};

export class UserFeed implements ξ.ClassComponent {
  username: string = '';
  feed: Feed = {} as Feed;
  error: Error;

  oncreate({attrs}) {
    this.username = attrs.username;
    this.fetchFeed();
  };

  fetchFeed() {
    const username = this.username;
    Network.request<Feed>('GET', `/feed?page=0&page-size=5&user=${username}`)
        .then((feed: Feed) => {
          this.feed = feed;
        }).catch((error) => {
          this.error = error;
        });
  };

  view(): ξ.Children {
    return [
      this.error ?
      ξ('.ui.error.message',
          ξ('.header',
              'Hmm, we failed to fetch the latest user events...',
          ),
          ξ('p', `${this.error.message}.`),
      ): '',
      this.feed.items ?
      ξ('.ui.feed', [
        ξ(FeedList, {
          'feed': this.feed,
        }),
      ]) : [
        ξ('h2', 'Oh no!'),
        ξ('p', 'This user hasn\'t got Roasted™ yet!'),
      ],
    ];
  };
}

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

    oncreate({attrs}) {
      this.friendCheck(attrs.username);
    };

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
            username: username,
            fullname: fullname,
            score: score,
            loggedIn: false,
          }),
          ξ('.ui.divider')),
          ξ('.ui.main.text.stackable.two.column.grid.container',
              ξ('.ui.row',
                  ξ('.column',
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
                      UserModel.isLoggedIn() ? [
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
                      ]: '',
                  ),
                  ξ('.column[min-height = 10em]',
                      ξ(RoastRatio, {
                        filter: StatisticsFilter.User,
                        username: username,
                      }),
                  ),
              ),
              ξ('.ui.one.column.row',
                  ξ('.column',
                      ξ('.ui.basic.segment',
                          ξ('h2.ui.dividing.header',
                              ξ('i.feed.icon'),
                              ξ('.content', 'USER FEED',
                                  ξ('.sub.header',
                                      `What has ${username} been up to?`)),
                          ),
                          ξ(UserFeed, {
                            username: username,
                          }),
                          ξ('.ui.basic.center.aligned.segment',
                              'Check out the ', ξ('a', {
                                href: '/feed',
                                oncreate: ξ.route.link,
                              }, 'feed page'), ' for more!',
                          ),
                      ),
                  ),
              ),
          ),
      );
    }
};

export default class UserView implements ξ.ClassComponent {
    downloadError: Error;
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
