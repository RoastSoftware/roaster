import ξ from 'mithril';
import {UserModel} from '../models/user';
import Auth from '../services/auth';
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

const headerTextStyle = `\
color: #00b5ad;\
`;

/**
 * Nav component provides a navigation bar for the top of the page.
 */
export default class Nav implements ξ.ClassComponent {
  setItemActive(url: string): string {
    return ξ.route.get() == url ? 'item active' : 'item';
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
    return ξ('nav.ui.massive.borderless.stackable.menu.attached', {
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
                // For some reason this dissappears when the
                // class above gets evaluated, bug in Mithril?
                class: 'item',
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
