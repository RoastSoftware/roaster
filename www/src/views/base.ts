import m from 'mithril';
import nav from './nav';
import header from './header';

export default {
  view(vnode: CVnode) {
    return [
      m(nav),
      m('.ui.container',
          m(header)
      ),
      m('.ui.main.container.segment.inverted', vnode.children),
    ];
  },
} as m.Component;
