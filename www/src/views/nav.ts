import ξ from 'mithril';

const navBarStyle = `\
border: 1px solid #094959;\
z-index: 9999;\
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
      ξ('.ui.container',
          ξ('a.header.item', {href: '/', oncreate: ξ.route.link},
              ξ('i.brown.coffee.icon.logo'),
              'Roaster'
          ),
          ξ('a.item', {href: '/about', oncreate: ξ.route.link}, 'ABOUT'),
          ξ('a.item', {href: '/register', oncreate: ξ.route.link}, 'REGISTER'),
          ξ('a.item', {href: '/login', oncreate: ξ.route.link}, 'LOGIN'),
          ξ('a.item', {href: '/profile', oncreate: ξ.route.link}, 'PROFILE'),
          ξ('a.item', {
            href: '/statistics',
            oncreate: ξ.route.link},
          'STATISTICS')
      ),
    ]);
  },
} as ξ.Component;
