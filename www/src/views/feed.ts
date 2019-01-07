import ξ from 'mithril';
import base from './base';
import Network, {encodeURL} from '../services/network';
import {UserModel} from '../models/user';
import moment from 'moment';

interface FeedItem {
  category: number
  name: string
  description: string
  username: string
  createTime: string
};

interface Feed {
  items: Array<FeedItem>;
}

class UserLink implements ξ.ClassComponent {
  view({attrs}) {
    const username: string = attrs.username;

    return ξ('a.user', {
      href: '/user/' + username,
      oncreate: ξ.route.link,
    }, username.toUpperCase());
  };
}

export class FeedList implements ξ.ClassComponent {
  capitalize(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
  };

  view({attrs}) {
    const feed: Feed = attrs.feed;

    return [
      (feed.items && feed.items.length > 0) ?
      feed.items.map((i: FeedItem) => {
        switch (i.category) {
          case 0:
            return ξ('.event',
                ξ('.label',
                    ξ('img', {
                      src: encodeURL('user', i.username, 'avatar'),
                    }),
                ),
                ξ('.content',
                    ξ('.summary',
                        ξ(UserLink, {'username': i.username}),
                        ` JUST GOT ROASTED!`.toUpperCase(),
                        ξ('span.score[style=margin-left: 5px;]',
                            ξ('i.trophy.icon[style=color: gold;]'),
                            `${i.title}`),
                        ξ('.date[style=float:right;]',
                            ξ('i.clock.outline.icon'),
                            moment(new Date(i.createTime)).fromNow())),
                    ξ('.extra.text', `\
Oh hot damn, this is my jam, `, ξ(UserLink, {username: i.username}), ` \
got `, ξ('strong', `${i.title} Roast® Score™`), ` when they Roast'd their \
${this.capitalize(i.description)} code.`),
                ),
            );
        }
      }) : '',
    ];
  }
}

class PaginationListItem implements ξ.ClassComponent {
  view({attrs}) {
    const page = attrs.page;
    const active = attrs.active;

    if (page == -1) {
      return ξ('a.disabled.item', '...');
    }

    if (active) {
      return ξ('a.active.item', page+1);
    }

    return ξ('a.item', {
      href: `/feed?page=${page}`,
      oncreate: ξ.route.link,
      onupdate: ξ.route.link,
    }, page+1);
  };
};

class PaginationList implements ξ.ClassComponent {
  view({attrs}) {
    const currentPage: number = attrs.currentPage;
    let pages: Array<number> = [];

    if (currentPage < 2) {
      pages = [0, 1, 2, -1];
    } else {
      pages = [0, -1, currentPage-1, currentPage, currentPage+1];
    };

    return [
      pages.map((p: number) => {
        return ξ(PaginationListItem, {page: p, active: p == currentPage});
      }),
    ];
  };
};

const menuActiveItemStyle = `\
border-bottom: 3px solid #f2f2f2;\
`;

const menuStyle = `\
border-top: none;
border-left: none;
border-right: none;
border-bottom: 1px solid #094959;\
margin-right: -1px;\
margin-left: -1px;\
`;

export default class Feed implements ξ.ClassComponent {
  currentPage: number = 0;
  currentCategory: string = 'global';
  lastUpdatedPage: number = -1;
  feed: Feed = {} as Feed;

  fetchFeed() {
    let categoryQuery: string= '';

    switch (this.currentCategory) {
      case 'global':
        break;
      case 'followees':
        categoryQuery += `&followees=true`;
      case 'you':
        categoryQuery += `&user=${UserModel.getUsername()}`;
    }

    Network.request<Feed>('GET', '/feed?page='
      + this.currentPage + categoryQuery)

        .then((feed: Feed) => {
          this.feed = feed;
          this.lastUpdatedPage = this.currentPage;
        });
  };

  updatePage() {
    this.currentPage = parseInt(ξ.route.param('page')) || 0;
  };

  updateCategory(category: string) {
    this.currentCategory = category;
    this.fetchFeed();
  }

  setItemActiveClass(category: string) {
    return this.currentCategory == category ? 'item active' : 'item';
  }

  setItemActiveStyle(category: string) {
    return this.currentCategory == category ? menuActiveItemStyle : '';
  }

  oncreate() {
    this.updatePage();
    this.fetchFeed();
  };

  onupdate() {
    this.updatePage();

    // Only fetch new page if the actual query parameter has updated.
    if (this.currentPage != this.lastUpdatedPage) {
      this.fetchFeed();
    }
  };

  view(): ξ.Children {
    return ξ(base,
        ξ('.ui.main.text.container[style=margin-top: 1em;]',
            ξ('h1.ui.header',
                ξ('i.feed.icon'),
                ξ('.content', 'FEED',
                    ξ('.sub.header',
                        'Check out what everyone has been up to!')),
            ),
            ξ('.ui.divider'),
            UserModel.isLoggedIn() ?
              ξ('.ui.top.attached.secondary.pointing.menu', {style: menuStyle},
                  ξ('a.item', {
                    class: this.setItemActiveClass('global'),
                    style: this.setItemActiveStyle('global'),
                    onclick: () => {
                      this.updateCategory('global');
                    },
                  },
                  ξ('i.globe.icon'), 'GLOBAL'),
                  ξ('a.item', {
                    class: this.setItemActiveClass('followees'),
                    style: this.setItemActiveStyle('followees'),
                    onclick: () => {
                      this.updateCategory('followees');
                    },
                  },
                  ξ('i.users.icon'), 'FOLLOWING'),
                  ξ('a.item', {
                    class: this.setItemActiveClass('you'),
                    style: this.setItemActiveStyle('you'),
                    onclick: () => {
                      this.updateCategory('you');
                    },
                  },
                  ξ('i.user.icon'), 'YOU'),
              ): '',
            ξ('.ui.bottom.attached.segment',
                (this.feed.items ?
                  ξ('.ui.feed',
                      ξ(FeedList, {
                        'feed': this.feed,
                      }),
                  ) : [
                    ξ('h2', 'You\'ve reached the end.'),
                    ξ('p', 'Welp, there are no more events to show you.'),
                  ]
                ),
            ),
            ξ('.ui.center.aligned.container[style=margin-bottom: 2em;]',
                ξ('.ui.pagination.compact.menu',
                    ξ(PaginationList, {currentPage: this.currentPage}),
                ),
            ),
        ),
    );
  };
};
