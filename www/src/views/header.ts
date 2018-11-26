import m from 'mithril';

export default {
  view(vnode: CVnode) {
    return m('h1.ui.header',
        m('a', {href: '/', oncreate: m.route.link},
            'ROASTER INC.'
        )
    );
  },
} as m.Component;
