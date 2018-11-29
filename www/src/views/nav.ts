import ξ from 'mithril';
import {User} from '../models/user';
import roasterLogo from '../assets/icons/roaster-icon-teal.svg';

const navBarStyle = `\
border: 1px solid #094959;\
z-index: 9999;\
`;

const logotypeStyle = `\
margin: -2em 0em -2em 0em;\
height: 1.5em;\
color: #fff;\
`;

/**
 * Nav component provides a navigation bar for the top of the page.
 */
export default {
  /**
   * Creates a navigation bar.
   * @param {CVnode} vnode - Virtual node.
   * @return {CVnode}
   */
  view(vnode: CVnode) {
    return ξ('nav.ui.massive.borderless.stackable.menu.attached', {
      style: navBarStyle,
    }, [
      ξ('a.header.item', {href: '/', oncreate: ξ.route.link},
          ξ('img', {
            src: roasterLogo,
            style: logotypeStyle,
          }),
          'ROASTER INC.'
      ),

      ξ('a.item', {href: '/about', oncreate: ξ.route.link},
          ξ('i.question.circle.outline.icon'),
          'ABOUT'),

      ξ('a.item', {href: '/statistics', oncreate: ξ.route.link},
          ξ('i.chart.bar.icon'),
          'STATISTICS'),
      ξ('.right.menu',

          (User.isLoggedIn() ?
            ξ('a.item', {href: '/profile', oncreate: ξ.route.link},
                'PROFILE')
          :
            [
              ξ('a.item', {href: '/register', oncreate: ξ.route.link},
                  ξ('i.user.plus.icon'),
                  'REGISTER'),
              ξ('a.item', {href: '/login', oncreate: ξ.route.link},
                  ξ('i.sign.in.icon'),
                  'LOGIN'),
            ]
          ),
      ),
    ]);
  },
} as ξ.Component;
