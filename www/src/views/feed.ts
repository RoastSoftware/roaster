import ξ from 'mithril';
import base from './base';
import Network from '../services/network';
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
  view(vnode: ξ.CVnode) {
    const username: string = vnode.attrs.username;

    return ξ('a.user', {
      href: `/user/${username.toLowerCase()}`,
      oncreate: ξ.route.link,
    }, username);
  };
}

class FeedList implements ξ.ClassComponent {
  capitalize(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
  };

  view(vnode: ξ.CVnode) {
    const feed: Feed = vnode.attrs.feed;

    return [
      (feed.items && feed.items.length > 0) ?
      feed.items.map((i: FeedItem) => {
        switch (i.category) {
          case 0:
            return ξ('.event',
                ξ('.label',
                    ξ('img', {
                      src: `/user/${i.username}/avatar`,
                    }),
                ),
                ξ('.content',
                    ξ('.summary',
                        ξ(UserLink, {'username': i.username.toUpperCase()}),
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
  view(vnode: ξ.CVnode) {
    const page = vnode.attrs.page;
    const active = vnode.attrs.active;

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
  view(vnode: ξ.CVnode) {
    const currentPage: number = vnode.attrs.currentPage;
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


export default class Feed implements ξ.ClassComponent {
  currentPage: number = 0;
  lastUpdatedPage: number = -1;
  feed: Feed = {} as Feed;

  fetchFeed() {
    Network.request<Feed>('GET', '/feed?page=' + this.currentPage)
        .then((feed: Feed) => {
          this.feed = feed;
          this.lastUpdatedPage = this.currentPage;
        });
  };

  updatePage() {
    this.currentPage = parseInt(ξ.route.param('page')) || 0;
  };

  oncreate(vnode: ξ.CVnodeDOM) {
    this.updatePage();
    this.fetchFeed();
  };

  onupdate(vnode: ξ.CVnodeDOM) {
    this.updatePage();

    // Only fetch new page if the actual query parameter has updated.
    if (this.currentPage != this.lastUpdatedPage) {
      this.fetchFeed();
    }
  };

  view(vnode: ξ.CVnode): ξ.Children {
    console.log(this.feed.items);
    return ξ(base,
        ξ('.ui.main.text.container[style=margin-top: 1em;]',
            ξ('h1', 'GLOBAL FEED'),
            ξ('.ui.divider'),
            (this.feed.items ?
            ξ('.ui.feed', [
              ξ(FeedList, {
                'feed': this.feed,
              }),
            ]) : [
              ξ('h2', 'You\'ve reached the end.'),
              ξ('p', 'Welp, there are no more events to show you.'),
            ]),
            ξ('.ui.center.aligned.container[style=margin-bottom: 2em;]',
                ξ('.ui.pagination.compact.menu',
                    ξ(PaginationList, {currentPage: this.currentPage}),
                ),
            ),
        ),
    );
  };
};
