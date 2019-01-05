// Forgive me Father, for I have imported jQuery.
// I pray that you guide to forgiveness and move past this dark spot in my life.
// In the name of the Father, Son, and Holy Spirit. Amen.
import $ from 'jquery';
window.$ = window.jQuery = $;
import 'jquery-ui/ui/effect.js'; // semantic-ui-search requires `easing` effect
$.fn.api = require('semantic-ui-api');
$.fn.search = require('semantic-ui-search');

import ξ from 'mithril';
import {UserModel} from '../models/user';
import Auth from '../services/auth';
import roasterLogo from '../assets/icons/roaster-icon-teal.svg';

const navBarStyle = `\
border: 1px solid #094959;\
z-index: 9999;\
min-height: 50px;\
`;

const logotypeStyle = `\
margin: -2em 0em -2em 0em;\
height: 1.5em;\
color: #fff;\
`;

const headerTextStyle = `\
color: #00b5ad;\
`;

class SearchItem implements ξ.ClassComponent {
  oncreate({dom}) {
    $(dom)
      .search({
        type          : 'category',
        minCharacters : 3,
        apiSettings   : {
          onResponse: function(githubResponse) {
            var
            response = {
              results : {}
            }
            ;
            // translate GitHub API response to work with search
            $.each(githubResponse.items, function(index, item) {
              var
              language   = item.language || 'Unknown',
                maxResults = 8
              ;
              if(index >= maxResults) {
                return false;
              }
              // create new language category
              if(response.results[language] === undefined) {
                response.results[language] = {
                  name    : language,
                  results : []
                };
              }
              // add result to category
              response.results[language].results.push({
                title       : item.name,
                description : item.description,
                url         : item.html_url
              });
            });
            return response;
          },
          url: '//api.github.com/search/repositories?q={query}'
        }
      });
  };

  view() {
    return ξ('.ui.fluid.category.search.loading.item',
        ξ('.ui.transparent.icon.input',
            ξ('input.prompt', {
              placeholder: 'Search people...',
              type: 'text',
            }),
            ξ('i.search.link.icon')),
        ξ('.results'));
  };
};

/**
 * Nav component provides a navigation bar for the top of the page.
 */
export default class Nav implements ξ.ClassComponent {
  setItemActive(url: string): string {
    return ξ.route.get() == url ? 'active' : '';
  };

  deauthenticate(): Promise {
    return Auth.logout()
        .then(() => {
          UserModel.empty();
          ξ.route.set('/');
        });
  };

  /**
   * Creates a navigation bar.
   * @param {CVnode} vnode - Virtual node.
   * @return {CVnode}
   */
  view(vnode: CVnode) {
    return ξ('nav.ui.borderless.stackable.menu.attached', {
      style: navBarStyle,
    }, [
      ξ('a.header.item', {
        href: '/',
        oncreate: ξ.route.link,
        style: headerTextStyle},
      ξ('img', {
        src: roasterLogo,
        style: logotypeStyle,
      }),
      'ROASTER INC.'
      ),

      // TODO: Make this DRY, generate the navbar instead?
      ξ('.right.menu',

          ξ(SearchItem),

          ξ.route.get() != '/' ?
          ξ('.item',
              ξ('a.ui.primary.button[style=max-height=0.1em;]', {
                href: '/',
                oncreate: ξ.route.link,
              }, 'GET ROASTED!')): '',

          ξ('a.item', {
            href: '/about',
            oncreate: ξ.route.link,
            class: this.setItemActive('/about')},
          ξ('i.question.circle.outline.icon'), 'ABOUT'),

          ξ('a.item', {
            href: '/statistics',
            oncreate: ξ.route.link,
            class: this.setItemActive('/statistics')},
          ξ('i.chart.bar.icon'), 'STATISTICS'),

          ξ('a.item', {
            href: '/feed',
            oncreate: ξ.route.link,
            class: this.setItemActive('/feed')},
          ξ('i.feed.icon'), 'FEED'),

          (UserModel.isLoggedIn() ?
            [
              ξ('a.item', {
                href: '/profile',
                oncreate: ξ.route.link,
                onupdate: ξ.route.link,
                class: this.setItemActive('/profile')},
              ξ('i.user.icon'),
              UserModel.getUsername().toUpperCase()),
              ξ('a.item', {
                href: '#!/',
                onclick: this.deauthenticate},
              ξ('i.sign.out.icon'), 'LOGOUT'),
            ]
          :
            [
              ξ('a.item', {
                href: '/register',
                oncreate: ξ.route.link,
                onupdate: ξ.route.link,
                class: this.setItemActive('/register')},
              ξ('i.user.plus.icon'), 'REGISTER'),
              ξ('a.item', {
                href: '/login',
                oncreate: ξ.route.link,
                onupdate: ξ.route.link,
                class: this.setItemActive('/login')},
              ξ('i.sign.in.icon'), 'LOGIN'),
            ]
          ),
      ),
    ]);
  };
};
