import m from 'mithril';
import base from './base';

export default {
  view(vnode: CVnode) {
    return m(base,
        m('p', 'Let there be GRAPHS! '),
        m('p', 'later...')
    );
  },
} as m.Component;
